package main

import (
	"context"
	"fmt"
	greeter "grpc_gateway/generated"
	"grpc_gateway/implementations"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

const (
	grpcAddr = "0.0.0.0:29090"
	httpGwAddr = "0.0.0.0:8081"
)

func startGrpcServer() {
	server := grpc.NewServer()
	// register all implementations
	greeter.RegisterGreeterServer(server, &implementations.Greeter{})
	lis, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		panic(err)
	}

	fmt.Printf("starting grpc service on addr %s\n", grpcAddr)

	err = server.Serve(lis)
	if err != nil {
		panic(err)
	}
}

func startHttpGateway() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	// register all handlers
	err := greeter.RegisterGreeterHandlerFromEndpoint(ctx, mux, grpcAddr, opts)
	if err != nil {
		panic(err)
	}

	fmt.Printf("starting http gateway service on addr %s\n", httpGwAddr)

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	http.ListenAndServe(httpGwAddr, mux)
}

func main() {
	go startGrpcServer()
	startHttpGateway()
}
