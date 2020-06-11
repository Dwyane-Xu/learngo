package main

import (
	"context"
	"io"
	"log"
	"net"

	"google.golang.org/grpc"

	hs "learngo/goAdvance/ch4.4/2/helloservice"
)

type HelloServiceImpl struct{}

func (p *HelloServiceImpl) Hello(
	ctx context.Context, in *hs.String) (*hs.String, error) {
	reply := &hs.String{Value: "hello:" + in.GetValue()}
	return reply, nil
}

func (p *HelloServiceImpl) Channel(stream hs.HelloService_ChannelServer) error {
	for {
		in, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}

		reply := &hs.String{Value: "hello:" + in.GetValue()}

		err = stream.Send(reply)
		if err != nil {
			return err
		}
	}
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
