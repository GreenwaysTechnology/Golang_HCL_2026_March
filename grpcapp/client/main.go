package main

import (
	"context"
	"fmt"
	pb "grpcapp/proto"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// -------------------- Unary --------------------
func unary(client pb.AppServiceClient) {
	res, _ := client.SayHello(context.Background(), &pb.HelloRequest{Name: "Subramanian"})
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
