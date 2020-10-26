package controllers

import (
	"fmt"

	"github.com/go-playground/validator"
	"github.com/labstack/echo"

	"github.com/oladotunsobande/housing-grpc/gateway/services"
	"github.com/oladotunsobande/housing-grpc/gateway/utils"
)

// payloadValidation validates request payload
func payloadValidation(_request **services.BreakEvenPayload) error {
	var validate *validator.Validate = validator.New()

	err := validate.Struct(_request)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println("Validation error: ", err)
			return err
		}

		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println(err.StructField())
			fmt.Println(err.ActualTag())
			fmt.Println(err.Kind())
			fmt.Println(err.Value())
			fmt.Println(err.Param())
		}
	}

	return nil
}

// CalculateBreakEven processes request for the computation of break-even analysis
func CalculateBreakEven(ctx echo.Context) (err error) {
	request := new(services.BreakEvenPayload)

	if err := ctx.Bind(request); err != nil {
		return err
	}

	// Request payload validation
	/*err = payloadValidation(&request)
	if err != nil {
		return ctx.JSON(400, utils.ErrorMessage(err.Error()))
	}*/

	result, err := request.CalculateBreakEven()
	if err != nil {
		return ctx.JSON(400, utils.ErrorMessage(err.Error()))
	}

	return ctx.JSON(200, utils.SuccessResult(result))
}
