package housinggrpc

// LoanDetails is the mortgage repayment object
type LoanDetails struct {
	LoanAmount       float64
	InterestRate     float64
	NumberOfPayments int64
}

// Service defines the interface for the loan repayment computation.
type Service interface {
	ComputeMonthlyRepayment(loanAmount float64, interestRate float32, numberOfPayments int64) (float64, error)
}
