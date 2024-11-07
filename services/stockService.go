package services

import (
	"go-stock-app/models"
	"go-stock-app/repositories"
)

type StockService struct {
	Repo *repositories.StockRepository
}

func (s *StockService) CreateStock(stock *models.Stock) error {
	return s.Repo.CreateStock(stock)
}

func (s *StockService) ListStocks() ([]models.Stock, error) {
	return s.Repo.ListStocks()
}

func (s *StockService) UpdateStock(stock *models.Stock) error {
	return s.Repo.UpdateStock(stock)
}

func (s *StockService) DeleteStock(id int) error {
	return s.Repo.DeleteStock(id)
}
