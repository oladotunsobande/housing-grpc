package main

import (
	"log"
	"net"
	"os"

	appcore "github.com/oladotunsobande/housing-grpc/core"
	appgrpc "github.com/oladotunsobande/housing-grpc/grpc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// configure our core service
	calculatorService := appcore.NewService()

	// configure our gRPC service controller
	calulatorServiceController := NewCalculatorServiceController(calculatorService)

	// start a gRPC server
	server := grpc.NewServer()
	appgrpc.RegisterCalculatorServiceServer(server, calulatorServiceController)
	reflection.Register(server)

	con, err := net.Listen("tcp", os.Getenv("GRPC_ADDR"))
	if err != nil {
		panic(err)
	}

	log.Printf("Starting gRPC calculator service on %s...\n", con.Addr().String())
	err = server.Serve(con)
	if err != nil {
		panic(err)
	}
}
