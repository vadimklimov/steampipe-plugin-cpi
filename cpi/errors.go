package cpi

import (
	"context"
	"strings"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func shouldIgnoreErrors(ignoreErrors []string) plugin.ErrorPredicateWithContext {
	return func(_ context.Context, _ *plugin.QueryData, _ *plugin.HydrateData, err error) bool {
		if err != nil {
			for _, pattern := range ignoreErrors {
				if strings.Contains(err.Error(), pattern) {
					return true
				}
			}
		}

		return false
	}
}
