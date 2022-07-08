package main

import (
	"fmt"
	"grpc/greet/greetpb"
	"log"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello, this is client")
	cc, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatal("could not connect to server: %v", err)
	}
	defer cc.Close()

	c := greetpb.NewGreetServiceClient(cc)
	fmt.Printf("Created client: %v\n", c)
}
