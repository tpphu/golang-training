# Service

This is an example of creating a micro service.

## Contents

- service.go - is the main definition of the service
- client.go - is the main definition of the client
- proto - contains the protobuf definition of the API

## Run the example

- Run Consul docker

```shell
docker-compose up -d
```

- Run the service

```shell
go run service.go --registry=consul
```

- Run the client

```shell
go run client.go --registry=consul
```

And that's all there is to it.

## Generate Proto

```shell
protoc --go_out=. ./proto/greeter.proto
protoc --micro_out=. ./proto/greeter.proto
```

> If got some error please install:

- [protoc](https://github.com/google/protobuf)
- [protoc-gen-go](https://github.com/golang/protobuf)
- [protoc-gen-micro](github.com/micro/protoc-gen-micro)
