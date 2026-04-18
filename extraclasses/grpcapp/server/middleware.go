package main

import (
	"context"
	"log"

	"google.golang.org/grpc"
)

// Unary Interceptor
func loggingUnaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log.Printf("--> Unary Call: %s", info.FullMethod)
	resp, err := handler(ctx, req) // Execute the handler
	log.Printf("<-- Unary Finished")
	return resp, err
}

// Stream Interceptor
func loggingStreamInterceptor(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	log.Printf("--> Stream Call: %s", info.FullMethod)
	err := handler(srv, ss)
	log.Printf("<-- Stream Finished")
	return err
}
