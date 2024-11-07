package models

type Order struct {
	ID       int
	UserID   int
	StockID  int
	Quantity int
}
