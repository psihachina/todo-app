package service

import (
	"github.com/psihachina/todo-app"
	"github.com/psihachina/todo-app/pkg/repository"
)

// Authorization - ...
type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
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
	Create(userId int, listId int, item todo.TodoItem) (int, error)
}

// Service - ...
type Service struct {
	Authorization
	TodoList
	TodoItem
}

// NewService - ...
func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		TodoList:      NewTodoListService(repos.TodoList),
	}
}
