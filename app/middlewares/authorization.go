package middlewares

import (
	// "learning-golang/app/internal/repository"
	// "learning-golang/app/pkg/utils"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func GetRolesFromToken(tokenReq string) (role []string) {
	token, _ := jwt.Parse(tokenReq, func(token *jwt.Token) (interface{}, error) {
		return []byte(""), nil
	})
	claim := token.Claims.(jwt.MapClaims)
	var roles []string
	rolesResource := claim["roles"].([]interface{})
	for _, role := range rolesResource {
		roles = append(roles, role.(string))
	}
	if len(roles) <= 0 {
		return nil
	}
	return roles
}

// func RequireAuthorization(name string) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		username := utils.GetUsername(c)
// 		isAccessible := false
// 		userRole, _, err := repository.GetUserRole(username, name)
// 		if err != nil || userRole == nil {
// 			notPermission(c)
// 			return
// 		}
// 		access := userRole.Access

// 		method := c.Request.Method

// 		if utils.Contains(access, method) {
// 			isAccessible = true
// 		}

// 		if isAccessible == false {
// 			notPermission(c)
// 			return
// 		}
// 		c.Next()
// 	}
// }

func invalidRequest(c *gin.Context) {
	c.JSON(http.StatusForbidden, gin.H{"error": "Invalid request, restricted endpoint"})
	c.Abort()
}

func notPermission(c *gin.Context) {
	c.JSON(http.StatusForbidden, gin.H{"error": "Dont have permission"})
	c.Abort()
}
