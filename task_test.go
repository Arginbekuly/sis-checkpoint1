package main

import (
	"errors"
	"testing"
)

func TestTask_Validate(t *testing.T) {
	tests := []struct {
		name      string
		id        string
		input     TaskInput
		expectErr bool
	}{
		{
			name:      "Valid input text",
			id:        "task-1",
			input:     TaskInput{Text: "Hello Go world"},
			expectErr: false,
		},
		{
			name:      "Empty input text",
			id:        "task-2",
			input:     TaskInput{Text: ""},
			expectErr: true,
		},
		{
			name:      "Whitespace only input text",
			id:        "task-3",
			input:     TaskInput{Text: "    \t \n  "},
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			task, err := NewTask(tt.id, tt.input)
			if tt.expectErr {
				if err == nil {
					t.Fatalf("expected error for input %q, got nil", tt.input.Text)
				}
				if !errors.Is(err, ErrEmptyInput) {
					t.Errorf("expected ErrEmptyInput, got %v", err)
				}
			} else {
				if err != nil {
					t.Fatalf("unexpected error: %v", err)
				}
				if task.State != StatePending {
					t.Errorf("expected initial state %s, got %s", StatePending, task.State)
				}
			}
		})
	}

}
func TestTaskProcessingAndLifecycle(t *testing.T) {
	task, err := NewTask("task-exec", TaskInput{Text: "Go is awesome"})
	if err != nil {
		t.Fatalf("failed to create task: %v", err)
	}

	if err := task.Process(); err != nil {
		t.Fatalf("Process() returned unexpected error: %v", err)
	}

	if task.State != StateCompleted {
		t.Errorf("expected state %s, got %s", StateCompleted, task.State)
	}

	if task.Result == nil {
		t.Fatal("expected result to be non-nil")
	}
	if task.Result.WordCount != 3 {
		t.Errorf("expected 3 words, got %d", task.Result.WordCount)
	}
	if task.Result.RuneCount != 13 {
		t.Errorf("expected 13 runes, got %d", task.Result.RuneCount)
	}
	if task.Result.SpaceCount != 2 {
		t.Errorf("expected 2 spaces, got %d", task.Result.SpaceCount)
	}

	err = task.Process()
	if !errors.Is(err, ErrInvalidTransition) {
		t.Errorf("expected ErrInvalidTransition on re-processing, got %v", err)
	}
}

func TestRepositoryOperations(t *testing.T) {
	repo := NewInMemoryTaskRepository()

	task1, _ := NewTask("task-1", TaskInput{Text: "First text"})
	task2, _ := NewTask("task-2", TaskInput{Text: "Second text with more words"})

	_ = task1.Process()

	_ = repo.Save(task1)
	_ = repo.Save(task2)

	retrieved1, err := repo.GetByID("task-1")
	if err != nil {
		t.Fatalf("unexpected error finding task-1: %v", err)
	}
	if retrieved1.State != StateCompleted {
		t.Errorf("expected task-1 state to be COMPLETED, got %s", retrieved1.State)
	}

	retrieved2, err := repo.GetByID("task-2")
	if err != nil {
		t.Fatalf("unexpected error finding task-2: %v", err)
	}
	if retrieved2.State != StatePending {
		t.Errorf("expected task-2 state to be PENDING, got %s", retrieved2.State)
	}
	if retrieved2.Result != nil {
		t.Error("expected task-2 result to be nil")
	}

	_, err = repo.GetByID("non-existent-id")
	if !errors.Is(err, ErrTaskNotFound) {
		t.Errorf("expected ErrTaskNotFound, got %v", err)
	}
}
