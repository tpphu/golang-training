#!/bin/bash


# Generate proto

protoc \
-I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
-I. \
--go_out=plugins=grpc:. ./proto/note.proto

# Generate gateway

protoc -I/usr/local/include -I. \
-I$GOPATH/src \
-I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
--go_out=plugins=grpc:. \
./proto/note.proto
