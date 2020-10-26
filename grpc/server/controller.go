package main

import (
	"context"

	housinggrpc "github.com/oladotunsobande/housing-grpc"
	appgrpc "github.com/oladotunsobande/housing-grpc/grpc"
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
func (ctlr *calculatorServiceController) ComputeMonthlyRepayment(ctx context.Context, req *appgrpc.LoanDataRequest) (result *appgrpc.LoanDataResponse, err error) {
	result, err := ctlr.calculatorService.ComputeMonthlyRepayment(req.GetLoanAmount(), req.GetInterestRate(), req.GetNumberOfPayments())
	if err != nil {
		return
	}

	return
}
