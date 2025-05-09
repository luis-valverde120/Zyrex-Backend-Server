package ports

import "github.com/luis-valverde120/Zyrex-Backend-Server/internal/core/models"

// UserService defines the interface for user-related operations in the application.
// It provides methods to create, read, update, and delete users, as well as to manage user roles and permissions.
// This interface is implemented by the user service, which interacts with the user repository to perform these operations.

type OwnerService interface {
	CreateOwner(owner *models.Owner) (*models.Owner, error) // Create a new owner in the database
	UpdateOwner(owner *models.Owner) (*models.Owner, error) // Update an existing owner in the database
	DeleteOwner(id int) error                               // Delete an owner by ID
}
