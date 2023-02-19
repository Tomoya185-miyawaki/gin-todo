/*
todoのcontroller用のパッケージ
*/
package todo

import (
	"strconv"

	"github.com/Tomoya185-miyawaki/gin-todo/errors"
	"github.com/Tomoya185-miyawaki/gin-todo/helpers"
	"github.com/Tomoya185-miyawaki/gin-todo/infrastructure/queryImpl/todo"
	"github.com/Tomoya185-miyawaki/gin-todo/request"
	"github.com/gin-gonic/gin"
)

type TodoController struct{}

// todo一覧を返却する
func (controller TodoController) Index(c *gin.Context) {
	result := todo.NewTodoQueryImpl().FindAll()
	helpers.Response(c, result, nil)
}

// idをキーにして、todoを返却する
func (controller TodoController) Show(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	result, err := todo.NewTodoQueryImpl().FindById(id)
	if err.HasErrors() {
		helpers.Response(c, nil, err)
	}
	helpers.Response(c, result, nil)
}

func (controller TodoController) Create(c *gin.Context) {
	var todoRequest request.Todo
	if err := c.Bind(&todoRequest); err != nil {
		errors := errors.NewAppError(err.Error())
		helpers.Response(c, nil, &errors)
	}
	result, err := todo.NewTodoQueryImpl().Create(todoRequest.Title)
	if err.HasErrors() {
		helpers.Response(c, nil, err)
	}
	helpers.Response(c, result, nil)
}
