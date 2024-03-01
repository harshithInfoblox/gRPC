package main

import (
	"context"
	"log"

	pb "grpc_example/pb" // Update the import path here

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewEchoServiceClient(conn)

	message := "Hello, Harshith!"
	response, err := c.Echo(context.Background(), &pb.EchoRequest{Message: message})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Response from server: %s", response.EchoedMessage)
}
