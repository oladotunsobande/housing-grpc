package core

import (
	"fmt"
	"math"
	"strconv"

	housinggrpc "github.com/oladotunsobande/housing-grpc"
	secrets "github.com/oladotunsobande/housing-grpc/config"
)

var (
	cummulativePurchaseExpense = make(map[int]Expense)
	cummulativeRentExpense     = make(map[int]Expense)
)

type service struct {
	data housinggrpc.AnalysisResponse
}

// LoanDetails is the mortgage monthly repayment request object
type LoanDetails struct {
	LoanAmount       float32
	InterestRate     float32
	NumberOfPayments int32
}

// AnalysisVariables defines the estimation parameters object
type AnalysisVariables struct {
	HomeValue                         float32
	DownPayment                       float32
	PropertyAppreciationRate          float32
	AnnualInflation                   float32
	ClosingCost                       float32
	RealtorFees                       float32
	LoanDuration                      int32
	LoanInterestRate                  float32
	AnnualRentIncreaseRate            float32
	PropertyTax                       float32
	PropertyTransferTax               float32
	MaintenanceRate                   float32
	UtilityRate                       float32
	HomeOwnerAssociationRate          float32
	HomeOwnerAssociationInsuranceRate float32
	AnnualRepayment                   float32
	SecurityDeposit                   float32
}

// Expense defines the annual and cummlative expense object
type Expense struct {
	Annual      float32
	Cummulative float32
}

// NewService instantiates a new Service.
func NewService() housinggrpc.Service {
	return &service{
		data: housinggrpc.AnalysisResponse{},
	}
}

func convertInt(val string) int32 {
	value, _ := strconv.ParseInt(val, 10, 32)
	result := int32(value)

	return result
}

func convertFloat(val string) float32 {
	value, _ := strconv.ParseFloat(val, 32)
	result := float32(value)

	return result
}

// initVariables initializes the estimation parameters. This would be replaced by reading these parameters from a database
func initVariables(payload housinggrpc.AnalysisRequest) (analysisVariables AnalysisVariables, err error) {
	details := LoanDetails{
		LoanAmount:       payload.HomeValue - ((payload.DownPayment / 100) * payload.HomeValue),
		InterestRate:     (convertFloat(secrets.GetSecrets().LoanInterestRate) / 100) / 12,
		NumberOfPayments: convertInt(secrets.GetSecrets().LoanDuration) * 12,
	}

	monthRepayment, err := details.computeMonthlyRepayment()
	if err != nil {
		return
	}

	analysisVariables = AnalysisVariables{
		HomeValue:                         payload.HomeValue,
		DownPayment:                       (payload.DownPayment / 100) * payload.HomeValue,
		PropertyAppreciationRate:          convertFloat(secrets.GetSecrets().PropertyAppreciationRate) / 100,
		AnnualInflation:                   convertFloat(secrets.GetSecrets().AnnualInflation) / 100,
		ClosingCost:                       convertFloat(secrets.GetSecrets().ClosingCost) / 100,
		RealtorFees:                       convertFloat(secrets.GetSecrets().RealtorFees) / 100,
		LoanDuration:                      convertInt(secrets.GetSecrets().LoanDuration),
		LoanInterestRate:                  convertFloat(secrets.GetSecrets().LoanInterestRate) / 100,
		AnnualRentIncreaseRate:            convertFloat(secrets.GetSecrets().AnnualRentIncreaseRate) / 100,
		PropertyTax:                       convertFloat(secrets.GetSecrets().PropertyTax) / 100,
		PropertyTransferTax:               convertFloat(secrets.GetSecrets().PropertyTransferTax) / 100,
		MaintenanceRate:                   convertFloat(secrets.GetSecrets().MaintenanceRate) / 100,
		UtilityRate:                       (convertFloat(secrets.GetSecrets().UtilityRate) / 100) * 12,
		HomeOwnerAssociationRate:          convertFloat(secrets.GetSecrets().HomeOwnerAssociationRate) / 100,
		HomeOwnerAssociationInsuranceRate: convertFloat(secrets.GetSecrets().HomeOwnerAssociationInsuranceRate) / 100,
		AnnualRepayment:                   monthRepayment * 12,
		SecurityDeposit:                   payload.MonthlyRent * 3, // Security deposit is equivalent to 3 months rent
	}

	return
}

func analyzeResult(occupancyDuration int) (float32, float32, string, string) {
	var rent float32 = 0.00000000
	var purchase float32 = 0.00000000
	var message, verdict string

	expenseMapSize := len(cummulativePurchaseExpense)

	if occupancyDuration > 0 {
		rentExpense := cummulativeRentExpense[occupancyDuration]
		purchaseExpense := cummulativePurchaseExpense[occupancyDuration]

		rent = rentExpense.Cummulative
		purchase = purchaseExpense.Cummulative

		if purchaseExpense.Cummulative > rentExpense.Cummulative {
			message = fmt.Sprintf("If you stay in this property for %d years, renting is cheaper than buying.", occupancyDuration)
		} else if purchaseExpense.Cummulative < rentExpense.Cummulative {
			message = fmt.Sprintf("If you stay in this property for %d years, buying is cheaper than renting.", occupancyDuration)
		} else {
			message = fmt.Sprintf("If you stay in this property for %d years, there is no difference.", occupancyDuration)
		}
	}

	i := 0
	breakEvenYear := 0
	for ; i <= expenseMapSize; i++ {
		if cummulativePurchaseExpense[i].Cummulative < cummulativeRentExpense[i].Cummulative {
			breakEvenYear = i
			break
		}
	}

	if breakEvenYear > 0 {
		verdict = fmt.Sprintf("Year %d is the break-even year", breakEvenYear)
	} else {
		verdict = "There is no point of balance"
	}

	return rent, purchase, verdict, message
}

func setExpense(expenseMap map[int]Expense, year int, sum float32) {
	if year == 0 {
		expenseMap[year] = Expense{
			Annual:      sum,
			Cummulative: sum,
		}
	} else {
		previous := expenseMap[year-1]

		expenseMap[year] = Expense{
			Annual:      sum,
			Cummulative: previous.Cummulative + sum,
		}
	}
}

// computeMonthlyRepayment computes the monthly loan repayment
func (payload LoanDetails) computeMonthlyRepayment() (float32, error) {
	p := payload.LoanAmount
	r := payload.InterestRate
	n := payload.NumberOfPayments

	result := float64(p) * ((float64(r) * (math.Pow(float64(1.00+r), float64(n)))) / (math.Pow(float64(1.00+r), float64(n)) - 1))

	return float32(result), nil
}

// ComputeComputePropertyBreakEven is the implementation of the Service interface function for the gRPC service
func (s *service) ComputePropertyBreakEven(payload housinggrpc.AnalysisRequest) (resp housinggrpc.AnalysisResponse, err error) {
	analysisVariables, err := initVariables(payload)
	if err != nil {
		return
	}

	var verdict string

	i := 0
	propertyValue := payload.HomeValue
	monthlyRent := payload.MonthlyRent
	loanDuration := analysisVariables.LoanDuration

	for ; i <= int(loanDuration); i++ {
		if i > 1 {
			propertyValue += (propertyValue * analysisVariables.PropertyAppreciationRate)
			monthlyRent += ((monthlyRent * analysisVariables.AnnualInflation) + (monthlyRent * analysisVariables.AnnualRentIncreaseRate))
		}

		// Compute purchase expenses
		analysisVariables.computePurchaseExpense(i, propertyValue)

		// Compute rent expenses
		analysisVariables.computeRentExpense(i, monthlyRent)
	}

	rent, purchase, verdict, message := analyzeResult(payload.OccupancyDuration)

	resp = housinggrpc.AnalysisResponse{
		Rent:     rent,
		Purchase: purchase,
		Message:  message,
		Verdict:  verdict,
	}

	return
}

func (_var AnalysisVariables) computePurchaseExpense(year int, propertyValue float32) {
	var sum float32 = 0.00
	var expenses []float32

	if year == 0 {
		expenses = append(expenses, _var.DownPayment)
		expenses = append(expenses, _var.ClosingCost*propertyValue)
		expenses = append(expenses, _var.RealtorFees*propertyValue)
		expenses = append(expenses, _var.PropertyTransferTax*propertyValue)
	} else {
		expenses = append(expenses, _var.AnnualRepayment)
		expenses = append(expenses, _var.HomeOwnerAssociationRate*propertyValue)
		expenses = append(expenses, _var.HomeOwnerAssociationInsuranceRate*propertyValue)
		expenses = append(expenses, _var.MaintenanceRate*propertyValue)
		expenses = append(expenses, _var.UtilityRate*propertyValue)
		expenses = append(expenses, _var.PropertyTax*propertyValue)
	}

	for _, item := range expenses {
		sum += item
	}

	setExpense(cummulativePurchaseExpense, year, sum)
}

func (_var AnalysisVariables) computeRentExpense(year int, monthlyRent float32) {
	var expense float32

	if year == 0 {
		expense = _var.SecurityDeposit
	} else {
		expense = monthlyRent * 12
	}

	setExpense(cummulativeRentExpense, year, expense)
}
