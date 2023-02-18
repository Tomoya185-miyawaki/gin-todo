package model

type TodoId int
type TodoTitle string

type Todo struct {
	Id    TodoId    `json:"todoId"`
	Title TodoTitle `json:"todoTitle"`
}

type Todos []Todo
