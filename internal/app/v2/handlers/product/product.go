package product

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetProduct(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Get Product from v2"})
}
