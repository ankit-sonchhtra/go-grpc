package main

import (
	"context"

	pb "github.com/ankit-sonchhtra/go-grpc/proto"
)

func (s *Server) Sum(ctx context.Context, request *pb.Request) (*pb.Response, error) {
	return &pb.Response{
		Result: request.FirstNumber + request.SecondNumber,
	}, nil
}
