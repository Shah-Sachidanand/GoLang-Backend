package app

import (
	"context"
	"learning-golang/config"
	v1 "learning-golang/internal/routes/v1"
	v2 "learning-golang/internal/routes/v2"
	"learning-golang/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func StartGin() {
	// Load configuration
	cfg := config.LoadConfig()

	// Connect to MongoDB
	client := config.ConnectMongo(cfg.MongoURI)
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	// Initialize Gin router
	router := gin.Default()

	router.Use(gin.Logger())
	// router.GET("swagger/*any", middlewares.NewSwagger())

	// Setup application routes
	v1.InitializeRoutes(router.Group("/api/v1"), client)
	v2.InitializeRoutes(router.Group("/api/v2"), client)

	// Load HTML templates from the views folder
	router.LoadHTMLGlob("app/views/*")

	// Route for home page
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home.html", nil)
	})

	// Custom 404 handler to return HTML for browsers and JSON for others
	router.NoRoute(func(c *gin.Context) {
		userAgent := c.Request.UserAgent()
		// Extract the method and path from the request
		method := c.Request.Method
		path := c.Request.URL.Path

		// Check if the User-Agent indicates a browser
		if utils.IsBrowser(userAgent) {
			c.HTML(http.StatusNotFound, "404.html", gin.H{
				"method": method,
				"path":   path,
			})
		} else {
			c.JSON(http.StatusNotFound, gin.H{
				"error":   "Not Found",
				"message": "The requested URL was not found on this server.",
				"method":  method,
				"path":    path,
			})
		}
	})

	// Start the server
	router.Run(":" + cfg.Port)
}
