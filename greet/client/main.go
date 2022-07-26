package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/ankit-sonchhtra/go-grpc/proto"
)

var addr string = "localhost:50052"

func main() {

	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect %v\n", err)
	}

	defer conn.Close()

	greetClient := pb.NewGreetServiceClient(conn)
	doGreet(greetClient)

	calcClient := pb.NewCalculatorServiceClient(conn)
	doSum(calcClient)

	doGreetManyTimes(greetClient)

	primeClient := pb.NewPrimeCalculatorServiceClient(conn)
	doPrime(primeClient)

	avgClient := pb.NewAvgCalculatorServiceClient(conn)
	doAvg(avgClient)

}
