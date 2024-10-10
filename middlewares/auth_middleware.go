package middlewares

import (
	"go-gin-restapi-boilerplate/errorhandler"
	"go-gin-restapi-boilerplate/helpers"
	"go-gin-restapi-boilerplate/initializers"
	"go-gin-restapi-boilerplate/models"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		secretKey := os.Getenv("JWT_SECRET")

		bearerToken := c.GetHeader("Authorization")
		tokenStr := ""

		parts := strings.Split(bearerToken, " ")

		if len(parts) == 2 && parts[0] == "Bearer" {
			tokenStr = parts[1]
		} else {
			errorhandler.ErrorHandler(c, &errorhandler.UnauthorizedError{
				Message: "Invalid Access token format",
			})
			c.Abort()
			return
		}

		if tokenStr == "" {
			errorhandler.ErrorHandler(c, &errorhandler.UnauthorizedError{
				Message: "Access token not specified",
			})
			c.Abort()
			return
		}

		userID, err := helpers.ValidateToken(tokenStr, []byte(secretKey))

		if err != nil {
			errorhandler.ErrorHandler(c, &errorhandler.UnauthorizedError{
				Message: err.Error(),
			})
			c.Abort()
			return
		}

		var user models.User
		initializers.DB.Find(&user, userID)

		c.Set("userID", userID)
		c.Next()
	}
}
