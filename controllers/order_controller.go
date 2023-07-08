package controllers

import (
	"net/http"
	"strconv"

	"github.com/arifmr/dbo/models"
	"github.com/arifmr/dbo/repositories"
	"github.com/gin-gonic/gin"
)

type OrderController struct {
	repo repositories.OrderRepository
}

func NewOrderController() *OrderController {
	return &OrderController{
		repo: *repositories.NewOrderRepository(),
	}
}

func (c *OrderController) GetOrders(ctx *gin.Context) {
	offset, _ := strconv.Atoi(ctx.DefaultQuery("offset", "0"))
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))

	orders, err := c.repo.GetOrders(offset, limit)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, orders)
}

func (c *OrderController) GetOrderDetail(ctx *gin.Context) {
	orderID, _ := strconv.Atoi(ctx.Param("order_id"))

	order, err := c.repo.GetOrderByID(orderID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if order == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Order not found"})
		return
	}

	ctx.JSON(http.StatusOK, order)
}

func (c *OrderController) CreateOrder(ctx *gin.Context) {
	var order models.Order
	if err := ctx.ShouldBindJSON(&order); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := c.repo.CreateOrder(&order)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, order)
}

func (c *OrderController) UpdateOrder(ctx *gin.Context) {
	orderID, _ := strconv.Atoi(ctx.Param("order_id"))

	var order models.Order
	if err := ctx.ShouldBindJSON(&order); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order.ID = int64(orderID)

	err := c.repo.UpdateOrder(&order)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, order)
}

func (c *OrderController) DeleteOrder(ctx *gin.Context) {
	orderID, _ := strconv.Atoi(ctx.Param("order_id"))

	err := c.repo.DeleteOrder(orderID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Order deleted successfully"})
}

func (c *OrderController) SearchOrders(ctx *gin.Context) {
	criteria := ctx.Query("criteria")

	orders, err := c.repo.SearchOrders(criteria)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, orders)
}
