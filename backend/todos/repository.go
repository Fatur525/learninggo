package todos

import (
	"github.com/Fatur525/learninggo/model"
	"gorm.io/gorm"
)

type Repository interface {
	GetTodos() ([]model.Todos, error)
	CreateTodos(task string) (model.Todos, error)
	UpdateTodos(id int) (model.Todos, error)
	DeleteTodos(id int) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetTodos() ([]model.Todos, error) {
	var todos []model.Todos
	res := r.db.Find(&todos)
	if res.Error != nil {
		return nil, res.Error
	}
	return todos, nil
}

func (r *repository) CreateTodos(task string) (model.Todos, error) {
	todo := model.Todos{
		Task: task,
		Done: false,
	}

	res := r.db.Create(&todo)
	if res.Error != nil {
		return model.Todos{}, res.Error
	}

	return todo, nil
}

func (r *repository) UpdateTodos(id int) (model.Todos, error) {
	var todo model.Todos
	res := r.db.Where("id = ?", id).First(&todo).Update("Done", true)
	if res.Error != nil {
		return model.Todos{}, res.Error
	}

	return todo, nil
}

func (r *repository) DeleteTodos(id int) error {
	var todo model.Todos
	res := r.db.Delete(&todo, id)
	if res.Error != nil {
		return (res.Error)
	}

	return nil
}
