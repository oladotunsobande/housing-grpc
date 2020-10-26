package services

import (
	"context"
	"fmt"
	"log"

	appgrpc "github.com/oladotunsobande/housing-grpc/grpc"
	"google.golang.org/grpc"
)

var gRPCClient appgrpc.CalculatorServiceClient

func init() {
	var conn *grpc.ClientConn

	conn, err := grpc.Dial(":7990", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Connection to calculator gRPC service failed: %s", err)
	}

	defer conn.Close()

	gRPCClient = appgrpc.NewCalculatorServiceClient(conn)
}

// CalculateBreakEven is the gRPC client that calls the calculator service
func (payload BreakEvenPayload) CalculateBreakEven() (interface{}, error) {
	requestConstruct := &appgrpc.BreakEvenRequest{
		HomeValue:         payload.PropertyValue,
		DownPayment:       payload.InitialDepositRate,
		MonthlyRent:       payload.MonthlyRent,
		OccupancyDuration: int32(payload.OccupancyDuration),
	}

	response, err := gRPCClient.ComputePropertyBreakEven(context.Background(), requestConstruct)
	if err != nil {
		fmt.Printf("Error when calling ComputeMonthlyRepayment: %s", err)
		return nil, err
	}

	return response, err
}
