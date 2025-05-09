package ports

import "github.com/luis-valverde120/Zyrex-Backend-Server/internal/core/models"

// UserService defines the interface for user-related operations in the application.
// It provides methods to create, read, update, and delete users, as well as to manage user roles and permissions.
// This interface is implemented by the user service, which interacts with the user repository to perform these operations.
type OrderService interface {
	CreateOrder(order *models.Order) (*models.Order, error)          // Create a new order in the database
	UpdateOrder(order *models.Order) (*models.Order, error)          // Update an existing order in the database
	DeleteOrder(id int) error                                        // Delete an order by ID
	GetOrderByID(id int) (*models.Order, error)                      // Get an order by ID
	GetOrderByClientID(clientID int) ([]*models.Order, error)        // Get all orders by client ID
	GetProductsListByOrderID(orderID int) ([]*models.Product, error) // Get all products by order ID
}
