package main

import (
	"fmt"
	"grpc/greet/greetpb"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
}

func main() {
	fmt.Println("Hello, world")
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v ", err)
	}
	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
