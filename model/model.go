package model

import "time"

type Todo struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	CreateAt  time.Time `json:"creaeted_at"`
}

type DBHandler interface {
	GetTodoList() []*Todo
	AddTodo(name string) *Todo
	RemoveTodo(id int) bool
	CompleteTodo(id int, complete bool) bool
	Close()
}

func NewDBHandler() DBHandler {
	return newSqliteHandler()
}
