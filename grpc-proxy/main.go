package main

import (
	"net"
	"strings"

	"github.com/mwitkow/grpc-proxy/proxy"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
)

var (
	director proxy.StreamDirector
)

func main() {
	director = func(ctx context.Context, fullMethodName string) (context.Context, *grpc.ClientConn, error) {
		// Make sure we never forward internal services.
		if strings.HasPrefix(fullMethodName, "/com.example.internal.") {
			return nil, nil, grpc.Errorf(codes.Unimplemented, "Unknown method")
		}
		md, ok := metadata.FromIncomingContext(ctx)
		// Copy the inbound metadata explicitly.
		outCtx, _ := context.WithCancel(ctx)
		outCtx = metadata.NewIncomingContext(outCtx, md.Copy())
		if ok {
			// Decide on which backend to dial
			if val, exists := md[":authority"]; exists && val[0] == "localhost:50050" {
				// Make sure we use DialContext so the dialing can be cancelled/time out together with the context.
				conn, err := grpc.DialContext(ctx, "localhost:50052", grpc.WithCodec(proxy.Codec()), grpc.WithInsecure())
				return outCtx, conn, err
			} else if val, exists := md[":authority"]; exists && val[0] == "api.example.com" {
				conn, err := grpc.DialContext(ctx, "api-service.prod.svc.local", grpc.WithCodec(proxy.Codec()))
				return outCtx, conn, err
			}
		}
		return nil, nil, grpc.Errorf(codes.Unimplemented, "Unknown method")
	}
	server := grpc.NewServer(
		grpc.CustomCodec(proxy.Codec()),
		grpc.UnknownServiceHandler(proxy.TransparentHandler(director)))
	lis, _ := net.Listen("tcp", ":50050")
	server.Serve(lis)

}
