package cpi

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func Plugin(_ context.Context) *plugin.Plugin {
	return &plugin.Plugin{
		Name: "steampipe-plugin-cpi",
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			Schema:      configSchema(),
			NewInstance: ConfigInstance,
		},
		DefaultTransform: transform.FromGo().NullIfZero(),
		DefaultIgnoreConfig: &plugin.IgnoreConfig{
			ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"404"}),
		},
		TableMap: map[string]*plugin.Table{
			"cpi_integration_flow":    tableCPIIntegrationFlow(),
			"cpi_integration_package": tableCPIIntegrationPackage(),
			"cpi_message_mapping":     tableCPIMessageMapping(),
			"cpi_script_collection":   tableCPIScriptCollection(),
			"cpi_value_mapping":       tableCPIValueMapping(),
		},
	}
}
