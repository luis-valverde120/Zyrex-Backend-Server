package models

/*
* Category Represents a product category in the e-commerce application.
* It contains fields for the category ID, name, and description.
 */
type Category struct {
	ID          int    `json:"id" gorm:"primaryKey"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
