package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	// Replace with your actual module path from go.mod
	pb "github.com/youruser/grpc-buf-example/gen/greet/v1"
)

type server struct {
	pb.UnimplementedGreetServiceServer
}

func (s *server) SayHello(ctx context.Context, in *pb.SayHelloRequest) (*pb.SayHelloResponse, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.SayHelloResponse{Greeting: "Hello " + in.GetName()}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterGreetServiceServer(s, &server{})

	log.Println("Server listening on :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}