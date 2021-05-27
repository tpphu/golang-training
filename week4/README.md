https://grpc.io/
https://developers.google.com/protocol-buffers/docs/overview
https://github.com/grpc-ecosystem/grpc-gateway    
https://github.com/namely/docker-protoc


protoc --go_out=. --go_opt=paths=source_relative \
--go-grpc_out=. --go-grpc_opt=paths=source_relative \
 week4/internal/api/patient.proto