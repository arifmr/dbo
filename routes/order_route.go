package routes

import (
	"github.com/arifmr/dbo/controllers"
	"github.com/arifmr/dbo/middlewares"
	"github.com/gin-gonic/gin"
)

func SetupOrderRoutes(router *gin.Engine, orderController *controllers.OrderController) {
	orderRoutes := router.Group("/orders")
	{
		orderRoutes.GET("", orderController.GetOrders)
		orderRoutes.GET("/search", orderController.SearchOrders)
		orderRoutes.POST("", middlewares.AuthMiddleware(), orderController.CreateOrder)
		orderRoutes.GET("/:order_id", orderController.GetOrderDetail)
		orderRoutes.PUT("/:order_id", middlewares.AuthMiddleware(), orderController.UpdateOrder)
		orderRoutes.DELETE("/:order_id", middlewares.AuthMiddleware(), orderController.DeleteOrder)
	}
}
