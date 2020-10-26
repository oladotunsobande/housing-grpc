package core

import (
	"math"

	housinggrpc "github.com/oladotunsobande/housing-grpc"
)

type service struct {
	data housinggrpc.LoanResponse
}

// NewService instantiates a new Service.
func NewService() housinggrpc.Service {
	return &service{
		data: housinggrpc.LoanResponse{},
	}
}

// ComputeMonthlyRepayment is the grpc service that computes the monthly loan repayment
func (s *service) ComputeMonthlyRepayment(loanAmount float64, interestRate float64, numberOfPayments int64) (housinggrpc.LoanResponse, error) {
	p := loanAmount
	r := interestRate / 100
	n := numberOfPayments

	result := p * ((r * (math.Pow((1.00 + r), float64(n)))) / (math.Pow((1.00+r), float64(n)) - 1))

	return housinggrpc.LoanResponse{Repayment: result}, nil
}
