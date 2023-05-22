package main

import (
	pb "Go-gRPC-Flows/gen"
	"context"
)

func (s *MyGreetingServer) SimpleRPC(ctx context.Context, req *pb.NoParam) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Message: "Hello, From the server",
	}, nil
}
