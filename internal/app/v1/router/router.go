package router

import (
	"api-template/internal/app/v1/handlers/product"
	"api-template/internal/app/v1/handlers/user"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// API v1 routes
	r.GET("/api/v1/user", user.GetUser)
	r.GET("/api/v1/product", product.GetProduct)

	return r
}
