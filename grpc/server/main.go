package main

import (
	"fmt"
	"log"
	"net"

	secrets "github.com/oladotunsobande/housing-grpc/config"

	"github.com/oladotunsobande/housing-grpc/core"
	appgrpc "github.com/oladotunsobande/housing-grpc/grpc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// configure our core service
	calculatorService := core.NewService()

	// configure our gRPC service controller
	calulatorServiceController := NewCalculatorServiceController(calculatorService)

	// start a gRPC server
	server := grpc.NewServer()
	appgrpc.RegisterCalculatorServiceServer(server, calulatorServiceController)
	reflection.Register(server)

	con, err := net.Listen("tcp", fmt.Sprintf(":%s", secrets.GetSecrets().CalculatorServicePort))
	if err != nil {
		panic(err)
	}

	log.Printf("Calculator gRPC service running on %s...\n", con.Addr().String())
	err = server.Serve(con)
	if err != nil {
		panic(err)
	}
}
