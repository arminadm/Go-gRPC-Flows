package main

import (
	pb "Go-gRPC-Flows/gen"
	"io"
	"log"
)

func (s *MyGreetingServer) ClientStreamingRPC(stream pb.GreetService_ClientStreamingRPCServer) error {
	var messages []string
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.MessageList{
				Messages: messages,
			})
		}
		if err != nil {
			return err
		}
		log.Printf("got request from client: %v", req.Name)
		messages = append(messages, "Hello "+req.Name)
	}
}
