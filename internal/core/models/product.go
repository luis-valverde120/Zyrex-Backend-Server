package models

import "errors"

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

// methods
// ChangePrice updates the price of the product. If the new price is negative, it sets the price to 0 and returns an error.
func (p *Product) ChangePrice(price float64) error {
	if price < 0 {
		price = 0
		//return error  the price not negative
		return errors.New("price cannot be negative")
	}
	p.Price = price
	return nil
}

// ChangeStock update the stock of the product. If the new stock is negative, it sets the stock to 0 and returns an error.
func (p *Product) ChangeStock(stock int) error {
	if stock < 0 {
		stock = 0
		//return error  the stock not negative
		return errors.New("stock cannot be negative")
	}
	p.Stock = stock
	return nil
}

// DecreaseStock decreases the stock of the product by the specified quantity. If the quantity is negative or greater than the current stock, it returns an error.
func (p *Product) DecreaseStock(quantity int) error {
	// Validate if has a stock of procut
	if quantity < 0 {
		return errors.New("quantity cannot be negative")
	}
	if p.Stock < quantity {
		return errors.New("not enough stock")
	}
	p.Stock -= quantity
	return nil
}

func (p *Product) Validate() error {
	if p.Name == "" {
		return errors.New("name is required")
	}
	if p.Description == "" {
		return errors.New("description is required")
	}
	if p.ImgURL == "" {
		return errors.New("image URL is required")
	}
	if p.Price <= 0 {
		return errors.New("price must be greater than 0")
	}
	if p.Stock < 0 {
		return errors.New("stock cannot be negative")
	}

	/*
	* Size, Weight, Color are optional fields, so we don't need to validate them.
	 */

	if p.StoreID <= 0 {
		return errors.New("store ID must be greater than 0")
	}
	if p.CategoryID <= 0 {
		return errors.New("category ID must be greater than 0")
	}
	return nil
}

// GetProduct returns the product details.
func (p *Product) GetProduct() Product {
	return *p
}
