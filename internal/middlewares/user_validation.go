package middlewares

import "github.com/gin-gonic/gin"

func UserValidationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("user-uuid", "1858f3ab-9cb3-4b15-96b6-c4407cc7cfc7")
		c.Next()
	}
}
