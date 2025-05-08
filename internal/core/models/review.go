package models

/*
* Review Represents a product review in the e-commerce application.
* It contains fields for the review ID, product ID, client ID, rating, comment, and creation date.
 */
type Review struct {
	ID        int    `json:"id" gorm:"primaryKey"`
	ProductID int    `json:"product_id"`
	ClientID  int    `json:"client_id"`
	Rating    int    `json:"rating"`
	Comment   string `json:"comment"`
	CreatedAt string `json:"created_at"`
}
