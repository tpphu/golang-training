#!/bin/sh

protoc --go_out=plugins=grpc:. ./proto/note.proto
protoc --micro_out=. ./proto/note.proto