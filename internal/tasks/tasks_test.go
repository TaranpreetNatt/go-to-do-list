package tasks_test

import (
	"bytes"
	"encoding/csv"
	"os"
	"testing"

	tasks "github.com/taranpreetnatt/todo/internal/tasks"
)

func TestNewTask(t *testing.T) {
	taskTests := []struct {
		name        string
		taskData    [][]string
		task        string
		expected    tasks.Task
		expectedErr bool
	}{
		{
			name:        "Happy path task",
			taskData:    [][]string{{"1", "Task one", "false"}, {"2", "Task two", "false"}, {"3", "Task three", "false"}},
			task:        "Task 4",
			expected:    tasks.Task{ID: 4, Task: "Task 4", Done: false},
			expectedErr: false,
		},
		{
			name:        "Empty file",
			taskData:    [][]string{},
			task:        "Task",
			expected:    tasks.Task{ID: 1, Task: "Task", Done: false},
			expectedErr: false,
		},
		{
			name:        "Random IDs",
			taskData:    [][]string{{"3", "Task one", "false"}, {"8", "two", "false"}, {"6", "Task three", "false"}},
			task:        "Task",
			expected:    tasks.Task{ID: 9, Task: "Task", Done: false},
			expectedErr: false,
		},
		{
			name:        "Invalid ID",
			taskData:    [][]string{{"asd", "Task one", "false"}, {"8", "two", "false"}, {"6", "Task three", "false"}},
			task:        "Task",
			expected:    tasks.Task{ID: 9, Task: "Task", Done: false},
			expectedErr: true,
		},
	}

	for _, tt := range taskTests {
		t.Run(tt.name, func(t *testing.T) {
			file, createTempErr := createTempCSV(t, tt.taskData)
			defer func() {
				err := file.Close()
				if err != nil {
					t.Fatalf("Error closing file in TestGetTaskID test: %v", err)
				}
				removeErr := os.Remove(file.Name())
				if removeErr != nil {
					t.Fatalf("Error removing temp file in TestGetTaskID: %v", removeErr)
				}
			}()

			if createTempErr != nil {
				t.Fatalf("Create temp file error: %v", createTempErr)
			}

			got, newTaskErr := tasks.NewTask(file, tt.task)
			if newTaskErr != nil && !tt.expectedErr {
				t.Fatalf("Error in new task: %v", newTaskErr)
			}

			if tt.expectedErr && newTaskErr == nil {
				t.Fatalf("Expected an error from NewTask: %v", newTaskErr)
			}

			if !tt.expectedErr {
				assertEqual(t, *got, tt.expected)
			}

		})
	}
}

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
			err := tasks.CreateTask(&buf, &tt.Task)

			if err != nil {
				t.Fatalf("Error testing creating a task: %v", err)
			}
			assertEqual(t, buf.String(), tt.expected)
		})
	}
}

func TestGetTaskID(t *testing.T) {
	taskTests := []struct {
		name        string
		taskData    [][]string
		expectedId  int
		expectedErr bool
	}{
		{
			name:        "Id of one task",
			taskData:    [][]string{{"1", "Task one", "false"}},
			expectedId:  1,
			expectedErr: false,
		},
		{
			name:        "Id of one task",
			taskData:    [][]string{{"2", "Task two", "false"}},
			expectedId:  2,
			expectedErr: false,
		},
		{
			name:        "Sequential task IDs",
			taskData:    [][]string{{"1", "Task one", "false"}, {"2", "Task two", "false"}, {"3", "Task three", "false"}},
			expectedId:  3,
			expectedErr: false,
		},
		{
			name:        "Non-sequential task IDs",
			taskData:    [][]string{{"2", "Task", "false"}, {"3", "Task", "false"}, {"1", "Task", "false"}},
			expectedId:  3,
			expectedErr: false,
		},
		{
			name:        "invalid task ID",
			taskData:    [][]string{{"invalidID", "Task", "false"}},
			expectedId:  3,
			expectedErr: true,
		},
	}

	for _, tt := range taskTests {
		t.Run(tt.name, func(t *testing.T) {
			file, createTempErr := createTempCSV(t, tt.taskData)
			defer func() {
				err := file.Close()
				if err != nil {
					t.Fatalf("Error closing file in TestGetTaskID test: %v", err)
				}

				removeErr := os.Remove(file.Name())
				if removeErr != nil {
					t.Fatalf("Error removing temp file in TestGetTaskID: %v", removeErr)
				}
			}()

			if createTempErr != nil {
				t.Fatalf("Error creating temp csv file in TestGetTaskID: %v", createTempErr)
			}

			got, err := tasks.GetHighestTaskID(file)
			if (err != nil) && (!tt.expectedErr) {
				t.Fatalf("Error getting TaskID in TestGetTaskID: %v", err)
			}

			if tt.expectedErr && err == nil {
				t.Fatalf("Expected an error, didn't get one %v", err)
			}

			if !tt.expectedErr {
				assertEqual(t, got, tt.expectedId)
			}
		})
	}
}

func assertEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func createTempCSV(t *testing.T, data [][]string) (*os.File, error) {
	t.Helper()

	file, err := os.CreateTemp("", "task-test-*.csv")
	if err != nil {
		return nil, err
	}

	writer := csv.NewWriter(file)
	csvErr := writer.WriteAll(data)

	if csvErr != nil {
		return nil, csvErr
	}

	return file, nil
}
