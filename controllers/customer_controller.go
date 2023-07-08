package controllers

import (
	"net/http"
	"strconv"

	"github.com/arifmr/dbo/models"
	"github.com/arifmr/dbo/repositories"
	"github.com/arifmr/dbo/utils"
	"github.com/gin-gonic/gin"
)

type CustomerController struct {
	repo repositories.CustomerRepository
}

func NewCustomerController() *CustomerController {
	return &CustomerController{
		repo: *repositories.NewCustomerRepository(),
	}
}

func (c *CustomerController) GetCustomers(ctx *gin.Context) {
	offset, _ := strconv.Atoi(ctx.DefaultQuery("offset", "0"))
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))

	customers, err := c.repo.GetCustomers(offset, limit)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, customers)
}

func (c *CustomerController) GetCustomerDetail(ctx *gin.Context) {
	customerID, _ := strconv.Atoi(ctx.Param("customer_id"))

	customer, err := c.repo.GetCustomerByID(customerID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if customer == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Customer not found"})
		return
	}

	ctx.JSON(http.StatusOK, customer)
}

func (c *CustomerController) CreateCustomer(ctx *gin.Context) {
	var customer models.Customer
	if err := ctx.ShouldBindJSON(&customer); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	encryptedPassword, err := utils.EncryptPassword(customer.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to encrypt password"})
		return
	}

	customer.Password = encryptedPassword

	err = c.repo.CreateCustomer(&customer)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, customer)
}

func (c *CustomerController) UpdateCustomer(ctx *gin.Context) {
	customerID, _ := strconv.Atoi(ctx.Param("customer_id"))

	var customer models.Customer
	if err := ctx.ShouldBindJSON(&customer); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if customer.Password != "" {
		encryptedPassword, err := utils.EncryptPassword(customer.Password)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to encrypt password"})
			return
		}

		customer.Password = encryptedPassword
	}

	customer.ID = int64(customerID)

	err := c.repo.UpdateCustomer(&customer)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, customer)
}

func (c *CustomerController) DeleteCustomer(ctx *gin.Context) {
	customerID, _ := strconv.Atoi(ctx.Param("customer_id"))

	err := c.repo.DeleteCustomer(customerID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Customer deleted successfully"})
}

func (c *CustomerController) SearchCustomers(ctx *gin.Context) {
	criteria := ctx.Query("criteria")

	customers, err := c.repo.SearchCustomers(criteria)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, customers)
}
