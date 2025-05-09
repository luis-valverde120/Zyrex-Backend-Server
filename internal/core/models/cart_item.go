package models

import "errors"

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

// * Methods for CartItem struct */

// ValidateAddItem checks if the CartItem has all the required fields to add an item to the cart.
func (cart_item *CartItem) ValidateAddItem() error {
	if cart_item.CartID == 0 {
		return errors.New("cart ID is required")
	}
	if cart_item.ProductID == 0 {
		return errors.New("product ID is required")
	}
	if cart_item.Quantity <= 0 {
		return errors.New("quantity must be greater than 0")
	}

	return nil
}
