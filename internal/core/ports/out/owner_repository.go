package ports

import "github.com/luis-valverde120/Zyrex-Backend-Server/internal/core/models"

// UserService defines the interface for user-related operations in the application.
// It provides methods to create, read, update, and delete users, as well as to manage user roles and permissions.
// This interface is implemented by the user service, which interacts with the user repository to perform these operations.
type OwnerRepository interface {
	Save(store *models.Owner) (*models.Owner, error)    // Save a new store to the database
	Update(store *models.Owner) (*models.Owner, error)  // Update an existing store in the database
	Delete(id int) error                                // Delete a store by ID
	FindAll() ([]*models.Owner, error)                  // Find all stores in the database
	FindByID(id int) (*models.Owner, error)             // Find a store by ID
	FindByStoreID(storeID int) ([]*models.Owner, error) // Find stores by owner ID
	verifyStore(id int) error                           // Verify a store by ID
	SendNotification(storeID int, message string) error // Send a notification to a store by ID
	GetStoreByID(id int) (*models.Store, error)         // Get a store by ID of owner
}
