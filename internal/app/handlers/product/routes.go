package product

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(rg *gin.RouterGroup, h *Handler) {
	group := rg.Group("/products")
	group.GET("/", h.List)                                // List all products
	group.POST("/", h.Create)                             // Create a new product
	group.GET("/:id", h.GetByID)                          // Get product by ID
	group.PUT("/:id", h.Update)                           // Update product by ID
	group.DELETE("/:id", h.Delete)                        // Delete product by ID
	group.GET("/store/:store_id", h.ListByStore)          // List products by store ID
	group.POST("/sync/:store_id", h.SyncByStore)          // Sync products by store ID
	group.POST("/:id/image", h.AddImage)                  // Add an image to a product
	group.GET("/category/:category_id", h.ListByCategory) // List products by category ID
	group.GET("/owner/:owner_id", h.ListByOwner)          // List products by owner ID
}
