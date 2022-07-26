package main

import (
	"fmt"
	"io"
	"log"

	pb "github.com/ankit-sonchhtra/go-grpc/proto"
)

func (s *Server) Avg(server pb.AvgCalculatorService_AvgServer) error {
	result := 0.0
	var sum int64 = 0
	var counter int64 = 0
	for {
		msg, err := server.Recv()

		if err == io.EOF {
			fmt.Println("Inside the err == io.EOF")
			result = float64(sum / counter)
			return server.SendAndClose(&pb.AvgResponse{
				Result: result,
			})
		}

		if err != nil {
			log.Fatalf("Error while sending the data over a stream : %v", err)
		}

		sum += msg.Number
		counter++

		fmt.Printf("Iteration : %v \n", sum)
		fmt.Printf("Counter : %v\n", counter)
	}
}
