package handler

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"go_blog/database"
	"net/http"
)

func AuthMiddleware(dbUser *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.Request.Header.Get("Authorization")
		user := database.V3GetUserByAuthHeader(dbUser, auth)
		if user.Email == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"msg": "invalid token",
			})
			c.Abort()
			return
		}
		c.Set("user", user)
		c.Next()
	}
}
