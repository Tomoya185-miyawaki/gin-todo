/*
dto用のパッケージ
*/
package dto

import (
	"time"

	"github.com/Tomoya185-miyawaki/gin-todo/domain/model"
)

type Todo struct {
	ID        int       `json:"id" binding:"required"`
	Title     string    `json:"title" binding:"required"`
	CreateAt  time.Time `gorm:"default:current_timestamp"`
	UpdatedAt time.Time `gorm:"default:current_timestamp"`
}

type Todos []Todo

// todoのstructをmodelに変換する
func (todo *Todo) ConvertToModel() *model.Todo {
	return &model.Todo{
		Id:    model.TodoId(todo.ID),
		Title: model.TodoTitle(todo.Title),
	}
}

// todosのstructをmodelに変換する
func (todos Todos) ConvertToModel() *model.Todos {
	result := make(model.Todos, len(todos))

	for idx, todo := range todos {
		todoModel := model.Todo{
			Id:    model.TodoId(todo.ID),
			Title: model.TodoTitle(todo.Title),
		}
		result[idx] = todoModel
	}

	return &result
}
