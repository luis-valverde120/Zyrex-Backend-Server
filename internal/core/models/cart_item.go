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

// GetCartID returns the cart ID of the CartItem.
func (cart_item *CartItem) GetID() int {
	return cart_item.ID
}

// GetCartID returns the cart ID of the CartItem.
func (cart_item *CartItem) GetCartID() int {
	return cart_item.CartID
}

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
