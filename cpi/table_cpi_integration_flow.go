package cpi

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableCPIIntegrationFlow() *plugin.Table {
	return &plugin.Table{
		Name:        "cpi_integration_flow",
		Description: "Integration flows.",
		Columns:     integrationFlowColumns(),
		List: &plugin.ListConfig{
			Hydrate: listIntegrationFlows,
		},
		Get: &plugin.GetConfig{
			KeyColumns: integrationFlowKeyColumns(),
			Hydrate:    getIntegrationFlow,
		},
	}
}

func integrationFlowKeyColumns() []*plugin.KeyColumn {
	return designtimeArtifactKeyColumns()
}

func integrationFlowColumns() []*plugin.Column {
	return designtimeArtifactColumns([]*plugin.Column{
		{
			Name:        "sender",
			Description: "Sender.",
			Type:        proto.ColumnType_STRING,
			Transform:   transform.FromField("Sender"),
		},
		{
			Name:        "receiver",
			Description: "Receiver.",
			Type:        proto.ColumnType_STRING,
			Transform:   transform.FromField("Receiver"),
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
			Transform:   transform.FromField("CreatedAt").Transform(convertEpochTimestampToUTCDateTime),
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
			Transform:   transform.FromField("ModifiedAt").Transform(convertEpochTimestampToUTCDateTime),
		},
	})
}

func listIntegrationFlows(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (any, error) {
	return listEntities(ctx, d, h, integrationFlows)
}

func getIntegrationFlow(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (any, error) {
	return getEntity(ctx, d, h, integrationFlowByIDAndVersion, []parameter{
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
