package main

import (
	"context"
	"log"

	"github.com/oladotunsobande/housing-grpc/calculator"

	"google.golang.org/grpc"
)

func main() {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := calculator.NewCalculatorServiceClient(conn)

	response, err := c.ComputeBreakEven(context.Background(), &calculator.Message{Body: "Hello From Client!"})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Response from server: %s", response.Body)

}
