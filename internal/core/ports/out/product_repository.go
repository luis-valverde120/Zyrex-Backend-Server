package ports

import "github.com/luis-valverde120/Zyrex-Backend-Server/internal/core/models"

// ProductRepository defines the interface for product-related database operations.
// It provides methods to create, read, update, and delete products in the database.
// This interface is implemented by the product repository, which interacts with the database to perform these operations.
// The methods in this interface return a pointer to the product model and an error if any occurs during the operation.
// The methods are designed to be used by the service layer of the application to manage product data.

type ProductRepository interface {
	Save(product *models.Product) (*models.Product, error)                    // Save a new product to the database
	FindByID(id int) (*models.Product, error)                                 // Find a product by its ID
	Update(*models.Product) (*models.Product, error)                          // Update an existing product
	Delete(id int) error                                                      // Delete a product by its ID
	FindAll() ([]*models.Product, error)                                      // Find all products in the database
	FindByStoreID(storeID int) ([]*models.Product, error)                     // Find products by store ID
	FindByCategoryID(categoryID int) ([]*models.Product, error)               // Find products by category ID
	CreateNewImage(productID int, image *models.Image) (*models.Image, error) // Create a new image for a product
}
