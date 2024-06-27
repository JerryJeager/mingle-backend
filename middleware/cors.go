package middleware

import (
	"os"
	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	var path string
	if os.Getenv("ENVIRONMENT") == "development"{
		path = "http://localhost:3000"
	}else{
		path = "https://we-mingle.vercel.app"
	}
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", path)
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
