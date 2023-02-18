package router

import (
	"github.com/Tomoya185-miyawaki/gin-todo/adapter/controller/todo"
	"github.com/gin-gonic/gin"
)

func Bind() {
	router := gin.Default()
	v1 := router.Group("/api/v1")
	{
		todoCtrl := todo.TodoController{}
		v1.GET("/todos", todoCtrl.Index)
	}
	router.Run(":3000")
}
