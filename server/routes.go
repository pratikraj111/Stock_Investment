package server

import (
	"database/sql"
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

	admin := router.Group("/api/admin")
	{
		admin.POST("/stocks", stockController.CreateStock)
		admin.GET("/stocks", stockController.ListStocks)
		admin.PUT("/stocks", stockController.UpdateStock)
		admin.DELETE("/stocks/:id", stockController.DeleteStock)
	}

	user := router.Group("/api/user")
	{
		user.GET("/stocks", stockController.ListStocks)
		user.POST("/orders", orderController.PlaceOrder)
	}
}
