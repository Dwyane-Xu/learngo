package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var port = ":5000"

type GreeterServiceImpl struct{}

func (g *GreeterServiceImpl) SayHello(
	ctx context.Context, in *HelloRequest) (*HelloReply, error) {
	return &HelloReply{Message: "Hello " + in.Name}, nil
}

func startServer() {
	creds, err := credentials.NewServerTLSFromFile("tls-config/server.crt", "tls-config/server.key")
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer(grpc.Creds(creds))
	RegisterGreeterServiceServer(grpcServer, new(GreeterServiceImpl))

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Println(w, "hello")
	})

	http.ListenAndServeTLS(port, "tls-config/server.crt", "tls-config/server.key",
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
				grpcServer.ServeHTTP(w, r)
			} else {
				mux.ServeHTTP(w, r)
			}
		}),
	)
}

func doClientWork() {
	creds, err := credentials.NewClientTLSFromFile("tls-config/server.crt", "server.grpc.io")
	if err != nil {
		log.Fatal(err)
	}

	conn, err := grpc.Dial(port, grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	c := NewGreeterServiceClient(conn)

	r, err := c.SayHello(context.Background(), &HelloRequest{Name: "gopher"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("doClientWork: %s", r.Message)
}

func main() {
	go startServer()
	time.Sleep(time.Second)

	doClientWork()

	time.Sleep(time.Second * 1000)
}
