package models

/*
* CartItem Represents an item in a shopping cart.
* It contains fields for the cart ID, product ID, and quantity of the item.
 */
type CartItem struct {
	ID        int `json:"id" gorm:"primaryKey"`
	CartID    int `json:"cart_id"`
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
}
