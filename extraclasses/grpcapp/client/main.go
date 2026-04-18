package main

import (
	"context"
	"fmt"
	pb "grpcapp/proto"
	"io"
	"log"
	"time"

	"google.golang.org/genproto/googleapis/rpc/errdetails" // For rich errors
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata" // For metadata
	"google.golang.org/grpc/status"   // For gRPC status
)

// -------------------- Unary --------------------
func unary(client pb.AppServiceClient) {
	// --- SENDING METADATA ---
	md := metadata.Pairs("authorization", "secret-token-123")
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	// Intentionally trigger an error by sending a short name
	// res, err := client.SayHello(ctx, &pb.HelloRequest{Name: "Su"})
	res, err := client.SayHello(ctx, &pb.HelloRequest{Name: "su"})

	if err != nil {
		// --- HANDLING RICH ERRORS ---
		st := status.Convert(err)
		log.Printf("gRPC Error: [%s] %s", st.Code(), st.Message())

		for _, detail := range st.Details() {
			switch t := detail.(type) {
			case *errdetails.BadRequest:
				for _, v := range t.GetFieldViolations() {
					log.Printf("Field Violation: %s -> %s", v.Field, v.Description)
				}
			}
		}
		return
	}
	log.Println("Unary Response:", res.Message)
}

// -------------------- Server Streaming --------------------
func serverStreaming(client pb.AppServiceClient) {
	stream, _ := client.GetNumbers(context.Background(), &pb.NumberRequest{Number: 5})

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		log.Println("Stream:", res.Number)
	}
}

// -------------------- Client Streaming --------------------
func clientStreaming(client pb.AppServiceClient) {
	stream, _ := client.UploadNumbers(context.Background())

	for i := 1; i <= 5; i++ {
		stream.Send(&pb.NumberRequest{Number: int32(i)})
	}

	res, _ := stream.CloseAndRecv()
	log.Println("Sum:", res.Number)
}

// -------------------- Bidirectional --------------------
func bidirectional(client pb.AppServiceClient) {
	stream, _ := client.Chat(context.Background())

	// send
	go func() {
		for i := 1; i <= 5; i++ {
			stream.Send(&pb.ChatMessage{Text: fmt.Sprintf("Msg %d", i)})
			time.Sleep(time.Second)
		}
	}()

	// receive
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		log.Println("Server:", res.Text)
	}
}
func main() {
	conn, err := grpc.NewClient("localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	client := pb.NewAppServiceClient(conn)
	//unary
	unary(client)
	//Server streaming
	serverStreaming(client)
	//Client Streaming
	clientStreaming(client)

	//Bidirectional streaming
	bidirectional(client)
}
