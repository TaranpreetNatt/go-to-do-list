package tasks

import (
	"encoding/csv"
	"io"
	"strconv"
)

type Task struct {
	ID   int
	Task string
	Done bool
}

func CreateTask(w io.Writer, task Task) error {
	writer := csv.NewWriter(w)
	writer.UseCRLF = true

	err := writer.Write([]string{strconv.Itoa(task.ID), task.Task, strconv.FormatBool(task.Done)})
	if err != nil {
		return err
	}

	writer.Flush()
	return nil
}
