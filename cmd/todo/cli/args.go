package cli

import (
	"flag"
	"fmt"
	"os"

	tasks "github.com/taranpreetnatt/todo/internal/tasks"
)

func GetArgs(args []string, file *os.File) (string, error) {
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

	flag.Parse()

	return "", nil
}
