package main

import (
	"context"
	"log"

	pb "github.com/ankit-sonchhtra/go-grpc/proto"
)

func (s *Server) Greet(ctx context.Context, request *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("Greet was invoked %v \n", request)
	return &pb.GreetResponse{
		Result: "Hello " + request.FirstName,
	}, nil
}
