package models

/*
* OrderDetail Represents the details of an order in the e-commerce application.
* It contains fields for the order detail ID, order ID, product ID, quantity, and unit price.
 */
type OrderDetail struct {
	ID        int     `json:"id" gorm:"primaryKey"`
	OrderID   int     `json:"order_id"`
	ProductID int     `json:"product_id"`
	Quantity  int     `json:"quantity"`
	UnitPrice float64 `json:"unit_price"`
}
