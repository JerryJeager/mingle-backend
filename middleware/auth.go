package middleware

import (
	"fmt"
	"net/http"

	auth "github.com/JerryJeager/mingle-backend/http"
	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		id, err := auth.ValidateToken(c)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"Unauthorized": "Authentication required"})
			fmt.Println(err)
			c.Abort()
			return
		}

		c.Set("user_id", id)

		c.Next()
	}
}
