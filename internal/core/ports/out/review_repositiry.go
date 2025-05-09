package ports

import "github.com/luis-valverde120/Zyrex-Backend-Server/internal/core/models"

// UserService defines the interface for user-related operations in the application.
// It provides methods to create, read, update, and delete users, as well as to manage user roles and permissions.
// This interface is implemented by the user service, which interacts with the user repository to perform these operations.
type ReviewRepository interface {
	Save(review *models.Review) (*models.Review, error)             // Save a new review to the database
	Update(review *models.Review) (*models.Review, error)           // Update an existing review in the database
	Delete(id int) error                                            // Delete a review by ID
	GetByID(id int) (*models.Review, error)                         // Get a review by ID
	GetAllReviewsByStoreID(storeID int) ([]*models.Review, error)   // Get all reviews by store ID
	GetAllReviewsByClientID(clientID int) ([]*models.Review, error) // Get all reviews by client ID
}
