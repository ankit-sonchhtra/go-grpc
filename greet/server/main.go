package main

import (
	"log"
	"net"

	pb "github.com/ankit-sonchhtra/go-grpc/proto"
	"google.golang.org/grpc"
)

var addr string = "0.0.0.0:50052"

type Server struct {
	pb.GreetServiceServer
	pb.CalculatorServiceServer
	pb.PrimeCalculatorServiceServer
	pb.AvgCalculatorServiceServer
}

func main() {

	listner, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Failed to listen : %v\n", err)
	}

	log.Printf("Listing on %s\n", addr)

	s := grpc.NewServer()
	pb.RegisterGreetServiceServer(s, &Server{})
	pb.RegisterCalculatorServiceServer(s, &Server{})
	pb.RegisterPrimeCalculatorServiceServer(s, &Server{})
	pb.RegisterAvgCalculatorServiceServer(s, &Server{})

	if err := s.Serve(listner); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}

}
