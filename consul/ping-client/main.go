package main

import (
	"fmt"

	consulapi "github.com/hashicorp/consul/api"
	consulconnect "github.com/hashicorp/consul/connect"
)

func main() {
	// Create a Consul API client
	client, err := consulapi.NewClient(consulapi.DefaultConfig())

	if err != nil {
		panic(err)
	}

	// Create an instance representing this service. "my-service" is the
	// name of _this_ service. The service should be cleaned up via Close.
	svc, _ := consulconnect.NewService("ping-service", client)
	defer svc.Close()

	// Get an HTTP client
	httpClient := svc.HTTPClient()

	// Perform a request, then use the standard response
	resp, _ := httpClient.Get("http://ping-service.service.consul:8080/ping")

	fmt.Println("la sao", resp)
}
