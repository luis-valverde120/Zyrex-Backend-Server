package ports

import "github.com/luis-valverde120/Zyrex-Backend-Server/internal/core/models"

// UserService defines the interface for user-related operations in the application.
// It provides methods to create, read, update, and delete users, as well as to manage user roles and permissions.
// This interface is implemented by the user service, which interacts with the user repository to perform these operations.
// The methods in this interface return a pointer to the user model and an error if any occurs during the operation.
type ClientService interface {
	CreateClient(client *models.Client) (*models.Client, error) // Create a new client in the database
	UpdateClient(client *models.Client) (*models.Client, error) // Update an existing client in the database
	DeleteClient(id int) error                                  // Delete a client by ID
}
