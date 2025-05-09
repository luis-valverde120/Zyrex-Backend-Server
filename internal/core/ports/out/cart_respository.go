package ports

import "github.com/luis-valverde120/Zyrex-Backend-Server/internal/core/models"

// UserService defines the interface for user-related operations in the application.
// It provides methods to create, read, update, and delete users, as well as to manage user roles and permissions.
// This interface is implemented by the user service, which interacts with the user repository to perform these operations.
type CartRepository interface {
	Save(cart *models.Cart) (*models.Cart, error)                               // Save a new cart to the database
	Update(cart *models.Cart) (*models.Cart, error)                             // Update an existing cart in the database
	FindByID(id int) (*models.Cart, error)                                      // Find a cart by ID
	FindProductsByCartID(cartID int) ([]*models.Product, error)                 // Find products by cart ID
	AddProductToCart(cartID int, product *models.Product) (*models.Cart, error) // Add a product to a cart
	FindCartByClientID(clientID int) ([]*models.Cart, error)                    // Find cart and product by client ID
}
