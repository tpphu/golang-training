## Generate proto

```shell
protoc --go_out=. ./proto/note.proto
protoc --go_out=plugins=grpc:. ./proto/note.proto
```