package zertocloudmanager

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"github.com/nolanprewit1/terraform-provider-zertoCloudManager/api"
)

// Provider ...define the provider schema, resources and configuration
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema:        providerSchema(),
		ResourcesMap:  providerResources(),
		ConfigureFunc: providerConfigure,
	}
}

// Provider schema. This is what information is need to connect to the provider.
func providerSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"address": {
			Type:        schema.TypeString,
			Required:    true,
			DefaultFunc: schema.EnvDefaultFunc("SERVICE_ADDRESS", ""),
			Description: "The IP address or hostname of the Zerto Cloud Manager.",
		},
		"port": {
			Type:        schema.TypeInt,
			Required:    true,
			DefaultFunc: schema.EnvDefaultFunc("SERVICE_PORT", ""),
			Description: "The API port of Zerto Cloud Manager. Default is 9989",
		},
		"username": {
			Type:        schema.TypeString,
			Required:    true,
			DefaultFunc: schema.EnvDefaultFunc("SERVICE_USERNAME", ""),
			Description: "Local or domain credentials of the Zerto Cloud Manager server.",
		},
		"password": {
			Type:        schema.TypeString,
			Required:    true,
			DefaultFunc: schema.EnvDefaultFunc("SERVICE_PASSWORD", ""),
			Description: "Local or domain credentials of the Zerto Cloud Manager server.",
		},
	}
}

// Define the resources available from the provider
func providerResources() map[string]*schema.Resource {
	return map[string]*schema.Resource{
		"zertocloudmanager_zorg": resourceZorg(),
	}
}

// Get the provider configuration information and pass it to the Client struct to reference later
func providerConfigure(schema *schema.ResourceData) (interface{}, error) {
	address := schema.Get("address").(string)
	port := schema.Get("port").(int)
	username := schema.Get("username").(string)
	password := schema.Get("password").(string)
	return api.ClientInfo(address, port, username, password), nil
}
