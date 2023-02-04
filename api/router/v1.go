package router

import (
	"github.com/gin-gonic/gin"
)

func Bind() {
	router := gin.Default()
	v1 := router.Group("/api/v1")
	{
		v1.GET("/todos", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "Hello aaa",
			})
		})
	}
	router.Run(":3000")
}
