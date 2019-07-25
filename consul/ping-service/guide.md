## Step 1 - Register a Service

```go
package main

import (
	"time"

	consulapi "github.com/hashicorp/consul/api"
)

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
	time.Sleep(10 * time.Second)
}

func createPingService() *consulapi.AgentServiceRegistration {
	service := new(consulapi.AgentServiceRegistration)
	service.Name = "ping-service"
	return service
}
```

## Step 2 - Understand healcheck
