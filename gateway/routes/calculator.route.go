package routes

import (
	"github.com/labstack/echo"

	"github.com/oladotunsobande/housing-grpc/gateway/controllers"
)

// Index configures all routes
func Index(root string, router *echo.Group) {
	// Compute break-even
	router.POST(root+"/compute", controllers.CalculateBreakEven)
}
