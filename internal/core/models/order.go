package models

/*
*	Order Represents an order in the e-commerce application.
*	It contains fields for the order ID, client ID, order date, status, and total amount.
 */
type Order struct {
	ID          int     `json:"id" gorm:"primaryKey"`
	ClientID    int     `json:"client_id"`
	OrderDate   string  `json:"order_date"`
	Status      string  `json:"status"`
	TotalAmount float64 `json:"total_amount"`
}
