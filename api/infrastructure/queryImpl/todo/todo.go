/*
todoのリポジトリ（実態）用パッケージ
*/
package todo

import (
	"time"

	"github.com/Tomoya185-miyawaki/gin-todo/domain/model"
	"github.com/Tomoya185-miyawaki/gin-todo/errors"
	"github.com/Tomoya185-miyawaki/gin-todo/infrastructure"
	"github.com/Tomoya185-miyawaki/gin-todo/infrastructure/dto"
	"github.com/Tomoya185-miyawaki/gin-todo/query"
)

type todoQueryImpl struct{}

// structを生成する
func NewTodoQueryImpl() query.TodoQuery {
	return &todoQueryImpl{}
}

// todoを全件取得する
func (repo todoQueryImpl) FindAll() *model.Todos {
	db := infrastructure.GetDB()
	todoDaos := dto.Todos{}

	db.Find(&todoDaos)

	return todoDaos.ConvertToModel()
}

// idをキーにして、todoを取得する
func (repo todoQueryImpl) FindById(id int) (*model.Todo, *errors.AppError) {
	db := infrastructure.GetDB()
	todoDao := dto.Todo{ID: id}

	if db.Find(&todoDao).RecordNotFound() {
		err := errors.NewAppError("id " + string(id) + "のtodoは見つかりませんでした")
		return nil, &err
	}

	return todoDao.ConvertToModel(), nil
}

// todoを作成する
func (repo todoQueryImpl) Create(title string) (*string, *errors.AppError) {
	db := infrastructure.GetDB()
	todoDao := dto.Todo{Title: title, CreateAt: time.Now(), UpdatedAt: time.Now()}

	result := db.Create(&todoDao)

	if result.Error != nil {
		err := errors.NewAppError(result.Error.Error())
		return nil, &err
	}

	successMessage := "todoの作成に成功しました"
	return &successMessage, nil
}
