package main

import (
	"context"
	"log"

	appgrpc "github.com/neocortical/housing-grpc/grpc"
	housinggrpc "github.com/oladotunsobande/housing-grpc"
)

// calculatorServiceController implements the gRPC CalculatorServiceServer interface.
type calculatorServiceController struct {
	calculatorService housinggrpc.Service
}

// NewCalculatorServiceController instantiates a new CalculatorServiceServer.
func NewCalculatorServiceController(calculatorService housinggrpc.Service) appgrpc.CalculatorServiceServer {
	return &calculatorServiceController{
		calculatorService: calculatorService,
	}
}

// ComputeMonthlyRepayment calls the core service's ComputeMonthlyRepayment method and maps the result to a grpc service response.
func (ctlr *calculatorServiceController) ComputeMonthlyRepayment(ctx context.Context, req *appgrpc.LoanDataRequest) (resp *appgrpc.LoanDataResponse, err error) {
	resultMap, err := ctlr.calculatorService.ComputeMonthlyRepayment(req.GetIds())
	if err != nil {
		return
	}

	resp = &mysvcgrpc.GetUsersResponse{}
	for _, u := range resultMap {
		resp.Users = append(resp.Users, marshalUser(&u))
	}

	log.Printf("handled GetUsers(%v)\n", req.GetIds())
	return
}
