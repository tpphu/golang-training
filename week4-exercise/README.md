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