package server

import (
	"go-stock-app/controllers"
	"go-stock-app/repositories"
	"go-stock-app/services"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, db *sql.DB) {
	// Initialize repositories, services, and controllers
	stockRepo := &repositories.StockRepository{DB: db}
	orderRepo := &repositories.OrderRepository{DB: db}

	stockService := &services.StockService{Repo: stockRepo}
	orderService := &services.OrderService{Repo: orderRepo}

	stockController := &controllers.StockController{Service: stockService}
	orderController := &controllers.OrderController{Service: orderService}

	router.POST("/api/admin/stocks", stockController.CreateStock)
	router.POST("/api/user/orders", orderController.PlaceOrder)
}
