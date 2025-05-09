package ports

import "github.com/luis-valverde120/Zyrex-Backend-Server/internal/core/models"

// UserService defines the interface for user-related operations in the application.
// It provides methods to create, read, update, and delete users, as well as to manage user roles and permissions.
// This interface is implemented by the user service, which interacts with the user repository to perform these operations.
type OrderRepository interface {
	Save(order *models.Order) (*models.Order, error)                  // Save a new order to the database
	Update(order *models.Order) (*models.Order, error)                // Update an existing order in the database
	Delete(id int) error                                              // Delete an order by ID
	FindByID(id int) (*models.Order, error)                           // Find an order by ID
	FindByClientID(clientID int) ([]*models.Order, error)             // Find orders by client ID
	FindProductsListByOrderID(orderID int) ([]*models.Product, error) // Find products by order ID
}
