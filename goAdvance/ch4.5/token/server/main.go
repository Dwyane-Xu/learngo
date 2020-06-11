package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	hw "learngo/goAdvance/ch4.5/token/helloworld"
)

type GreeterServiceImpl struct{}

func (g *GreeterServiceImpl) SayHello(
	ctx context.Context, in *hw.HelloRequest) (*hw.HelloReply, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("missing credentials")
	}

	var (
		login    string
		password string
	)

	if val, ok := md["login"]; ok {
		login = val[0]
	}
	if val, ok := md["password"]; ok {
		password = val[0]
	}

	if login != "gopher" || password != "password" {
		return nil, status.Errorf(codes.Unauthenticated,
			"invalid token: login=%s, password=%s", login, password)
	}

	return &hw.HelloReply{Message: "Hello " + in.Name}, nil
}

func main() {
	server := grpc.NewServer()
	hw.RegisterGreeterServiceServer(server, new(GreeterServiceImpl))

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Panicf("could not list on 1234 %s", err)
	}

	if err := server.Serve(listener); err != nil {
		log.Panicf("grpc serve error: %s", err)
	}
}
