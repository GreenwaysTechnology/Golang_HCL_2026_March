package main

import (
	pb "grpcapp/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	lis, _ := net.Listen("tcp", ":50051")

	// Register Interceptors here
	s := grpc.NewServer(
		grpc.UnaryInterceptor(loggingUnaryInterceptor),
		grpc.StreamInterceptor(loggingStreamInterceptor),
	)

	pb.RegisterAppServiceServer(s, &server{})
	log.Println("Server started on :50051")
	s.Serve(lis)
}
