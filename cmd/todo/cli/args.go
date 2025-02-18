package cli

import (
	"flag"
	"fmt"
	"os"

	tasks "github.com/taranpreetnatt/todo/internal/tasks"
	ui "github.com/taranpreetnatt/todo/internal/ui"
)

func GetArgs(args []string, file *os.File) error {
	os.Args = args

	flag.Func("create", "Create a task", func(s string) error {
		newTask, newTaskErr := tasks.NewTask(file, s)
		if newTaskErr != nil {
			return fmt.Errorf("Error creating new Task: %w", newTaskErr)
		}

		err := tasks.CreateTask(file, newTask)
		if err != nil {
			return fmt.Errorf("Error creating tasks: %w", err)
		}
		return nil
	})

	viewFlag := flag.Bool("view", false, "View tasks")

	flag.Parse()

	if *viewFlag {
		viewErr := ui.ViewTasks(file)
		if viewErr != nil {
			return fmt.Errorf("Error view tasks: %w", viewErr)
		}
	}

	return nil
}
