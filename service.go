package housinggrpc

// AnalysisRequest is the property rent/purchase request object
type AnalysisRequest struct {
	HomeValue         float32
	DownPayment       float32
	MonthlyRent       float32
	OccupancyDuration int
}

// AnalysisResponse is the break even analysic response object
type AnalysisResponse struct {
	Rent     float32
	Purchase float32
	Message  string
	Verdict  string
}

// Service defines the interface for the loan repayment computation.
type Service interface {
	ComputePropertyBreakEven(payload AnalysisRequest) (AnalysisResponse, error)
}
