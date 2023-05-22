package main

import (
	pb "Go-gRPC-Flows/gen"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const port = ":8080"

func main() {
	// create connection to the server
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("grpc dial connection failed: %v", err)
	}
	defer conn.Close()

	// create a new client
	client := pb.NewGreetServiceClient(conn)

	// create list of names as request body
	// names := &pb.NameList{
	// 	Names: []string{"Armin", "Mehrdad", "Asal", "Ali"},
	// }

	// calling the Flows, un comment each flow that you need and comment the rest
	CallSimpleRPC(client) // Simple gRPC

}
