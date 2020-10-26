package housinggrpc

// LoanResponse is the mortgage monthly repayment object
type LoanResponse struct {
	Repayment float64
}

// Service defines the interface for the loan repayment computation.
type Service interface {
	ComputeMonthlyRepayment(loanAmount float64, interestRate float64, numberOfPayments int64) (LoanResponse, error)
}
