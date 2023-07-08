package routes

import (
	"github.com/arifmr/dbo/controllers"
	"github.com/arifmr/dbo/middlewares"
	"github.com/gin-gonic/gin"
)

func SetupCustomerRoutes(router *gin.Engine, customerController *controllers.CustomerController) {
	customerRoutes := router.Group("/customers")
	{
		customerRoutes.GET("", customerController.GetCustomers)
		customerRoutes.POST("", customerController.CreateCustomer)
		customerRoutes.GET("/search", customerController.SearchCustomers)
		customerRoutes.GET("/:customer_id", customerController.GetCustomerDetail)
		customerRoutes.PUT("/:customer_id", middlewares.AuthMiddleware(), customerController.UpdateCustomer)
		customerRoutes.DELETE("/:customer_id", middlewares.AuthMiddleware(), customerController.DeleteCustomer)
	}
}
