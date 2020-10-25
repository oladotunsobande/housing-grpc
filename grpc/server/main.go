package main

import (
	"fmt"
	"log"
	"net"

	"github.com/oladotunsobande/housing-grpc/src/services/calculator"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9000))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	fmt.Println("gRPC server running...")

	s := calculator.Server{}

	grpcServer := grpc.NewServer()

	calculator.RegisterCalculatorServiceServer(grpcServer, CalculatorServiceServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
