package ports

import "github.com/luis-valverde120/Zyrex-Backend-Server/internal/core/models"

// UserService defines the interface for user-related operations in the application.
// It provides methods to create, read, update, and delete users, as well as to manage user roles and permissions.
// This interface is implemented by the user service, which interacts with the user repository to perform these operations.
// * When the client buy Cart this generate a order and save products in orders details
type CartService interface {
	CreateCart(cart *models.Cart) (*models.Cart, error)            // Create a new cart in the database
	UpdateCart(cart *models.Cart) (*models.Cart, error)            // Update an existing cart in the database
	DeleteProductFromCart(cartID, productID int) error             // Delete a product from a cart by cart ID and product ID
	BuyCart(cartID int) error                                      // Buy a cart by cart ID
	AddProductToCart(cartID, productID int) error                  // Add a product to a cart by cart ID and product ID
	GetListProductsByCartID(cartID int) ([]*models.Product, error) // Get a list of products by cart ID
}
