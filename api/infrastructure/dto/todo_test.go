package dto

import (
	"testing"

	"github.com/Tomoya185-miyawaki/gin-todo/domain/model"
	"github.com/stretchr/testify/assert"
)

func TestTodoConvertToModel(t *testing.T) {
	t.Parallel() // 並行実行
	todo := &Todo{
		ID:    1,
		Title: "テストTODO",
	}

	result := todo.ConvertToModel()
	expected := &model.Todo{
		Id:    model.TodoId(todo.ID),
		Title: model.TodoTitle(todo.Title),
	}
	assert.Equal(t, *result, *expected)
}

func TestTodosConvertToModel(t *testing.T) {
	t.Parallel() // 並行実行
	todo1 := &Todo{
		ID:    1,
		Title: "テストTODO1",
	}
	todo2 := &Todo{
		ID:    2,
		Title: "テストTODO2",
	}
	todos := &Todos{*todo1, *todo2}

	results := *todos.ConvertToModel()
	expected := &model.Todos{
		model.Todo{
			Id:    model.TodoId(todo1.ID),
			Title: model.TodoTitle(todo1.Title),
		},
		model.Todo{
			Id:    model.TodoId(todo2.ID),
			Title: model.TodoTitle(todo2.Title),
		},
	}

	assert.Equal(t, results, *expected)
}
