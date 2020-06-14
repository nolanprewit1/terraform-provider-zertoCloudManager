package api

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"strconv"
)

// Client ...struct to define client fields used to connect to the api
type Client struct {
	address  string
	port     int
	username string
	password string
}

// ClientInfo ...a funciton to save the terraform code values to the Client struct
func ClientInfo(address string, port int, username string, password string) *Client {
	return &Client{
		address:  address,
		port:     port,
		username: username,
		password: password,
	}
}

// GetURLandKey ...pass in the client struct values and return the full api url and header api key
func GetURLandKey(c *Client) (string, string) {
	// skip verifying the ssl cert due to it being self signed
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	// Get the address and port from the terraform file and combine them to create the api url
	url := c.address + ":" + strconv.Itoa(c.port) + "/v1"

	// create a new session with the zerto cloud manager getting the x-zerto-session key back from the returned header
	client := &http.Client{}
	request, err := http.NewRequest("POST", url+"/session/add", nil)
	request.SetBasicAuth(c.username, c.password)
	response, err := client.Do(request)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	}

	// return the session key from the response header
	return url, response.Header.Get("x-zerto-session")
}
