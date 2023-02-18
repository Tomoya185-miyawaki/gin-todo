package query

import "github.com/Tomoya185-miyawaki/gin-todo/domain/model"

type TodoQuery interface {
	FindAll() *model.Todos
}
