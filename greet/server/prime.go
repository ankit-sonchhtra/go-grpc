package main

import (
	"log"

	pb "github.com/ankit-sonchhtra/go-grpc/proto"
)

func (s *Server) Prime(primeReq *pb.PrimeRequest, stream pb.PrimeCalculatorService_PrimeServer) error {
	log.Printf("Invoking prime with number : %d", primeReq.Number)
	number := primeReq.Number
	var i int64 = 2
	for number > 1 {
		if number%i == 0 {
			stream.Send(&pb.PrimeResponse{
				Result: i,
			})
			number = number / i
		} else {
			i = i + 1
		}
	}
	return nil
}
