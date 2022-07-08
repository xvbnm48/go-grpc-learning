package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/xvbnm48/go-grpc-learning/greet/greetpb"

	"google.golang.org/grpc"
)

type server struct{}

func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	fmt.Printf("Greet function was invoked with %v", req)
	firtsname := req.GetGreeting().GetFirstName()
	result := "Hello " + firtsname
	response := &greetpb.GreetResponse{
		Result: result,
	}

	return response, nil
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
