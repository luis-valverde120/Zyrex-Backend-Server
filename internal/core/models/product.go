package models

/*
* Product Represents a product in the e-commerce application.
* It contains fields for the product ID, name, description, image URL, price, stock, size, weight,
* color, store ID, and category ID.
 */
type Product struct {
	ID          int     `json:"id" gorm:"primaryKey"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	ImgURL      string  `json:"img_url"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
	Size        string  `json:"size"`
	Weight      string  `json:"weight"`
	Color       string  `json:"color"`
	StoreID     int     `json:"store_id"`
	CategoryID  int     `json:"category_id"`
}
