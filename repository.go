package main

import "errors"

var ErrTaskNotFound = errors.New("task not found")

type TaskRepository interface {
	Save(task *Task) error
	GetByID(id string) (*Task, error)
}

type InMemoryTaskRepository struct {
	tasks map[string]*Task
}

var _ TaskRepository = (*InMemoryTaskRepository)(nil)

func NewInMemoryTaskRepository() *InMemoryTaskRepository {
	return &InMemoryTaskRepository{
		tasks: make(map[string]*Task),
	}
}

func (r *InMemoryTaskRepository) Save(task *Task) error {
	r.tasks[task.ID] = task
	return nil
}

func (r *InMemoryTaskRepository) GetByID(id string) (*Task, error) {
	task, ok := r.tasks[id]
	if !ok {
		return nil, ErrTaskNotFound
	}
	return task, nil
}
