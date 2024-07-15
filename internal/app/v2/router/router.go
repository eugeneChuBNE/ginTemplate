package router

import (
	"api-template/internal/app/v2/handlers/product"
	"api-template/internal/app/v2/handlers/user"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// API v2 routes
	r.GET("/api/v2/user", user.GetUser)
	r.GET("/api/v2/product", product.GetProduct)

	return r
}
