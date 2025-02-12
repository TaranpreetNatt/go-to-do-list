package tasks_test

import (
	"bytes"
	"testing"

	tasks "github.com/taranpreetnatt/todo/internal/tasks"
)

func TestCreatingTasks(t *testing.T) {
	taskTests := []struct {
		name     string
		Task     tasks.Task
		expected string
	}{
		{
			name:     "task with one word",
			Task:     tasks.Task{ID: 1, Task: "Task", Done: false},
			expected: "1,Task,false\r\n",
		},
		{
			name:     "task with two words",
			Task:     tasks.Task{ID: 2, Task: "Task word", Done: false},
			expected: "2,Task word,false\r\n",
		},
		{
			name:     "Empty task",
			Task:     tasks.Task{ID: 3, Task: "", Done: true},
			expected: "3,,true\r\n",
		},
	}

	for _, tt := range taskTests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			err := tasks.CreateTask(&buf, tt.Task)

			if err != nil {
				t.Fatalf("Error testing creating a task: %v", err)
			}
			assertTask(t, buf.String(), tt.expected)
		})
	}
}

func assertTask(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}

}
