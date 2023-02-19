/*
todoのリポジトリ（実態）用パッケージ
*/
package todo

import (
	"github.com/Tomoya185-miyawaki/gin-todo/domain/model"
	"github.com/Tomoya185-miyawaki/gin-todo/errors"
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

func (repo todoQueryImpl) FindById(id int) (*model.Todo, *errors.AppError) {
	db := infrastructure.GetDB()
	todoDao := dto.Todo{ID: id}

	if db.Find(&todoDao).RecordNotFound() {
		err := errors.NewAppError("id " + string(id) + "のtodoは見つかりませんでした")
		return nil, &err
	}

	return todoDao.ConvertToModel(), nil
}
