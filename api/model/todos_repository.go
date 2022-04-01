package repository

type TodoRepository interface {
	GetTodos() (todos string, err error)
}

type todoRepository struct{}

func NewTodoRepository() TodoRepository {
	return &todoRepository{}
}

func (tr *todoRepository) GetTodos() (todos string, err error) {
	todos = "hello"
	return
}
