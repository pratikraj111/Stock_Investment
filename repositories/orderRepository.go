package repositories

import (
	"database/sql"
	"go-stock-app/models"
)

type OrderRepository struct {
	DB *sql.DB
}

func (r *OrderRepository) PlaceOrder(order *models.Order) error {
	query := "INSERT INTO orders (user_id, stock_id, quantity) VALUES (?, ?, ?)"
	_, err := r.DB.Exec(query, order.UserID, order.StockID, order.Quantity)
	return err
}

func (r *OrderRepository) ListOrders(userID int) ([]models.Order, error) {
	query := "SELECT id, user_id, stock_id, quantity FROM orders WHERE user_id = ?"
	rows, err := r.DB.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []models.Order
	for rows.Next() {
		var order models.Order
		if err := rows.Scan(&order.ID, &order.UserID, &order.StockID, &order.Quantity); err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}

	return orders, nil
}
