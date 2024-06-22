package cmd

import (
	"log"
	"os"

	"github.com/JerryJeager/mingle-backend/api"
	"github.com/JerryJeager/mingle-backend/manualwire"
	"github.com/JerryJeager/mingle-backend/middleware"
	"github.com/gin-gonic/gin"
)

var userController = manualwire.GetUserController()

func ExecuteApiRoutes() {

	r := gin.Default()
	r.Use(middleware.CORSMiddleware())

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello!",
		})
	})

	v1 := r.Group("/api/v1")

	v1.GET("/info/openapi.yaml", func(c *gin.Context) {
		c.String(200, api.OpenApiDocs())
	})

	googleUser := v1.Group("users/google")

	googleUser.GET("/signup", userController.CreateUserWithGoogle)

	users := v1.Group("/user")
	users.POST("/signup", userController.CreateUser)
	users.POST("/login", userController.CreateToken)

	user := v1.Group("/user")
	user.Use(middleware.JwtAuthMiddleware())
	user.GET("", middleware.JwtAuthMiddleware(), userController.GetUser)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	if err := r.Run(":" + port); err != nil {
		log.Panicf("error: %s", err)
	}

}
