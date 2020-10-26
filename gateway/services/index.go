package services

// BreakEvenPayload defines the object for break-even analysis
type BreakEvenPayload struct {
	PropertyValue      float32 `json:"propertyValue" validate:"min=1.00required"`
	InitialDepositRate float32 `json:"initialDepositRate" validate:"min=1.00,required"`
	MonthlyRent        float32 `json:"monthlyRent" validate:"min=1.00,required"`
	OccupancyDuration  int     `json:"occupancyDuration" validate:"min=1,required"`
}
