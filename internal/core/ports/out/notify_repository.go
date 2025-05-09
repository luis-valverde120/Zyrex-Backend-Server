package ports

import "github.com/luis-valverde120/Zyrex-Backend-Server/internal/core/models"

// UserService defines the interface for user-related operations in the application.
// It provides methods to create, read, update, and delete users, as well as to manage user roles and permissions.
// This interface is implemented by the user service, which interacts with the user repository to perform these operations.
type NotifyRepository interface {
	Save(notification *models.Notification) (*models.Notification, error)   // Save a new notify to the database
	Update(notification *models.Notification) (*models.Notification, error) // Update an existing notify in the database
	Delete(id int) error                                                    // Delete a notify by ID
	FindByOwnerID(ownerID int) ([]*models.Notification, error)              // Find notifies by owner ID
	FindByID(id int) (*models.Notification, error)                          // Find a notify by ID
	FindByClientID(clientID int) ([]*models.Notification, error)            // Find notifies by client ID
}
