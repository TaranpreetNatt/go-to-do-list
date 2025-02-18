package view

import (
	"fmt"
	"io"

	"github.com/rivo/tview"
	tasks "github.com/taranpreetnatt/todo/internal/tasks"
)

func ViewTasks(r io.ReadSeeker) error {
	taskData, getTasksErr := tasks.GetTasks(r, true)
	if getTasksErr != nil {
		return fmt.Errorf("Erorr getting tasks in ViewTasks: %v", getTasksErr)
	}
	app := tview.NewApplication()
	table := tview.NewTable().SetBorders(true)

	for row, task := range taskData {
		for col, property := range task {
			tableCell := tview.NewTableCell(property)
			table.SetCell(row, col, tableCell)
		}
	}

	if err := app.SetRoot(table, true).Run(); err != nil {
		panic(err)
	}

	return nil
}
