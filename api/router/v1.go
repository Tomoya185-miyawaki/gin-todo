/*
ルーティング用パッケージ
*/
package router

import (
	"github.com/Tomoya185-miyawaki/gin-todo/adapter/controller/todo"
	cors "github.com/Tomoya185-miyawaki/gin-todo/middleware"
	"github.com/gin-gonic/gin"
)

// ルーティング処理を行う
func Bind() *gin.Engine {
	router := gin.Default()
	// corsの設定
	router.Use(cors.NewCorsConfig())

	// v1apiのルーティング
	v1 := router.Group("/api/v1")
	{
		todoCtrl := todo.TodoController{}
		v1.GET("/todo/:id", todoCtrl.Show)
		v1.POST("/todo/create", todoCtrl.Create)
		v1.GET("/todos", todoCtrl.Index)
	}
	return router
}
