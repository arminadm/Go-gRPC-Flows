package main

import (
	pb "Go-gRPC-Flows/gen"
	"context"
	"fmt"
	"log"
	"time"
)

func CallClientStreamingRPC(client pb.GreetServiceClient, names *pb.NameList) {
	stream, err := client.ClientStreamingRPC(context.Background())
	if err != nil {
		log.Fatalf("Failed to make client streaming with the server: %v", err)
	}
	for _, name := range names.Names {
		if err := stream.Send(&pb.HelloRequest{
			Name: name,
		}); err != nil {
			log.Fatalf("Failed to send %v during this error: %v", name, err)
		}
		time.Sleep(3 * time.Second)
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Failed to close streaming: %v", err)
	}
	fmt.Printf("received response after client streaming: %v", res)
}
