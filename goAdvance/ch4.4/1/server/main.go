package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	hs "learngo/goAdvance/ch4.4/1/helloservice"
)

type HelloServiceImpl struct{}

func (p *HelloServiceImpl) Hello(
	ctx context.Context, in *hs.String) (*hs.String, error) {
	reply := &hs.String{Value: "hello:" + in.GetValue()}
	return reply, nil
}

func main() {
	grpcServer := grpc.NewServer()
	hs.RegisterHelloServiceServer(grpcServer, new(HelloServiceImpl))

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}
	grpcServer.Serve(listener)
}
