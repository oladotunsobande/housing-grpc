package config

import (
	"fmt"
	"log"
	"os"
	"path"

	"github.com/joho/godotenv"
)

// Secrets Struct for all environment variables
type Secrets struct {
	Environment                       string
	ApplicationPort                   string
	ApplicationName                   string
	CalculatorServicePort             string
	PropertyAppreciationRate          string
	AnnualInflation                   string
	ClosingCost                       string
	RealtorFees                       string
	LoanDuration                      string
	LoanInterestRate                  string
	AnnualRentIncreaseRate            string
	PropertyTax                       string
	PropertyTransferTax               string
	MaintenanceRate                   string
	UtilityRate                       string
	HomeOwnerAssociationRate          string
	HomeOwnerAssociationInsuranceRate string
}

func init() {
	dirname, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	if err := godotenv.Load(path.Join(dirname, ".env")); err != nil {
		fmt.Println("Error: ", err)
		log.Fatal("Error loading .env file")
	}
	log.Println(".env loaded")
}

// GetSecrets Get all loaded secrets
func GetSecrets() Secrets {
	var secrets Secrets

	secrets.Environment = os.Getenv("GO_ENV")
	secrets.ApplicationPort = os.Getenv("APP_PORT")
	secrets.ApplicationName = os.Getenv("APP_NAME")
	secrets.CalculatorServicePort = os.Getenv("CALCULATOR_SERVICE_PORT")
	secrets.PropertyAppreciationRate = os.Getenv("PROPERTY_APPRECIATION_RATE")
	secrets.AnnualInflation = os.Getenv("ANNUAL_INFLATION")
	secrets.ClosingCost = os.Getenv("CLOSING_COST")
	secrets.RealtorFees = os.Getenv("REALTOR_FEES")
	secrets.LoanDuration = os.Getenv("LOAN_DURATION")
	secrets.LoanInterestRate = os.Getenv("LOAN_INTEREST_RATE")
	secrets.AnnualRentIncreaseRate = os.Getenv("ANNUAL_RENT_INCREASE_RATE")
	secrets.PropertyTax = os.Getenv("PROPERTY_TAX")
	secrets.PropertyTransferTax = os.Getenv("PROPERTY_TRANSFER_TAX")
	secrets.MaintenanceRate = os.Getenv("MAINTENANCE_RATE")
	secrets.UtilityRate = os.Getenv("UTILITIES_RATE")
	secrets.HomeOwnerAssociationRate = os.Getenv("HOME_OWNER_ASSOCIATION_RATE")
	secrets.HomeOwnerAssociationInsuranceRate = os.Getenv("HOME_OWNER_ASSOCIATION_INSURANCE_RATE")

	return secrets
}
