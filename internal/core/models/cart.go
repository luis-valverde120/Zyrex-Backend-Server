package models

/*
* Cart Represents a shopping cart for a client.
* It contains fields for the cart ID, client ID, and creation date.
 */
type Cart struct {
	ID        int    `json:"id" gorm:"primaryKey"`
	ClientID  int    `json:"client_id"`
	CreatedAt string `json:"created_at"`
}
