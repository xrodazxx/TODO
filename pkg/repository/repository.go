package repository

import "github.com/xrodazxx/TODO/pkg/repository"

type Authorization interface {
}
type TodoList interface {
}
type TodoItem interface {
}
type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(repos repository.Repository) *Repository {
	return &Repository{}
}
