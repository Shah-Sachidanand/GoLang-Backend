package v1

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func InitializeRoutes(router *gin.RouterGroup, client *mongo.Client) {
	// User Routes Group
	userGroup := router.Group("/user")
	RegisterUserRoutes(userGroup, client)

	// Product Routes Group
	productGroup := router.Group("/product")
	RegisterProductRoutes(productGroup, client)
}
