package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Check if the user is logged in
		username, err := c.Cookie("username")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "You are not logged in"})
			c.Abort()
			return
		}
		if username == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "You are not logged in"})
			c.Abort()
			return
		}

		// The user is logged in, proceed to the route handler
		c.Next()
	}
}
