package owner

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/luis-valverde120/Zyrex-Backend-Server/internal/core/models"
	ports "github.com/luis-valverde120/Zyrex-Backend-Server/internal/core/ports/in"
)

type Handler struct {
	service ports.OwnerService
}

func NewHandler(service ports.OwnerService) *Handler {
	return &Handler{service: service}
}

// POST /owners
func (h *Handler) Create(c *gin.Context) {
	var owner models.Owner
	if err := c.ShouldBindJSON(&owner); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := h.service.CreateOwner(&owner)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, result)
}

// PUT /owners/:id
func (h *Handler) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid owner ID"})
		return
	}

	var owner models.Owner
	if err := c.ShouldBindJSON(&owner); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	owner.ID = id

	updated, err := h.service.UpdateOwner(&owner)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updated)
}

// DELETE /owners/:id
func (h *Handler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid owner ID"})
		return
	}

	if err := h.service.DeleteOwner(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

// GET /owners
func (h *Handler) List(c *gin.Context) {
	owners, err := h.service.GetAllOwners()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, owners)
}

// GET /owners/store/:store_id
func (h *Handler) ListByStore(c *gin.Context) {
	storeID, err := strconv.Atoi(c.Param("store_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid store ID"})
		return
	}

	owners, err := h.service.GetOwnerByStore(storeID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, owners)
}

// GET /owners/:id/store
func (h *Handler) GetStoreByOwnerID(c *gin.Context) {
	ownerID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid owner ID"})
		return
	}

	store, err := h.service.GetStoreByOwnerID(ownerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, store)
}
