package cmd


import (
	"github.com/JerryJeager/mingle-backend/api"
	"github.com/gin-gonic/gin"
)

func ExecuteApiRoutes() {

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello!",
		})
	})


	v1 := r.Group("/api/v1")

	v1.GET("/info/openapi.yaml", func(c *gin.Context) {
		c.String(200, api.OpenApiDocs())
	})

}