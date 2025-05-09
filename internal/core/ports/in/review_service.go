package ports

import "github.com/luis-valverde120/Zyrex-Backend-Server/internal/core/models"

// UserService defines the interface for user-related operations in the application.
// It provides methods to create, read, update, and delete users, as well as to manage user roles and permissions.
// This interface is implemented by the user service, which interacts with the user repository to perform these operations.
type ReviewService interface {
	CreateReview(review *models.Review) (*models.Review, error) // Create a new review in the database
	UpdateReview(review *models.Review) (*models.Review, error) // Update an existing review in the database
	DeleteReview(id int) error                                  // Delete a review by ID
}
