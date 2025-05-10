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
	Images      []Image `json:"images"`
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
func (product *Product) ChangePrice(price float64) error {
	if price < 0 {
		price = 0
		//return error  the price not negative
		return errors.New("price cannot be negative")
	}
	product.Price = price
	return nil
}

// ChangeStock update the stock of the product. If the new stock is negative, it sets the stock to 0 and returns an error.
func (product *Product) ChangeStock(stock int) error {
	if stock < 0 {
		stock = 0
		//return error  the stock not negative
		return errors.New("stock cannot be negative")
	}
	product.Stock = stock
	return nil
}

// DecreaseStock decreases the stock of the product by the specified quantity. If the quantity is negative or greater than the current stock, it returns an error.
func (product *Product) DecreaseStock(quantity int) error {
	// Validate if has a stock of procut
	if quantity < 0 {
		return errors.New("quantity cannot be negative")
	}
	if product.Stock < quantity {
		return errors.New("not enough stock")
	}
	product.Stock -= quantity
	return nil
}

func (product *Product) Validate() error {
	if product.Name == "" {
		return errors.New("name is required")
	}
	if product.Description == "" {
		return errors.New("description is required")
	}
	if product.Price <= 0 {
		return errors.New("price must be greater than 0")
	}
	if product.Stock < 0 {
		return errors.New("stock cannot be negative")
	}

	/*
	* Size, Weight, Color are optional fields, so we don't need to validate them.
	 */

	if product.StoreID <= 0 {
		return errors.New("store ID must be greater than 0")
	}
	if product.CategoryID <= 0 {
		return errors.New("category ID must be greater than 0")
	}

	// Validate images
	for _, image := range product.Images {
		if err := image.Validate(); err != nil {
			return err
		}
	}

	return nil
}

// GetProduct returns the product details.
func (product *Product) GetProduct() Product {
	return *product
}

func (product *Product) ValidateToBuyProduct() error {
	if product.Stock <= 0 {
		return errors.New("product is out of stock")
	}
	if product.Price <= 0 {
		return errors.New("price must be greater than 0")
	}
	return nil
}
