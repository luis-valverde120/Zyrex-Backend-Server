package ports

import "github.com/luis-valverde120/Zyrex-Backend-Server/internal/core/models"

// AuthService defines the interface for authentication-related operations in the application.
// It provides methods to create, read, update, and delete users, as well as to manage user roles and permissions
// This interface is implemented by the user service, which interacts with the user repository to perform these operations.
type AuthService interface {
	// user related operations
	RegisterUser(user *models.User) (*models.User, error)   // Register a new user
	LoginUser(email, password string) (*models.User, error) // Login a user
	LogoutUser(userID int) error                            // Logout a user
	// password related operations
	ChangePassword(userID int, oldPassword, newPassword string) error // Change a user's password
	ResetPassword(email string) error                                 // Reset a user's password
	// token related operations
	GenerateToken(userID int) (string, error) // Generate a token for a user
	ValidateToken(token string) (int, error)  // Validate a token and return the user ID
}
