package server

import (
	"github.com/gin-gonic/gin"
	"github.com/luis-valverde120/Zyrex-Backend-Server/internal/app/handlers/product"
)

type HandlerModules struct {
	Product *product.Handler // Ensure the Handler type is defined in the product package
	// add more modules here
}

func RegisterRoutes(router *gin.Engine, h *HandlerModules) {
	api := router.Group("/api")

	product.RegisterRoutes(api, h.Product)
}
