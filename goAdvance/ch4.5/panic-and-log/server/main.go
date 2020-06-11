package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	hw "learngo/goAdvance/ch4.5/panic-and-log/helloworld"
)

type GreeterServiceImpl struct{}

func (g *GreeterServiceImpl) SayHello(
	ctx context.Context, in *hw.HelloRequest) (*hw.HelloReply, error) {
	panic("debug")
	return &hw.HelloReply{Message: "Hello " + in.Name}, nil
}

func filter(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler,
) (resp interface{}, err error) {
	log.Println("filter:", info)

	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic: %v", r)
		}
	}()

	return handler(ctx, req)
}

func main() {
	server := grpc.NewServer(grpc.UnaryInterceptor(filter))
	hw.RegisterGreeterServiceServer(server, new(GreeterServiceImpl))

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Panicf("could not list on 1234 %s", err)
	}

	if err := server.Serve(listener); err != nil {
		log.Panicf("grpc serve error: %s", err)
	}
}
