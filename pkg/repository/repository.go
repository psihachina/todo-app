package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/psihachina/todo-app"
)

// Authorization - ...
type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GetUser(username, password string) (todo.User, error)
}

// TodoList - ...
type TodoList interface {
	Create(userId int, list todo.TodoList) (int, error)
	GetAll(userId int) ([]todo.TodoList, error)
	GetById(userId int, listId int) (todo.TodoList, error)
	Delete(userId int, listId int) error
	Update(userId int, listId int, input todo.UpdateListInput) error
}

// TodoItem - ...
type TodoItem interface {
	Create(listId int, item todo.TodoItem) (int, error)
}

// Repository - ...
type Repository struct {
	Authorization
	TodoList
	TodoItem
}

// NewRepository - ...
func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		TodoList:      NewTodoListPostgres(db),
	}
}
