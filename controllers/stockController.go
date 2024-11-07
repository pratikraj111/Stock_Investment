package controllers

import (
	"go-stock-app/models"
	"go-stock-app/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type StockController struct {
	Service *services.StockService
}

func (c *StockController) CreateStock(ctx *gin.Context) {
	var stock models.Stock
	if err := ctx.ShouldBindJSON(&stock); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.Service.CreateStock(&stock); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create stock"})
		return
	}

	ctx.JSON(http.StatusCreated, stock)
}
