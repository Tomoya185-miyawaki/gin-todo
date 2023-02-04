package main

import (
	db "Tomoya185-miyawaki/gin-todo/infrastructure"
	"github.com/gin-gonic/gin"
)

func main() {
	db := db.Init()
	defer db.Close()
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello aaa",
		})
	})

	router.Run(":3000")
}
