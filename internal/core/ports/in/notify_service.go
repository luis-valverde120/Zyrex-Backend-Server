package ports

import "github.com/luis-valverde120/Zyrex-Backend-Server/internal/core/models"

// UserService defines the interface for user-related operations in the application.
// It provides methods to create, read, update, and delete users, as well as to manage user roles and permissions.
// This interface is implemented by the user service, which interacts with the user repository to perform these operations.
type NotifyService interface {
	NotifyOwner(owner *models.Owner, message string) error                  // Notify the owner with a message
	NotifyClient(client *models.Client, message string) error               // Notify the client with a message
	CreateNotification(notification *models.Notify) (*models.Notify, error) // Create a new notification
	GetNotificationByIDOwner(id int) (*models.Notify, error)                // Get a notification by ID for the owner
	GetNotificationByIDClient(id int) (*models.Notify, error)               // Get a notification by ID for the client
	MarkNotificationAsRead(id int) error                                    // Mark a notification as read
}
