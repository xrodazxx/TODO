package repository

type Authorization interface {
	// Добавьте методы для интерфейса
}

type TodoList interface {
	// Добавьте методы для интерфейса
}

type TodoItem interface {
	// Добавьте методы для интерфейса
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository() *Repository {
	return &Repository{}
}
