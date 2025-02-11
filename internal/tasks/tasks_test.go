package tasks_test

import (
	"bytes"
	"testing"

	tasks "github.com/taranpreetnatt/todo/internal/tasks"
)

func TestCreatingTasks(t *testing.T) {
	var buf bytes.Buffer

	task := tasks.Task{ID: 1, Task: "task", Done: false}
	err := tasks.CreateTask(&buf, task)

	if err != nil {
		t.Fatalf("Error testing creating a task: %v", err)
	}

	got := buf.String()
	want := "1,task,false\r\n"

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
