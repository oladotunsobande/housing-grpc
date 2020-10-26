package controllers

import (
	"github.com/labstack/echo"

	"github.com/oladotunsobande/housing-grpc/gateway/services"
	"github.com/oladotunsobande/housing-grpc/gateway/utils"
)

// CalculateBreakEven processes request for the computation of break-even analysis
func CalculateBreakEven(ctx echo.Context) (err error) {
	request := new(services.BreakEvenPayload)

	if err := ctx.Bind(request); err != nil {
		return err
	}

	// Request payload validation
	err = PayloadValidation(request)
	if err != nil {
		return ctx.JSON(400, utils.ErrorMessage(err.Error()))
	}

	result, err := request.CalculateBreakEven()
	if err != nil {
		return ctx.JSON(400, utils.ErrorMessage(err.Error()))
	}

	return ctx.JSON(200, utils.SuccessResult(result))
}
