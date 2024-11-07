package repositories

import (
	"database/sql"
	"go-stock-app/models"
)

type StockRepository struct {
	DB *sql.DB
}

func (r *StockRepository) CreateStock(stock *models.Stock) error {
	query := "INSERT INTO stocks (name, price, quantity) VALUES (?, ?, ?)"
	_, err := r.DB.Exec(query, stock.Name, stock.Price, stock.Quantity)
	return err
}

func (r *StockRepository) ListStocks() ([]models.Stock, error) {
	query := "SELECT id, name, price, quantity FROM stocks"
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var stocks []models.Stock
	for rows.Next() {
		var stock models.Stock
		if err := rows.Scan(&stock.ID, &stock.Name, &stock.Price, &stock.Quantity); err != nil {
			return nil, err
		}
		stocks = append(stocks, stock)
	}

	return stocks, nil
}

func (r *StockRepository) UpdateStock(stock *models.Stock) error {
	query := "UPDATE stocks SET name = ?, price = ?, quantity = ? WHERE id = ?"
	_, err := r.DB.Exec(query, stock.Name, stock.Price, stock.Quantity, stock.ID)
	return err
}

func (r *StockRepository) DeleteStock(id int) error {
	query := "DELETE FROM stocks WHERE id = ?"
	_, err := r.DB.Exec(query, id)
	return err
}
