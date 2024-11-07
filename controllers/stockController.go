package controllers

import (
	"go-stock-app/models"
	"go-stock-app/services"
	"net/http"
	"strconv"

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

func (c *StockController) ListStocks(ctx *gin.Context) {
	stocks, err := c.Service.ListStocks()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list stocks"})
		return
	}

	ctx.JSON(http.StatusOK, stocks)
}

func (c *StockController) UpdateStock(ctx *gin.Context) {
	var stock models.Stock
	if err := ctx.ShouldBindJSON(&stock); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.Service.UpdateStock(&stock); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update stock"})
		return
	}

	ctx.JSON(http.StatusOK, stock)
}

func (c *StockController) DeleteStock(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid stock ID"})
		return
	}

	if err := c.Service.DeleteStock(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete stock"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Stock deleted"})
}
