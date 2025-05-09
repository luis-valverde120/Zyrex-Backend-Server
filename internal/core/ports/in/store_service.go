package ports

import "github.com/luis-valverde120/Zyrex-Backend-Server/internal/core/models"

// UserService defines the interface for user-related operations in the application.
// It provides methods to create, read, update, and delete users, as well as to manage user roles and permissions.
// This interface is implemented by the user service, which interacts with the user repository to perform these operations.
type StoreService interface {
	CreateStore(store *models.Store) (*models.Store, error) // Create a new store in the database
	UpdateStore(store *models.Store) (*models.Store, error) // Update an existing store in the database
	DeleteStore(id int) error                               // Delete a store by ID
	GetStoreByID(id int) (*models.Store, error)
	GetStoreByOwner(ownerID string) (*models.Store, error)
	SyncStoreData(ownerID string) error // si trabajas offline/online           // Get a store by ID
}
