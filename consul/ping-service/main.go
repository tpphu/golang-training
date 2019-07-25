package main

import (
	"fmt"
	"net/http"
	"os"

	consulapi "github.com/hashicorp/consul/api"
)

const port = 8080

func main() {
	// Step 1 - Create config
	config := consulapi.DefaultConfig()
	// Step 2 - New client
	client, err := consulapi.NewClient(config)
	if err != nil {
		panic(err)
	}
	// Step 3 - Create a service
	service := createPingService()
	// Step 4 - Register service to consul
	err = client.Agent().ServiceRegister(service)
	if err != nil {
		panic(err)
	}
	// Step 5 - Listen on HTTP Port
	http.HandleFunc("/ping", pingHandler)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}

func createPingService() *consulapi.AgentServiceRegistration {
	service := new(consulapi.AgentServiceRegistration)
	service.Name = "ping-service"
	service.Port = port
	service.Address, _ = os.Hostname()
	return service
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world!"))
}
