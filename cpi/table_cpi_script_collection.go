package cpi

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableCPIScriptCollection() *plugin.Table {
	return &plugin.Table{
		Name:        "cpi_script_collection",
		Description: "Script collections.",
		Columns:     scriptCollectionColumns(),
		List: &plugin.ListConfig{
			Hydrate: listScriptCollections,
		},
		Get: &plugin.GetConfig{
			KeyColumns: scriptCollectionKeyColumns(),
			Hydrate:    getScriptCollection,
		},
	}
}

func scriptCollectionKeyColumns() []*plugin.KeyColumn {
	return designtimeArtifactKeyColumns()
}

func scriptCollectionColumns() []*plugin.Column {
	return designtimeArtifactColumns(nil)
}

func listScriptCollections(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (any, error) {
	return listEntities(ctx, d, h, scriptCollections)
}

func getScriptCollection(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (any, error) {
	return getEntity(ctx, d, h, scriptCollectionByIDAndVersion, []parameter{
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
