package main

import (
	"context"
	"log"
	"time"

	pbV1 "github.com/username/grpcschemasorganization/gen/go/v1"
	pbV2 "github.com/username/grpcschemasorganization/gen/go/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// --- Test V1 ---
	clientV1 := pbV1.NewUserServiceClient(conn)
	resp1, err := clientV1.GetUser(ctx, &pbV1.UserRequest{UserId: "123"})
	if err != nil {
		log.Printf("V1 Error: %v", err)
	} else {
		log.Printf("V1 Response: %s (Username: %s)", resp1.UserId, resp1.Username)
	}

	// --- Test V2 ---
	clientV2 := pbV2.NewUserServiceClient(conn)
	resp2, err := clientV2.GetUser(ctx, &pbV2.UserRequest{UserId: "456"})
	if err != nil {
		log.Printf("V2 Error: %v", err)
	} else {
		log.Printf("V2 Response: %s (FullName: %s)", resp2.UserId, resp2.FullName)
	}
}
