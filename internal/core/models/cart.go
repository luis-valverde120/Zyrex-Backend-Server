package models

import "errors"

/*
* Cart Represents a shopping cart for a client.
* It contains fields for the cart ID, client ID, and creation date.
 */
type Cart struct {
	ID        int         `json:"id" gorm:"primaryKey"`
	ClientID  int         `json:"client_id"`
	CartItems []*CartItem `json:"cart_items"`
	CreatedAt string      `json:"created_at"`
}

// * Methods to manage the cart can be added here, such as adding items, removing items, and calculating the total amount.

func (cart *Cart) AddItem(product Product, quantity int) error {
	if quantity <= 0 {
		return errors.New("quantity must be greater than 0")
	}

	// if the product is already in the cart, update the quantity
	for _, item := range cart.CartItems {
		if item.ProductID == product.ID {
			item.Quantity += quantity
			return nil
		}
	}

	// if the product is not in the cart, add a new item
	cart.CartItems = append(cart.CartItems, &CartItem{
		CartID:    cart.ID,
		ProductID: product.ID,
		Quantity:  quantity,
	})

	return nil
}

// RemoveItem removes an item from the cart by product ID.
func (cart *Cart) RemoveItem(productID int) error {
	for i, item := range cart.CartItems {
		if item.ProductID == productID {
			cart.CartItems = append(cart.CartItems[:i], cart.CartItems[i+1:]...)
			return nil
		}
	}
	return errors.New("item not found in cart")
}
