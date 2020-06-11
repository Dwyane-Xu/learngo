package main

import (
	"context"
	"log"

	"google.golang.org/grpc"

	hw "learngo/goAdvance/ch4.5/panic-and-log/helloworld"
)

func main() {
	conn, err := grpc.Dial(":1234", grpc.WithInsecure())
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
