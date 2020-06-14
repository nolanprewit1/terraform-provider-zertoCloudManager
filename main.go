package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/nolanprewit1/terraform-provider-zertoCloudManager/zertocloudmanager"
)

func main() {
	opts := plugin.ServeOpts{
		ProviderFunc: zertocloudmanager.Provider,
	}
	plugin.Serve(&opts)
}
