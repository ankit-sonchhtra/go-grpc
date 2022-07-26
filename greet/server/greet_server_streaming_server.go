package main

import (
	"fmt"

	pb "github.com/ankit-sonchhtra/go-grpc/proto"
)

func (s *Server) GreetMany(in *pb.GreetRequest, stream pb.GreetService_GreetManyServer) error {
	for i := 0; i < 10; i++ {
		res := fmt.Sprintf("Hello %s, ivonking %d time", in.FirstName, i)
		stream.Send(&pb.GreetResponse{
			Result: res,
		})
	}
	return nil
}
