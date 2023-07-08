package routes

import (
	"github.com/arifmr/dbo/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	SetupCustomerRoutes(router, controllers.NewCustomerController())
	SetupOrderRoutes(router, controllers.NewOrderController())
	SetupLoginDataRoutes(router, controllers.NewLoginDataController())
}
