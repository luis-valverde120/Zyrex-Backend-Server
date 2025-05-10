package ports

import "github.com/luis-valverde120/Zyrex-Backend-Server/internal/core/models"

// UserService defines the interface for user-related operations in the application.
// It provides methods to create, read, update, and delete users, as well as to manage user roles and permissions.
// This interface is implemented by the user service, which interacts with the user repository to perform these operations.
type ProductService interface {
	CreateProduct(product *models.Product) (*models.Product, error)
	UpdateProduct(product *models.Product) (*models.Product, error)
	DeleteProduct(id int) error
	ListProductsByStore(StoreID int) ([]*models.Product, error)
	SyncProductsByStore(StoreID int) ([]*models.Product, error)
	AddNewImageToProduct(productID int, image *models.Image) (*models.Image, error)
}
