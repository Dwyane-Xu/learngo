package main

import (
	"context"
	"log"

	"google.golang.org/grpc"

	hw "learngo/goAdvance/ch4.5/token/helloworld"
)

type Authentication struct {
	Login    string
	Password string
}

func (a *Authentication) GetRequestMetadata(context.Context, ...string) (map[string]string, error) {
	return map[string]string{"login": a.Login, "password": a.Password}, nil
}

func (a *Authentication) RequireTransportSecurity() bool {
	return false
}

func main() {
	auth := &Authentication{
		Login:    "gopher",
		Password: "password",
	}

	conn, err := grpc.Dial(":1234", grpc.WithInsecure(), grpc.WithPerRPCCredentials(auth))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	c := hw.NewGreeterServiceClient(conn)

	reply, err := c.SayHello(context.Background(), &hw.HelloRequest{Name: "gopher"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("doClientWork: %s", reply.Message)
}
