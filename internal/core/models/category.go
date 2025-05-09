package models

import (
	"errors"
	"strings"
)

/*
* Category Represents a product category in the e-commerce application.
* It contains fields for the category ID, name, and description.
 */
type Category struct {
	ID          int    `json:"id" gorm:"primaryKey"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// * methods for Category struct can be added here if needed

// Validate checks if the category name is valid.
func (cart *Category) Validate() error {
	if cart.Name == "" {
		return errors.New("category name is required")
	}

	if strings.TrimSpace(cart.Name) == "" {
		return errors.New("category name cannot be empty or whitespace")
	}

	return nil
}
