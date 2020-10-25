package calculator

import (
	"math"

	housinggrpc "github.com/oladotunsobande/housing-grpc"
)

// ComputeMonthlyRepayment is the grpc service that computes the monthly loan repayment
func ComputeMonthlyRepayment(loanAmount float64, interestRate float64, numberOfPayments int64) float64 {
	payload := &housinggrpc.LoanDetails{
		LoanAmount:       loanAmount,
		InterestRate:     interestRate,
		NumberOfPayments: numberOfPayments,
	}

	p := payload.LoanAmount
	r := payload.InterestRate / 100
	n := payload.NumberOfPayments

	result := p * ((r * (math.Pow((1.00 + r), float64(n)))) / (math.Pow((1.00+r), float64(n)) - 1))

	return result
}
