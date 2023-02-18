package todo

import (
	"github.com/Tomoya185-miyawaki/gin-todo/helpers"
	"github.com/Tomoya185-miyawaki/gin-todo/infrastructure/queryImpl/todo"
	"github.com/gin-gonic/gin"
)

type TodoController struct{}

func (controller TodoController) Index(c *gin.Context) {
	result := todo.NewTodoQueryImpl().FindAll()
	helpers.Response(c, result, nil)
}
