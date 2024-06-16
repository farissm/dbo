package auth

import (
	"strings"
	"net/http"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"

	"dbo/helper"
	"dbo/model"
)

// type auth struct {
// 	db *gorm.DB
// }

// func NewAuth(db *gorm.DB) *auth {
// 	return &auth{db}
// }

func AuthMiddleware(db *gorm.DB) gin.HandlerFunc{
	return func (c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
	
		if !strings.Contains(authHeader, "Bearer") {
			response := helper.APIResponse(http.StatusUnauthorized, "Invalid token")
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
	
		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}
	
		token, err := helper.ValidateToken(tokenString)
		if err != nil {
			response := helper.APIResponse(http.StatusUnauthorized, "Unauthorized")
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)

		if !ok || !token.Valid {
			response := helper.APIResponse(http.StatusUnauthorized, "Unauthorized")
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return			
		}

		username := claim["Username"].(string)

		var customer model.Customer
		err = db.Where("username = ?", username).First(&customer).Error
		if err != nil {
			response := helper.APIResponse(http.StatusUnauthorized, "Unauthorized")
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return			
		}

		c.Set("currentUser", customer)
		c.Next()
	}
} 