package v1

import (
	"github.com/gin-gonic/gin"

	"learning-golang/app/config"
)

func InitializeRoutes(router *gin.RouterGroup, client *config.Resource) {
	// User Routes Group
	userGroup := router.Group("/user")
	RegisterUserRoutes(userGroup, client)

	// Product Routes Group
	productGroup := router.Group("/product")
	RegisterProductRoutes(productGroup, client)
}
