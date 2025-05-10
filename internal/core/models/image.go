package models

import (
	"errors"
	"strings"
)

/*
* Image Represents an image in the e-commerce application.
* It contains fields for the image ID, product ID, and URL.
 */
type Image struct {
	ID        int    `json:"id" gorm:"primaryKey"`
	ProductID int    `json:"product_id"`
	URL       string `json:"url"`
}

// * methods of Image model

func (image *Image) Validate() error {
	if image.URL == "" {
		return errors.New("image URL is required")
	}

	if strings.TrimSpace(image.URL) == "" {
		return errors.New("image URL cannot be empty or whitespace")
	}

	return nil
}
