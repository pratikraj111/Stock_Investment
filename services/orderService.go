package services

import (
	"go-stock-app/models"
	"go-stock-app/repositories"
)

type OrderService struct {
	Repo *repositories.OrderRepository
}

func (s *OrderService) PlaceOrder(order *models.Order) error {
	return s.Repo.PlaceOrder(order)
}

func (s *OrderService) ListOrders(userID int) ([]models.Order, error) {
	return s.Repo.ListOrders(userID)
}
