package controllers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/arifmr/dbo/models"
	"github.com/arifmr/dbo/repositories"
	"github.com/arifmr/dbo/utils"
	"github.com/gin-gonic/gin"
)

type LoginDataController struct {
	repo         repositories.LoginDataRepository
	customerRepo repositories.CustomerRepository
}

func NewLoginDataController() *LoginDataController {
	return &LoginDataController{
		repo: *repositories.NewLoginDataRepository(),
	}
}

func (l *LoginDataController) GetLoginDataByCustomerID(ctx *gin.Context) {
	customerID, _ := strconv.Atoi(ctx.Param("customer_id"))

	loginData, err := l.repo.GetLoginDataByCustomerID(int64(customerID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, loginData)
}

func (l *LoginDataController) Login(ctx *gin.Context) {
	var credentials models.AuthData
	if err := ctx.ShouldBindJSON(&credentials); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if credentials.Email == "" || credentials.Password == "" || !strings.Contains(credentials.Email, "@") {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Email and Password cannot be empty"})
	}

	if !strings.Contains(credentials.Email, "@") {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Email Format"})
	}

	customer, err := l.customerRepo.GetCustomerByEmail(credentials.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if customer.ID == 0 {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Email or Password"})
		return
	}

	isPasswordCorrect := utils.ComparePassword(customer.Password, credentials.Password)

	if !isPasswordCorrect {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Email or Password"})
		return
	}

	token, err := utils.GenerateToken(credentials.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	reqInsertLoginData := models.LoginData{
		CustomerID: customer.ID,
		Token:      token,
		IPAddress:  ctx.ClientIP(),
	}

	err = l.repo.InsertLoginData(&reqInsertLoginData)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})
}
