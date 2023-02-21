package dto

import (
	"testing"

	"github.com/Tomoya185-miyawaki/gin-todo/domain/model"
)

func TestTodoConvertToModel(t *testing.T) {
	t.Parallel() // 並行実行
	todo := &Todo{
		ID:    1,
		Title: "テストTODO",
	}

	result := *todo.ConvertToModel()
	if result.Id != 1 {
		t.Error("IDの値が異なります")
	}
	if result.Title != "テストTODO" {
		t.Error("TITLEの値が異なります")
	}
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
	expected1 := &model.Todo{
		Id:    model.TodoId(todo1.ID),
		Title: model.TodoTitle(todo1.Title),
	}
	expected2 := &model.Todo{
		Id:    model.TodoId(todo2.ID),
		Title: model.TodoTitle(todo2.Title),
	}

	if results[0] != *expected1 {
		t.Error("1番目のスライスの値が異なります")
	}
	if results[1] != *expected2 {
		t.Error("2番目のスライスの値が異なります")
	}
}
