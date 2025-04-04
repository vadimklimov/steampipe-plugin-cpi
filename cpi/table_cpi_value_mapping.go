package cpi

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableCPIValueMapping() *plugin.Table {
	return &plugin.Table{
		Name:        "cpi_value_mapping",
		Description: "Value mappings.",
		Columns:     valueMappingColumns(),
		List: &plugin.ListConfig{
			Hydrate: listValueMappings,
		},
		Get: &plugin.GetConfig{
			KeyColumns: valueMappingKeyColumns(),
			Hydrate:    getValueMapping,
		},
	}
}

func valueMappingKeyColumns() []*plugin.KeyColumn {
	return designtimeArtifactKeyColumns()
}

func valueMappingColumns() []*plugin.Column {
	return designtimeArtifactColumns(nil)
}

func listValueMappings(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (any, error) {
	return listEntities(ctx, d, h, valueMappings)
}

func getValueMapping(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (any, error) {
	return getEntity(ctx, d, h, valueMappingByIDAndVersion, []parameter{
		{
			Name:     "id",
			Required: true,
		},
		{
			Name:         "version",
			Required:     false,
			DefaultValue: "active",
		},
	})
}
