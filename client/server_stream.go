package main

import (
	pb "Go-gRPC-Flows/gen"
	"context"
	"io"
	"log"
)

func CallServerStreamingRPC(client pb.GreetServiceClient, names *pb.NameList) {
	log.Printf("Client Started the streaming request")
	stream, err := client.ServerStreamingRPC(context.Background(), names)
	if err != nil {
		log.Fatalf("Could not send names to server for streaming: %v", err)
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Failed to receive streaming response: %v", err)
		}
		log.Printf("Message received from server: %v", res.Message)
	}

	log.Printf("Streaming Finished")
}
