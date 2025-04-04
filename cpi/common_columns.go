package cpi

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func designtimeArtifactKeyColumns() []*plugin.KeyColumn {
	return []*plugin.KeyColumn{
		{
			Name:    "id",
			Require: plugin.Required,
		},
		{
			Name:    "version",
			Require: plugin.Optional,
		},
	}
}

func commonColumnsForDesigntimeArtifact() []*plugin.Column {
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
			Name:        "package_id",
			Description: "Integration package ID.",
			Type:        proto.ColumnType_STRING,
			Transform:   transform.FromField("PackageID"),
		},
		{
			Name:        "name",
			Description: "Name.",
			Type:        proto.ColumnType_STRING,
			Transform:   transform.FromField("Name"),
		},
		{
			Name:        "description",
			Description: "Description.",
			Type:        proto.ColumnType_STRING,
			Transform:   transform.FromField("Description"),
		},
	}
}

func designtimeArtifactColumns(columns []*plugin.Column) []*plugin.Column {
	return append(commonColumnsForDesigntimeArtifact(), columns...)
}
