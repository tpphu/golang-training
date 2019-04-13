package main

import (
	"fmt"
	"log"

	consul "github.com/hashicorp/consul/api"
)

// Client provides an interface for getting data out of Consul
type Client interface {
	// Get a Service from consul
	Service(string, string) ([]*consul.ServiceEntry, *consul.QueryMeta, error)
	// Register a service with local agent
	Register(string, int) error
	// Deregister a service with local agent
	DeRegister(string) error
}

type client struct {
	consul *consul.Client
}

// NewConsulClient returns a Client interface for given consul address
func NewConsulClient(addr string) (Client, error) {
	config := consul.DefaultConfig()
	config.Address = addr
	c, err := consul.NewClient(config)
	if err != nil {
		return nil, err
	}
	return &client{consul: c}, nil
}

// Register a service with consul local agent
func (c *client) Register(name string, port int) error {
	reg := &consul.AgentServiceRegistration{
		ID:   name,
		Name: name,
		Port: port,
	}
	return c.consul.Agent().ServiceRegister(reg)
}

// DeRegister a service with consul local agent
func (c *client) DeRegister(id string) error {
	return c.consul.Agent().ServiceDeregister(id)
}

// Service return a service
func (c *client) Service(service, tag string) ([]*consul.ServiceEntry, *consul.QueryMeta, error) {
	passingOnly := true
	addrs, meta, err := c.consul.Health().Service(service, tag, passingOnly, nil)
	if len(addrs) == 0 && err == nil {
		return nil, nil, fmt.Errorf("service ( %s ) was not found", service)
	}
	if err != nil {
		return nil, nil, err
	}
	return addrs, meta, nil
}

func main() {
	serviceConsul, err := NewConsulClient("http://localhost:8500")
	if err != nil {
		log.Fatalln("Can't find consul:", err)
	}
	services, _, err := serviceConsul.Service("greeter", "")
	if err != nil {
		log.Fatalln("Discover failed:", err)
	}
	log.Println("Found service at these locations:")
	for _, v := range services {
		log.Println(fmt.Sprintf("%s:%d", v.Node.Address, v.Service.Port))
	}
}
