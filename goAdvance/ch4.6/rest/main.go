package main

import (
	"context"
	"flag"
	"log"
	"net"
	"net/http"

	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

var (
	port         = ":5000"
	echoEndpoint = flag.String("echo_endpoint", "localhost"+port, "endpoint of YourService")
)

type myGrpcServer struct{}

func (s *myGrpcServer) Get(ctx context.Context, in *StringMessage) (*StringMessage, error) {
	return &StringMessage{Value: "Get: " + in.Value}, nil
}
func (s *myGrpcServer) Post(ctx context.Context, in *StringMessage) (*StringMessage, error) {
	return &StringMessage{Value: "Post: " + in.Value}, nil
}

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := RegisterRestServiceHandlerFromEndpoint(ctx, mux, *echoEndpoint, opts)
	if err != nil {
		return err
	}

	return http.ListenAndServe(":8080", mux)
}

func main() {
	flag.Parse()
	defer glog.Flush()

	go startGrpcServer()

	if err := run(); err != nil {
		glog.Fatal(err)
	}
}

func startGrpcServer() {
	server := grpc.NewServer()
	RegisterRestServiceServer(server, new(myGrpcServer))

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Panicf("could not list on %s: %s", port, err)
	}

	if err := server.Serve(lis); err != nil {
		log.Panicf("grpc serve error: %s", err)
	}
}
