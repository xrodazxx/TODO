package service

import "github.com/xrodazxx/TODO/pkg/repository"

type Authorization interface {
	// Добавьте методы для интерфейса Authorization
}

type TodoList interface {
	// Добавьте методы для интерфейса TodoList
}

type TodoItem interface {
	// Добавьте методы для интерфейса TodoItem
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		// Инициализируйте сервисы, если нужно
		// Authorization: ...,
		// TodoList: ...,
		// TodoItem: ...,
	}
}
