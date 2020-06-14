package zertocloudmanager

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nolanprewit1/terraform-provider-zertoCloudManager/api"
)

// Define the schema of the resource fields found the terraform configuration
func resourceZorg() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name used to identify the ZORG. The name must be unique",
				ForceNew:    false,
			},
			"crmidentifier": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    false,
				Description: "A internal CRM identification for the ZORG",
			},
		},
		// specify the CRUD operation functions
		Create: resourceZorgCreate,
		Read:   resourceZorgRead,
		Update: resourceZorgUpdate,
		Delete: resourceZorgDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
	}
}

func resourceZorgCreate(schema *schema.ResourceData, c interface{}) error {
	// Get the api client informaiton then pass it to the the URL and Key function
	apiClientInfo := c.(*api.Client)
	url, sessionKey := api.GetURLandKey(apiClientInfo)

	// Use the informaiton provided in the terraform configuration file to build a json object
	jsonData := map[string]string{
		"Name":          schema.Get("name").(string),
		"CrmIdentifier": schema.Get("crmidentifier").(string)}
	jsonValue, _ := json.Marshal(jsonData)

	// Post the json object to the api using the url and session key retrieved earlier
	client := &http.Client{}
	request, err := http.NewRequest("POST", url+"/zorgs", bytes.NewBuffer(jsonValue))
	request.Header.Set("x-zerto-session", sessionKey)
	request.Header.Set("Content-Type", "application/json")
	response, err := client.Do(request)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	}

	// remove the " " from the string
	zorgID, _ := ioutil.ReadAll(response.Body)
	zorgIDFormatted := strings.ReplaceAll(string(zorgID), "\"", "")

	// Set the schema id so that the resource can be referenced in later runs of terraform
	schema.SetId(zorgIDFormatted)

	return nil
}

func resourceZorgRead(schema *schema.ResourceData, c interface{}) error {
	// Get the api client informaiton then pass it to the the URL and Key function
	apiClientInfo := c.(*api.Client)
	url, sessionKey := api.GetURLandKey(apiClientInfo)

	// Get the zorg id
	zorgIDFormatted := schema.Id()

	// get the requested zorg
	client := &http.Client{}
	request, err := http.NewRequest("GET", url+"/zorgs/"+zorgIDFormatted, nil)
	request.Header.Set("x-zerto-session", sessionKey)
	response, err := client.Do(request)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	}

	// define the structure of the json response
	type Response struct {
		CrmIdentifier  string `json:"CrmIdentifier"`
		Name           string `json:"Name"`
		ZorgIdentifier string `json:"ZorgIdentifier"`
	}
	var apiResponse Response

	// read the response body
	data, _ := ioutil.ReadAll(response.Body)
	json.Unmarshal([]byte(data), &apiResponse)

	schema.SetId(zorgIDFormatted)
	schema.Set("name", apiResponse.Name)
	schema.Set("crmidentifier", apiResponse.CrmIdentifier)

	return nil
}

func resourceZorgUpdate(schema *schema.ResourceData, c interface{}) error {
	// Get the api client informaiton then pass it to the the URL and Key function
	apiClientInfo := c.(*api.Client)
	url, sessionKey := api.GetURLandKey(apiClientInfo)

	// Get the zorg id
	zorgIDFormatted := schema.Id()

	// Use the informaiton provided in the terraform configuration file to build a json object
	jsonData := map[string]string{
		"Name":          schema.Get("name").(string),
		"CrmIdentifier": schema.Get("crmidentifier").(string)}
	jsonValue, _ := json.Marshal(jsonData)

	// put the json object to the api using the url and session key retrieved earlier
	client := &http.Client{}
	request, err := http.NewRequest("PUT", url+"/zorgs/"+zorgIDFormatted, bytes.NewBuffer(jsonValue))
	request.Header.Set("x-zerto-session", sessionKey)
	request.Header.Set("Content-Type", "application/json")
	_, err = client.Do(request)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	}

	return nil
}

func resourceZorgDelete(schema *schema.ResourceData, c interface{}) error {
	// Get the api client informaiton then pass it to the the URL and Key function
	apiClientInfo := c.(*api.Client)
	url, sessionKey := api.GetURLandKey(apiClientInfo)

	// Get the zorg id
	zorgIDFormatted := schema.Id()

	// get the requested zorg
	client := &http.Client{}
	request, err := http.NewRequest("DELETE", url+"/zorgs/"+zorgIDFormatted, nil)
	request.Header.Set("x-zerto-session", sessionKey)
	_, err = client.Do(request)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	}

	schema.SetId("")

	return nil
}
