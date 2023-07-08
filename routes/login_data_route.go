package routes

import (
	"github.com/arifmr/dbo/controllers"
	"github.com/gin-gonic/gin"
)

func SetupLoginDataRoutes(router *gin.Engine, loginData *controllers.LoginDataController) {
	authRoutes := router.Group("/login")
	{
		authRoutes.POST("", loginData.Login)
		authRoutes.GET("/:customer_id", loginData.GetLoginDataByCustomerID)
	}
}
