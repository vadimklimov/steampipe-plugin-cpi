package cpi

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableCPIIntegrationPackage() *plugin.Table {
	return &plugin.Table{
		Name:        "cpi_integration_package",
		Description: "Integration packages.",
		Columns:     integrationPackageColumns(),
		List: &plugin.ListConfig{
			Hydrate: listIntegrationPackages,
		},
		Get: &plugin.GetConfig{
			KeyColumns: integrationPackageKeyColumns(),
			Hydrate:    getIntegrationPackage,
		},
	}
}

func integrationPackageKeyColumns() []*plugin.KeyColumn {
	return plugin.SingleColumn("id")
}

func integrationPackageColumns() []*plugin.Column {
	return []*plugin.Column{
		{
			Name:        "id",
			Description: "ID.",
			Type:        proto.ColumnType_STRING,
			Transform:   transform.FromField("ID"),
		},
		{
			Name:        "version",
			Description: "Version.",
			Type:        proto.ColumnType_STRING,
			Transform:   transform.FromField("Version"),
		},
		{
			Name:        "name",
			Description: "Name.",
			Type:        proto.ColumnType_STRING,
			Transform:   transform.FromField("Name"),
		},
		{
			Name:        "short_text",
			Description: "Short text.",
			Type:        proto.ColumnType_STRING,
			Transform:   transform.FromField("ShortText"),
		},
		{
			Name:        "description",
			Description: "Description.",
			Type:        proto.ColumnType_STRING,
			Transform:   transform.FromField("Description"),
		},
		{
			Name:        "vendor",
			Description: "Vendor.",
			Type:        proto.ColumnType_STRING,
			Transform:   transform.FromField("Vendor"),
		},
		{
			Name:        "partner_content",
			Description: "Is partner content.",
			Type:        proto.ColumnType_BOOL,
			Transform:   transform.FromField("PartnerContent"),
		},
		{
			Name:        "mode",
			Description: "Mode.",
			Type:        proto.ColumnType_STRING,
			Transform:   transform.FromField("Mode"),
		},
		{
			Name:        "update_available",
			Description: "Is update available.",
			Type:        proto.ColumnType_BOOL,
			Transform:   transform.FromField("UpdateAvailable"),
		},
		{
			Name:        "supported_platform",
			Description: "Supported platform.",
			Type:        proto.ColumnType_STRING,
			Transform:   transform.FromField("SupportedPlatform"),
		},
		{
			Name:        "products",
			Description: "Products.",
			Type:        proto.ColumnType_STRING,
			Transform:   transform.FromField("Products"),
		},
		{
			Name:        "keywords",
			Description: "Keywords.",
			Type:        proto.ColumnType_STRING,
			Transform:   transform.FromField("Keywords"),
		},
		{
			Name:        "countries",
			Description: "Countries.",
			Type:        proto.ColumnType_STRING,
			Transform:   transform.FromField("Countries"),
		},
		{
			Name:        "industries",
			Description: "Industries.",
			Type:        proto.ColumnType_STRING,
			Transform:   transform.FromField("Industries"),
		},
		{
			Name:        "line_of_business",
			Description: "Line of business.",
			Type:        proto.ColumnType_STRING,
			Transform:   transform.FromField("LineOfBusiness"),
		},
		{
			Name:        "resource_id",
			Description: "Resource ID.",
			Type:        proto.ColumnType_STRING,
			Transform:   transform.FromField("ResourceID"),
		},
		{
			Name:        "created_by",
			Description: "Created by.",
			Type:        proto.ColumnType_STRING,
			Transform:   transform.FromField("CreatedBy"),
		},
		{
			Name:        "created_at",
			Description: "Created at.",
			Type:        proto.ColumnType_TIMESTAMP,
			Transform:   transform.FromField("CreationDate").Transform(convertEpochTimestampToUTCDateTime),
		},
		{
			Name:        "modified_by",
			Description: "Modified by.",
			Type:        proto.ColumnType_STRING,
			Transform:   transform.FromField("ModifiedBy"),
		},
		{
			Name:        "modified_at",
			Description: "Modified at.",
			Type:        proto.ColumnType_TIMESTAMP,
			Transform:   transform.FromField("ModifiedDate").Transform(convertEpochTimestampToUTCDateTime),
		},
	}
}

func listIntegrationPackages(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (any, error) {
	return listEntities(ctx, d, h, integrationPackages)
}

func getIntegrationPackage(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (any, error) {
	return getEntity(ctx, d, h, integrationPackageByID, []parameter{
		{
			Name:     "id",
			Required: true,
		},
	})
}
