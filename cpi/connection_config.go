package cpi

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/schema"
)

type cpiConfig struct {
	BaseURL        *string `cty:"base_url"        hcl:"base_url"`
	TokenURL       *string `cty:"token_url"       hcl:"token_url"`
	ClientID       *string `cty:"client_id"       hcl:"client_id"`
	ClientSecret   *string `cty:"client_secret"   hcl:"client_secret"`
	MaxConcurrency *int    `cty:"max_concurrency" hcl:"max_concurrency"`
	Timeout        *string `cty:"timeout"         hcl:"timeout"`
}

func configSchema() map[string]*schema.Attribute {
	return map[string]*schema.Attribute{
		"base_url": {
			Type:     schema.TypeString,
			Required: true,
		},
		"token_url": {
			Type:     schema.TypeString,
			Required: true,
		},
		"client_id": {
			Type:     schema.TypeString,
			Required: true,
		},
		"client_secret": {
			Type:     schema.TypeString,
			Required: true,
		},
		"max_concurrency": {
			Type:     schema.TypeInt,
			Required: false,
		},
		"timeout": {
			Type:     schema.TypeString,
			Required: false,
		},
	}
}

func ConfigInstance() any {
	return &cpiConfig{}
}

func GetConfig(connection *plugin.Connection) cpiConfig {
	if connection == nil || connection.Config == nil {
		return cpiConfig{}
	}

	config, _ := connection.Config.(cpiConfig)

	return config
}
