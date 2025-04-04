package cpi

import (
	"context"
	"time"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func convertEpochTimestampToUTCDateTime(_ context.Context, d *transform.TransformData) (any, error) {
	epochTimestamp := d.Value.(*int64)
	if epochTimestamp != nil {
		return time.UnixMilli(*epochTimestamp).UTC().Format(time.RFC3339), nil
	}

	return nil, nil
}
