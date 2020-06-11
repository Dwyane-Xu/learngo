package main

import (
	"context"
	"log"
	"net"
	"strings"
	"time"

	"github.com/docker/docker/pkg/pubsub"
	"google.golang.org/grpc"

	pb "learngo/goAdvance/ch4.4/3/pubsubservice"
)

type PubsubService struct {
	pub *pubsub.Publisher
}

func NewPubsubService() *PubsubService {
	return &PubsubService{
		pub: pubsub.NewPublisher(100*time.Millisecond, 10),
	}
}

func (p *PubsubService) Publish(
	ctx context.Context, in *pb.String) (*pb.String, error) {
	p.pub.Publish(in.GetValue())
	return &pb.String{}, nil
}

func (p *PubsubService) Subscribe(
	in *pb.String, stream pb.PubsubService_SubscribeServer) error {
	ch := p.pub.SubscribeTopic(func(v interface{}) bool {
		if key, ok := v.(string); ok {
			if strings.HasPrefix(key, in.GetValue()) {
				return true
			}
		}
		return false
	})

	for v := range ch {
		if err := stream.Send(&pb.String{Value: v.(string)}); err != nil {
			return err
		}
	}

	return nil
}

func main() {
	grpcServer := grpc.NewServer()
	pb.RegisterPubsubServiceServer(grpcServer, NewPubsubService())

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}

	grpcServer.Serve(listener)
}
