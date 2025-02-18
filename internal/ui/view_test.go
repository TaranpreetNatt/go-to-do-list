package view_test

import (
	"os"
	"testing"

	tasks "github.com/taranpreetnatt/todo/internal/tasks"
	ui "github.com/taranpreetnatt/todo/internal/ui"
)

func TestViewTasks(t *testing.T) {
	taskTests := []struct {
		name     string
		taskData [][]string
	}{
		{
			name:     "one task",
			taskData: [][]string{{"1", "Task one", "false"}},
		},
		{
			name:     "two tasks",
			taskData: [][]string{{"1", "Task one", "false"}, {"2", "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.", "false"}},
		},
		// {
		// 	name:        "no tasks",
		// 	taskData:    [][]string{},
		// },
	}

	for _, tt := range taskTests {
		t.Run(tt.name, func(t *testing.T) {
			file, createTempErr := tasks.CreateTempCSV(t, tt.taskData)
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
				t.Fatalf("Error creating temp file in TestGetTasks: %v", createTempErr)
			}

			viewTasksErr := ui.ViewTasks(file)
			if viewTasksErr != nil {
				t.Fatalf("Error viewing tasks: %v", viewTasksErr)
			}
		})
	}
}
