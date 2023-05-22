package main

import (
	pb "Go-gRPC-Flows/gen"
	"context"
	"fmt"
	"log"
	"time"
)

func CallSimpleRPC(client pb.GreetServiceClient) {
	// creating context
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// calling SimpleRPC in the server side
	res, err := client.SimpleRPC(ctx, &pb.NoParam{})
	if err != nil {
		log.Fatalf("error accrued during simple RPC call from client to the server: %v", err)
	}

	// print the response
	fmt.Printf("Client got this response from the server: \n%v\n", res)
	fmt.Println("==============================================")
	fmt.Printf("response message: %v\n", res.Message)
}
