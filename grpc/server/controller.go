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

// ComputePropertyBreakEven calls the core service's ComputePropertyBreakEven method and maps the result to a grpc service response.
func (ctlr *calculatorServiceController) ComputePropertyBreakEven(ctx context.Context, req *appgrpc.BreakEvenRequest) (result *appgrpc.BreakEvenResponse, err error) {
	payload := housinggrpc.AnalysisRequest{
		HomeValue:         req.GetHomeValue(),
		DownPayment:       req.GetDownPayment(),
		MonthlyRent:       req.GetMonthlyRent(),
		OccupancyDuration: int(req.GetOccupancyDuration()),
	}

	res, err := ctlr.calculatorService.ComputePropertyBreakEven(payload)
	if err != nil {
		return
	}

	result = &appgrpc.BreakEvenResponse{
		Rent:     res.Rent,
		Purchase: res.Purchase,
		Message:  res.Message,
		Verdict:  res.Verdict,
	}

	return
}
