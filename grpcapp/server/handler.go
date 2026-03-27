package main

import (
	"context"
	pb "grpcapp/proto"
	"io"
	"log"
	"time"
)

type server struct {
	pb.UnimplementedAppServiceServer
}

// 1️⃣ Unary
func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	log.Println("Unary:", req.Name)
	return &pb.HelloResponse{Message: "Hello " + req.Name}, nil
}

// 2️⃣ Server Streaming
func (s *server) GetNumbers(req *pb.NumberRequest, stream pb.AppService_GetNumbersServer) error {
	for i := 1; i <= int(req.Number); i++ {
		stream.Send(&pb.NumberResponse{Number: int32(i)})
		time.Sleep(time.Millisecond * 500)
	}
	return nil
}

// 3️⃣ Client Streaming
func (s *server) UploadNumbers(stream pb.AppService_UploadNumbersServer) error {
	sum := int32(0)

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.NumberResponse{Number: sum})
		}
		sum += req.Number
	}
}

// 4️⃣ Bidirectional Streaming
func (s *server) Chat(stream pb.AppService_ChatServer) error {
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			return nil
		}

		log.Println("Client:", msg.Text)

		stream.Send(&pb.ChatMessage{
			Text: "Echo: " + msg.Text,
		})
	}
}
