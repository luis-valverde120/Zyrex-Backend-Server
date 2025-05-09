package ports

import "github.com/luis-valverde120/Zyrex-Backend-Server/internal/core/models"

// UserService defines the interface for user-related operations in the application.
// It provides methods to create, read, update, and delete users, as well as to manage user roles and permissions.
// This interface is implemented by the user service, which interacts with the user repository to perform these operations.
type ClientRepository interface {
	Save(store *models.Client) (*models.Client, error)   // Save a new store to the database
	Update(store *models.Client) (*models.Client, error) // Update an existing store in the database
	Delete(id int) error                                 // Delete a store by ID
	GetDirectionByID(id int) (*models.Client, error)     // Get a direction by ID
	FindByID(id int) (*models.Client, error)             // Find a store by ID
}
