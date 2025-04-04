package main

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/vadimklimov/steampipe-plugin-cpi/cpi"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{PluginFunc: cpi.Plugin})
}
