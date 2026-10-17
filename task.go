package main

import (
	"errors"
	"strings"
	"unicode/utf8"
)

type TaskState string

const (
	StatePending    TaskState = "PENDING"
	StateProcessing TaskState = "PROCESSING"
	StateCompleted  TaskState = "COMPLETED"
)

var (
	ErrEmptyInput        = errors.New("input text cannot be empty or whitespace only")
	ErrInvalidTransition = errors.New("cannot process a task that is already completed")
)

type TaskInput struct {
	Text string
}

type TaskResult struct {
	WordCount  int
	RuneCount  int
	SpaceCount int
}

type Task struct {
	ID     string
	Input  TaskInput
	State  TaskState
	Result *TaskResult
}

func NewTask(id string, input TaskInput) (*Task, error) {
	t := &Task{
		ID:    id,
		Input: input,
		State: StatePending,
	}

	if err := t.Validate(); err != nil {
		return nil, err
	}
	return t, nil
}

func (t *Task) Validate() error {
	if strings.TrimSpace(t.Input.Text) == "" {
		return ErrEmptyInput
	}
	return nil
}

func (t *Task) Process() error {
	if t.State == StateCompleted {
		return ErrInvalidTransition
	}

	t.State = StateProcessing

	words := strings.Fields(t.Input.Text)
	runes := utf8.RuneCountInString(t.Input.Text)
	spaces := strings.Count(t.Input.Text, " ")

	t.Result = &TaskResult{
		WordCount:  len(words),
		RuneCount:  runes,
		SpaceCount: spaces,
	}

	t.State = StateCompleted
	return nil
}
