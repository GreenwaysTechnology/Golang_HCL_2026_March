package main

import (
	"context"
	"fmt"
	pb "grpcapp/proto"
	"io"
	"log"
	"time"

	"google.golang.org/genproto/googleapis/rpc/errdetails" // For rich errors
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata" // For metadata
	"google.golang.org/grpc/status"   // For gRPC status
)

type server struct {
	pb.UnimplementedAppServiceServer
}

func printMetadata(ctx context.Context, methodName string) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		log.Printf("[%s] No metadata found", methodName)
		return
	}

	fmt.Printf("\n--- Metadata Received in %s ---\n", methodName)
	for key, values := range md {
		for _, value := range values {
			fmt.Printf("%s: %s\n", key, value)
		}
	}
	fmt.Println("------------------------------------")
}

// 1️⃣ Unary
// 1️⃣ Unary with Metadata and Rich Errors
func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	// --- RECEIVING METADATA ---
	// md, ok := metadata.FromIncomingContext(ctx)
	// if ok {
	// 	log.Printf("Received Metadata: %v", md["authorization"])
	// }
	printMetadata(ctx, "SayHello")
	// --- RICH ERROR HANDLING ---
	if len(req.Name) < 3 {
		st := status.New(codes.InvalidArgument, "Name is too short")

		// Add structured details to the error
		v := &errdetails.BadRequest{
			FieldViolations: []*errdetails.BadRequest_FieldViolation{
				{Field: "name", Description: "Should be at least 3 characters"},
			},
		}

		st, _ = st.WithDetails(v)
		return nil, st.Err()
	}

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
