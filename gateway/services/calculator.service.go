package services

import (
	"context"
	"fmt"
	"log"

	"github.com/oladotunsobande/housing-grpc/config"
	appgrpc "github.com/oladotunsobande/housing-grpc/grpc"
	"google.golang.org/grpc"
)

// CalculateBreakEven is the gRPC client that calls the calculator service
func (payload BreakEvenPayload) CalculateBreakEven() (interface{}, error) {
	requestConstruct := &appgrpc.BreakEvenRequest{
		HomeValue:         payload.PropertyValue,
		DownPayment:       payload.InitialDepositRate,
		MonthlyRent:       payload.MonthlyRent,
		OccupancyDuration: int32(payload.OccupancyDuration),
	}

	conn, err := grpc.Dial(fmt.Sprintf(":%s", config.GetSecrets().CalculatorServicePort), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Connection to calculator gRPC service failed: %s", err)
	}

	gRPCClient := appgrpc.NewCalculatorServiceClient(conn)

	response, err := gRPCClient.ComputePropertyBreakEven(context.Background(), requestConstruct)
	if err != nil {
		fmt.Printf("Error when calling ComputeMonthlyRepayment: %s", err)
		return nil, err
	}

	defer conn.Close()

	return response, err
}
