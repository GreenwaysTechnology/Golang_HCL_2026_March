package main

import (
	"context"
	"log"
	"net"

	// Import generated code with aliases
	pbV1 "github.com/username/grpcschemasorganization/gen/go/v1"
	pbV2 "github.com/username/grpcschemasorganization/gen/go/v2"
	"google.golang.org/grpc"
)

// --- V1 Implementation ---
type serverV1 struct {
	pbV1.UnimplementedUserServiceServer
}

func (s *serverV1) GetUser(ctx context.Context, req *pbV1.UserRequest) (*pbV1.UserResponse, error) {
	log.Printf("V1 Request Received: %s", req.GetUserId())
	return &pbV1.UserResponse{
		UserId:   req.GetUserId(),
		Username: "old_gopher_v1",
	}, nil
}

// --- V2 Implementation ---
type serverV2 struct {
	pbV2.UnimplementedUserServiceServer
}

func (s *serverV2) GetUser(ctx context.Context, req *pbV2.UserRequest) (*pbV2.UserResponse, error) {
	log.Printf("V2 Request Received: %s", req.GetUserId())
	return &pbV2.UserResponse{
		UserId:   req.GetUserId(),
		FullName: "Gopher The Second", // Matching the V2 field name
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	// Register separate instances for each version
	pbV1.RegisterUserServiceServer(s, &serverV1{})
	pbV2.RegisterUserServiceServer(s, &serverV2{})

	log.Println("Server listening on :50051 (supporting v1 and v2)")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}