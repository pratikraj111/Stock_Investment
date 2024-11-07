package controllers

import (
	"go-stock-app/models"
	"go-stock-app/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type OrderController struct {
	Service *services.OrderService
}

func (c *OrderController) PlaceOrder(ctx *gin.Context) {
	var order models.Order
	if err := ctx.ShouldBindJSON(&order); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.Service.PlaceOrder(&order); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to place order"})
		return
	}

	ctx.JSON(http.StatusCreated, order)
}
