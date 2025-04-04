package cpi

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableCPIMessageMapping() *plugin.Table {
	return &plugin.Table{
		Name:        "cpi_message_mapping",
		Description: "Message mappings.",
		Columns:     messageMappingColumns(),
		List: &plugin.ListConfig{
			Hydrate: listMessageMappings,
		},
		Get: &plugin.GetConfig{
			KeyColumns: messageMappingKeyColumns(),
			Hydrate:    getMessageMapping,
		},
	}
}

func messageMappingKeyColumns() []*plugin.KeyColumn {
	return designtimeArtifactKeyColumns()
}

func messageMappingColumns() []*plugin.Column {
	return designtimeArtifactColumns(nil)
}

func listMessageMappings(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (any, error) {
	return listEntities(ctx, d, h, messageMappings)
}

func getMessageMapping(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (any, error) {
	return getEntity(ctx, d, h, messageMappingByIDAndVersion, []parameter{
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
