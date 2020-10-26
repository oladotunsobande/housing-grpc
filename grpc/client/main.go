package main

import (
	"context"
	"log"

	appgrpc "github.com/oladotunsobande/housing-grpc/grpc"
	"google.golang.org/grpc"
)

func main() {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":7990", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := appgrpc.NewCalculatorServiceClient(conn)

	response, err := c.ComputePropertyBreakEven(context.Background(), &appgrpc.BreakEvenRequest{HomeValue: 300000, DownPayment: 20, MonthlyRent: 3000, OccupancyDuration: 3})
	if err != nil {
		log.Fatalf("Error when calling ComputeMonthlyRepayment: %s", err)
	}
	log.Println("gRPC Response...\n")
	log.Printf("Rent:     $%d\n", response.Rent)
	log.Printf("Purchase: $%d\n", response.Purchase)
	log.Printf("Message:  %s\n", response.Message)
	log.Printf("Verdict:  %s\n", response.Verdict)
}
