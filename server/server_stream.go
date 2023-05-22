package main

import (
	pb "Go-gRPC-Flows/gen"
	"log"
	"time"
)

func (s *MyGreetingServer) ServerStreamingRPC(req *pb.NameList, stream pb.GreetService_ServerStreamingRPCServer) error {
	log.Printf("Server received names:\n%v", req.Names)
	for _, name := range req.Names {
		err := stream.Send(&pb.HelloResponse{
			Message: "Hello " + name,
		})
		if err != nil {
			return err
		}
		time.Sleep(2 * time.Second)
	}
	return nil
}
