package todo

import (
	"github.com/Tomoya185-miyawaki/gin-todo/domain/model"
	"github.com/Tomoya185-miyawaki/gin-todo/infrastructure"
	"github.com/Tomoya185-miyawaki/gin-todo/infrastructure/dto"
	"github.com/Tomoya185-miyawaki/gin-todo/query"
)

type todoQueryImpl struct{}

func NewTodoQueryImpl() query.TodoQuery {
	return &todoQueryImpl{}
}

func (repo todoQueryImpl) FindAll() *model.Todos {
	db := infrastructure.GetDB()
	todoDaos := dto.Todos{}

	db.Find(&todoDaos)

	return todoDaos.ConvertToModel()
}
