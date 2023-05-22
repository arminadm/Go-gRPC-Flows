package main

import (
	"fmt"
	"log"
	"net"

	pb "Go-gRPC-Flows/gen"

	"google.golang.org/grpc"
)

const port = ":8080"

type MyGreetingServer struct {
	pb.GreetServiceServer
}

func main() {
	// Listen on port 8080
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen (tcp, port:8080) : %v", err)
	}
	// Create gRPC server
	grpcServer := grpc.NewServer()
	// Registration
	pb.RegisterGreetServiceServer(grpcServer, &MyGreetingServer{})
	fmt.Printf("server registration passed and started at %v", lis.Addr())
	// serve the server
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve at %v, due to: %v", lis.Addr(), err)
	}
}
