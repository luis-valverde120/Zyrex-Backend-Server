package ports

import "github.com/luis-valverde120/Zyrex-Backend-Server/internal/core/models"

// UserService defines the interface for user-related operations in the application.
// It provides methods to create, read, update, and delete users, as well as to manage user roles and permissions.
// This interface is implemented by the user service, which interacts with the user repository to perform these operations.
type StoreRepository interface {
	Save(store *models.Store) (*models.Store, error)    // Save a new store to the database
	Update(store *models.Store) (*models.Store, error)  // Update an existing store in the database
	Delete(id int) error                                // Delete a store by ID
	FindAll() ([]*models.Store, error)                  // Find all stores in the database
	FindByID(id int) (*models.Store, error)             // Find a store by ID
	FindByOwnerID(ownerID int) ([]*models.Store, error) // Find stores by owner ID
}
