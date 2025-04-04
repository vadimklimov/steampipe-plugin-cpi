package cpi

import (
	"context"
	"errors"
	"fmt"
	"runtime"
	"strings"
	"sync"
	"time"

	"github.com/hashicorp/go-hclog"
	"github.com/panjf2000/ants/v2"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"golang.org/x/oauth2/clientcredentials"
	"resty.dev/v3"
)

type apiClient struct {
	httpClient     *resty.Client
	maxConcurrency int
}

type restyToSteampipeLoggerAdapter struct {
	steampipeLogger hclog.Logger
}

type parameter struct {
	Name         string
	Required     bool
	DefaultValue string
}

type integrationPackage struct {
	ID                *string `json:"Id"`
	Version           *string `json:"Version"`
	Name              *string `json:"Name"`
	ShortText         *string `json:"ShortText"`
	Description       *string `json:"Description"`
	Vendor            *string `json:"Vendor"`
	PartnerContent    *bool   `json:"PartnerContent"`
	Mode              *string `json:"Mode"`
	UpdateAvailable   *bool   `json:"UpdateAvailable"`
	SupportedPlatform *string `json:"SupportedPlatform"`
	Products          *string `json:"Products"`
	Keywords          *string `json:"Keywords"`
	Countries         *string `json:"Countries"`
	Industries        *string `json:"Industries"`
	LineOfBusiness    *string `json:"LineOfBusiness"`
	ResourceID        *string `json:"ResourceId"`
	CreatedBy         *string `json:"CreatedBy"`
	CreationDate      *int64  `json:"CreationDate,string"`
	ModifiedBy        *string `json:"ModifiedBy"`
	ModifiedDate      *int64  `json:"ModifiedDate,string"`
}

type integrationFlow struct {
	ID          *string `json:"Id"`
	Version     *string `json:"Version"`
	PackageID   *string `json:"PackageId"`
	Name        *string `json:"Name"`
	Description *string `json:"Description"`
	Sender      *string `json:"Sender"`
	Receiver    *string `json:"Receiver"`
	CreatedBy   *string `json:"CreatedBy"`
	CreatedAt   *int64  `json:"CreatedAt,string"`
	ModifiedBy  *string `json:"ModifiedBy"`
	ModifiedAt  *int64  `json:"ModifiedAt,string"`
}

type messageMapping struct {
	ID          *string `json:"Id"`
	Version     *string `json:"Version"`
	PackageID   *string `json:"PackageId"`
	Name        *string `json:"Name"`
	Description *string `json:"Description"`
}

type scriptCollection struct {
	ID          *string `json:"Id"`
	Version     *string `json:"Version"`
	PackageID   *string `json:"PackageId"`
	Name        *string `json:"Name"`
	Description *string `json:"Description"`
}

type valueMapping struct {
	ID          *string `json:"Id"`
	Version     *string `json:"Version"`
	PackageID   *string `json:"PackageId"`
	Name        *string `json:"Name"`
	Description *string `json:"Description"`
}

var (
	instance *apiClient
	once     sync.Once
)

func (l *restyToSteampipeLoggerAdapter) Errorf(format string, v ...any) {
	l.steampipeLogger.Error(fmt.Sprintf(format, v...))
}

func (l *restyToSteampipeLoggerAdapter) Warnf(format string, v ...any) {
	l.steampipeLogger.Warn(fmt.Sprintf(format, v...))
}

func (l *restyToSteampipeLoggerAdapter) Debugf(format string, v ...any) {
	l.steampipeLogger.Debug(fmt.Sprintf(format, v...))
}

func client(ctx context.Context, d *plugin.QueryData) (*apiClient, error) {
	var initError error

	once.Do(func() {
		const (
			logPrefix             = "service.client"
			basePath              = "api/v1"
			defaultTimeoutSeconds = 30
		)

		logger := plugin.Logger(ctx)

		config := GetConfig(d.Connection)

		maxConcurrency := runtime.NumCPU()

		if config.MaxConcurrency != nil {
			maxConcurrency = *config.MaxConcurrency
		}

		timeout := defaultTimeoutSeconds * time.Second

		if config.Timeout != nil {
			t, err := time.ParseDuration(*config.Timeout)
			if err != nil {
				logger.Warn(logPrefix, "config_parameter_parsing_error: timeout", err)
			} else {
				timeout = t
			}
		}

		oauthConfig := &clientcredentials.Config{
			TokenURL:     *config.TokenURL,
			ClientID:     *config.ClientID,
			ClientSecret: *config.ClientSecret,
		}

		httpClient := resty.NewWithClient(oauthConfig.Client(ctx)).
			SetContext(ctx).
			SetLogger(&restyToSteampipeLoggerAdapter{logger}).
			SetDebug(logger.IsDebug()).
			SetTimeout(timeout).
			SetBaseURL(*config.BaseURL + "/" + basePath).
			SetHeaders(map[string]string{
				"Accept":     "application/json",
				"User-Agent": "Steampipe",
			})

		instance = &apiClient{
			httpClient:     httpClient,
			maxConcurrency: maxConcurrency,
		}

		logger.Debug(logPrefix, "client_config: base_url", instance.httpClient.BaseURL())
		logger.Debug(logPrefix, "client_config: token_url", oauthConfig.TokenURL)
		logger.Debug(logPrefix, "client_config: max_concurrency", instance.maxConcurrency)
		logger.Debug(logPrefix, "client_config: timeout", instance.httpClient.Timeout().String())
	})

	return instance, initError
}

func send(request *resty.Request) (*resty.Response, error) {
	response, err := request.Send()
	if err != nil {
		return nil, fmt.Errorf("failed to call API: %w", err)
	}

	if response.IsError() {
		return nil, fmt.Errorf("failed to complete API request: response status: %s", response.Status())
	}

	return response, nil
}

func run[I any, O any](
	ctx context.Context,
	client *apiClient,
	values []I,
	task func(context.Context, *apiClient, I) ([]O, error),
) ([]O, []error, error) {
	var (
		results []O
		errs    []error
		wg      sync.WaitGroup
	)

	resultc := make(chan O)
	errc := make(chan error)

	pool, err := ants.NewPoolWithFunc(client.maxConcurrency, func(value any) {
		defer wg.Done()

		entries, err := task(ctx, client, value.(I))
		if err != nil {
			select {
			case errc <- err:
			case <-ctx.Done():
				return
			}

			return
		}

		for _, entry := range entries {
			select {
			case resultc <- entry:
			case <-ctx.Done():
				return
			}
		}
	})
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create worker pool: %w", err)
	}

	defer pool.Release()

	done := make(chan struct{})
	go func() {
		defer close(done)

		for {
			select {
			case result, ok := <-resultc:
				if !ok {
					resultc = nil
				} else {
					results = append(results, result)
				}
			case err, ok := <-errc:
				if !ok {
					errc = nil
				} else {
					errs = append(errs, err)
				}
			case <-ctx.Done():
				return
			}

			if resultc == nil && errc == nil {
				return
			}
		}
	}()

	for _, value := range values {
		wg.Add(1)

		_ = pool.Invoke(value)
	}

	wg.Wait()
	close(resultc)
	close(errc)
	<-done

	if ctx.Err() != nil {
		return results, errs, fmt.Errorf("failed to execute parallel tasks: %w", ctx.Err())
	}

	return results, errs, nil
}

func listEntities[T any](
	ctx context.Context,
	d *plugin.QueryData,
	_ *plugin.HydrateData,
	lister func(context.Context, *apiClient) ([]T, error),
) (any, error) {
	client, err := client(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("service.listEntities", "client_init_error", err)

		return nil, fmt.Errorf("failed to initialize API client: %w", err)
	}

	entities, err := lister(ctx, client)
	if err != nil {
		plugin.Logger(ctx).Error("service.listEntities", "error", err)

		return nil, fmt.Errorf("failed to execute lister function: %w", err)
	}

	for _, entity := range entities {
		d.StreamListItem(ctx, entity)

		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	return nil, nil
}

func getEntity[T any](
	ctx context.Context,
	d *plugin.QueryData,
	_ *plugin.HydrateData,
	getter func(context.Context, *apiClient, map[string]string) (*T, error),
	parameters []parameter,
) (any, error) {
	params := make(map[string]string, len(parameters))
	missingRequiredParams := make([]string, 0)

	for _, parameter := range parameters {
		name := parameter.Name

		value := d.EqualsQualString(parameter.Name)
		if value == "" {
			if parameter.Required {
				missingRequiredParams = append(missingRequiredParams, name)
			}

			value = parameter.DefaultValue
		}

		params[name] = value
	}

	if len(missingRequiredParams) > 0 {
		return nil, fmt.Errorf("SQL query error: following column(s) cannot be empty in WHERE clause: %s",
			strings.Join(missingRequiredParams, ", "),
		)
	}

	client, err := client(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("service.getEntity", "client_init_error", err)

		return nil, fmt.Errorf("failed to initialize API client: %w", err)
	}

	entity, err := getter(ctx, client, params)
	if err != nil {
		plugin.Logger(ctx).Error("service.getEntity", "error", err)

		return nil, fmt.Errorf("failed to execute getter function: %w", err)
	}

	return entity, nil
}

func integrationPackages(
	ctx context.Context,
	client *apiClient,
) ([]integrationPackage, error) {
	type responseBody struct {
		Data struct {
			Results []integrationPackage `json:"results"`
		} `json:"d"`
	}

	request := client.httpClient.R().
		SetMethod(resty.MethodGet).
		SetURL("IntegrationPackages").
		SetResult(&responseBody{})

	response, err := send(request)
	if err != nil {
		plugin.Logger(ctx).Error("service.integrationPackages", "api_error", err)

		return nil, fmt.Errorf("failed to retrieve integration packages: %w", err)
	}

	return response.Result().(*responseBody).Data.Results, nil
}

func integrationPackageByID(
	ctx context.Context,
	client *apiClient,
	parameters map[string]string,
) (*integrationPackage, error) {
	type responseBody struct {
		Data integrationPackage `json:"d"`
	}

	request := client.httpClient.R().
		SetMethod(resty.MethodGet).
		SetPathParam("id", parameters["id"]).
		SetURL("IntegrationPackages('{id}')").
		SetResult(&responseBody{})

	response, err := send(request)
	if err != nil {
		plugin.Logger(ctx).Error("service.integrationPackageByID", "api_error", err)

		return nil, fmt.Errorf("failed to retrieve integration package by ID: %w", err)
	}

	return &response.Result().(*responseBody).Data, nil
}

func integrationFlows(
	ctx context.Context,
	client *apiClient,
) ([]integrationFlow, error) {
	pkgs, err := integrationPackages(ctx, client)
	if err != nil {
		plugin.Logger(ctx).Error("service.integrationFlows", "api_error", err)

		return nil, err
	}

	pkgIDs := make([]string, 0, len(pkgs))
	for _, pkg := range pkgs {
		pkgIDs = append(pkgIDs, *pkg.ID)
	}

	iflws, errs, sysErr := run(ctx, client, pkgIDs, integrationFlowsByPackageID)

	if sysErr != nil {
		plugin.Logger(ctx).Error("service.integrationFlows", "system_error", err)

		return nil, fmt.Errorf("failed to retrieve integration flows due to system error: %w", err)
	}

	if len(errs) > 0 {
		plugin.Logger(ctx).Error("service.integrationFlows", "api_error", err)

		return nil, fmt.Errorf("failed to retrieve integration flows: %w", errors.Join(errs...))
	}

	return iflws, nil
}

func integrationFlowByIDAndVersion(
	ctx context.Context,
	client *apiClient,
	parameters map[string]string,
) (*integrationFlow, error) {
	type responseBody struct {
		Data integrationFlow `json:"d"`
	}

	request := client.httpClient.R().
		SetMethod(resty.MethodGet).
		SetPathParams(map[string]string{
			"id":      parameters["id"],
			"version": parameters["version"],
		}).
		SetURL("IntegrationDesigntimeArtifacts(Id='{id}',Version='{version}')").
		SetResult(&responseBody{})

	response, err := send(request)
	if err != nil {
		plugin.Logger(ctx).Error("service.integrationFlowByIDAndVersion", "api_error", err)

		return nil, fmt.Errorf("failed to retrieve integration flow by ID and version: %w", err)
	}

	return &response.Result().(*responseBody).Data, nil
}

func integrationFlowsByPackageID(
	ctx context.Context,
	client *apiClient,
	packageID string,
) ([]integrationFlow, error) {
	type responseBody struct {
		Data struct {
			Results []integrationFlow `json:"results"`
		} `json:"d"`
	}

	request := client.httpClient.R().
		SetMethod(resty.MethodGet).
		SetPathParam("packageID", packageID).
		SetURL("IntegrationPackages('{packageID}')/IntegrationDesigntimeArtifacts").
		SetResult(&responseBody{})

	response, err := send(request)
	if err != nil {
		plugin.Logger(ctx).Error("service.integrationFlowsByPackageID", "api_error", err)

		return nil, fmt.Errorf("failed to retrieve integration flows by package ID: %w", err)
	}

	return response.Result().(*responseBody).Data.Results, nil
}

func messageMappings(
	ctx context.Context,
	client *apiClient,
) ([]messageMapping, error) {
	type responseBody struct {
		Data struct {
			Results []messageMapping `json:"results"`
		} `json:"d"`
	}

	request := client.httpClient.R().
		SetMethod(resty.MethodGet).
		SetURL("MessageMappingDesigntimeArtifacts").
		SetResult(&responseBody{})

	response, err := send(request)
	if err != nil {
		plugin.Logger(ctx).Error("service.messageMappings", "api_error", err)

		return nil, fmt.Errorf("failed to retrieve message mappings: %w", err)
	}

	return response.Result().(*responseBody).Data.Results, nil
}

func messageMappingByIDAndVersion(
	ctx context.Context,
	client *apiClient,
	parameters map[string]string,
) (*messageMapping, error) {
	type responseBody struct {
		Data messageMapping `json:"d"`
	}

	request := client.httpClient.R().
		SetMethod(resty.MethodGet).
		SetPathParams(map[string]string{
			"id":      parameters["id"],
			"version": parameters["version"],
		}).
		SetURL("MessageMappingDesigntimeArtifacts(Id='{id}',Version='{version}')").
		SetResult(&responseBody{})

	response, err := send(request)
	if err != nil {
		plugin.Logger(ctx).Error("service.messageMappingByIDAndVersion", "api_error", err)

		return nil, fmt.Errorf("failed to retrieve message mapping by ID and version: %w", err)
	}

	return &response.Result().(*responseBody).Data, nil
}

func scriptCollections(
	ctx context.Context,
	client *apiClient,
) ([]scriptCollection, error) {
	pkgs, err := integrationPackages(ctx, client)
	if err != nil {
		plugin.Logger(ctx).Error("service.scriptCollections", "api_error", err)

		return nil, err
	}

	pkgIDs := make([]string, 0, len(pkgs))
	for _, pkg := range pkgs {
		pkgIDs = append(pkgIDs, *pkg.ID)
	}

	scs, errs, sysErr := run(ctx, client, pkgIDs, scriptCollectionsByPackageID)

	if sysErr != nil {
		plugin.Logger(ctx).Error("service.scriptCollections", "system_error", err)

		return nil, fmt.Errorf("failed to retrieve script collections due to system error: %w", err)
	}

	if len(errs) > 0 {
		plugin.Logger(ctx).Error("service.scriptCollections", "api_error", err)

		return nil, fmt.Errorf("failed to retrieve script collections: %w", errors.Join(errs...))
	}

	return scs, nil
}

func scriptCollectionByIDAndVersion(
	ctx context.Context,
	client *apiClient,
	parameters map[string]string,
) (*scriptCollection, error) {
	type responseBody struct {
		Data scriptCollection `json:"d"`
	}

	request := client.httpClient.R().
		SetMethod(resty.MethodGet).
		SetPathParams(map[string]string{
			"id":      parameters["id"],
			"version": parameters["version"],
		}).
		SetURL("ScriptCollectionDesigntimeArtifacts(Id='{id}',Version='{version}')").
		SetResult(&responseBody{})

	response, err := send(request)
	if err != nil {
		plugin.Logger(ctx).Error("service.scriptCollectionByIDAndVersion", "api_error", err)

		return nil, fmt.Errorf("failed to retrieve script collection by ID and version: %w", err)
	}

	return &response.Result().(*responseBody).Data, nil
}

func scriptCollectionsByPackageID(
	ctx context.Context,
	client *apiClient,
	packageID string,
) ([]scriptCollection, error) {
	type responseBody struct {
		Data struct {
			Results []scriptCollection `json:"results"`
		} `json:"d"`
	}

	request := client.httpClient.R().
		SetMethod(resty.MethodGet).
		SetPathParam("packageID", packageID).
		SetURL("IntegrationPackages('{packageID}')/ScriptCollectionDesigntimeArtifacts").
		SetResult(&responseBody{})

	response, err := send(request)
	if err != nil {
		plugin.Logger(ctx).Error("service.scriptCollectionsByPackageID", "api_error", err)

		return nil, fmt.Errorf("failed to retrieve script collections by package ID: %w", err)
	}

	return response.Result().(*responseBody).Data.Results, nil
}

func valueMappings(
	ctx context.Context,
	client *apiClient,
) ([]valueMapping, error) {
	type responseBody struct {
		Data struct {
			Results []valueMapping `json:"results"`
		} `json:"d"`
	}

	request := client.httpClient.R().
		SetMethod(resty.MethodGet).
		SetURL("ValueMappingDesigntimeArtifacts").
		SetResult(&responseBody{})

	response, err := send(request)
	if err != nil {
		plugin.Logger(ctx).Error("service.valueMappings", "api_error", err)

		return nil, fmt.Errorf("failed to retrieve value mappings: %w", err)
	}

	return response.Result().(*responseBody).Data.Results, nil
}

func valueMappingByIDAndVersion(
	ctx context.Context,
	client *apiClient,
	parameters map[string]string,
) (*valueMapping, error) {
	type responseBody struct {
		Data valueMapping `json:"d"`
	}

	request := client.httpClient.R().
		SetMethod(resty.MethodGet).
		SetPathParams(map[string]string{
			"id":      parameters["id"],
			"version": parameters["version"],
		}).
		SetURL("ValueMappingDesigntimeArtifacts(Id='{id}',Version='{version}')").
		SetResult(&responseBody{})

	response, err := send(request)
	if err != nil {
		plugin.Logger(ctx).Error("service.valueMappingByIDAndVersion", "api_error", err)

		return nil, fmt.Errorf("failed to retrieve value mapping by ID and version: %w", err)
	}

	return &response.Result().(*responseBody).Data, nil
}
