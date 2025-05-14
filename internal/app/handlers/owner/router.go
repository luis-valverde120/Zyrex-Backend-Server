package owner

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(rg *gin.RouterGroup, h *Handler) {
	group := rg.Group("/owners")
	group.POST("/", h.Create)                    // Create a new owner
	group.PUT("/:id", h.Update)                  // Update an owner by ID
	group.DELETE("/:id", h.Delete)               // Delete an owner by ID
	group.GET("/", h.List)                       // List all owners
	group.GET("/store/:store_id", h.ListByStore) // List owners by store ID
	group.GET("/:id/store", h.GetStoreByOwnerID) // Get store by owner ID
}
