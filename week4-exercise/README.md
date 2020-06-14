# Cach thong thuong khong xu dung cai GRPC Gateway
protoc --go_out=plugins=grpc:. ./proto/note.proto

# Doi voi su dung GRPC Gateway thi minh chay 2 cau lenh sau

## Cau 1st: la de generate note.pb.go

### Truong hop simple
protoc \
-I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
-I. \
--go_out=plugins=grpc:. ./proto/note.proto

### Truong hop full
protoc -I/usr/local/include -I. \
-I$GOPATH/src \
-I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
--go_out=plugins=grpc:. \
./proto/note.proto


## Cau 2nd: la de generate file gateway: note.pb.gw.go
protoc -I/usr/local/include -I. \
-I$GOPATH/src \
-I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
--grpc-gateway_out=logtostderr=true:. \
./proto/note.proto

## Cau 3nd:

docker run --rm -v $(pwd):$(pwd) -w $(pwd) znly/protoc --go_out=plugins=grpc:. -I. proto/note.proto

docker run --rm -v $(pwd):$(pwd) -w $(pwd) znly/protoc --grpc-gateway_out=logtostderr=true:. -I. proto/note.proto

## Build

```shell
$ go build -o server_default server/main.go
$ go build -o client_default client/main.go
```

## Gateway

```
curl -XDELETE http://localhost:8080/note/123
```

