package main

import (
	"learning-golang/app"
)

// @title Swagger Example API
// @version 2.0
// @description This is a sample swagger
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1

// @securityDefinitions.apiKey BearerAuth
// @type Bearer Token
// @name Authorization
// @in header
// @description Bearer token for accessing the API
func main() {
	app.StartGin()
}
