package tasks

import (
	"encoding/csv"
	"fmt"
	"io"
	"strconv"
)

type Task struct {
	ID   int
	Task string
	Done bool
}

func createCSVReader(r io.ReadSeeker) (*csv.Reader, error) {
	if r == nil {
		return nil, fmt.Errorf("reader is nil")
	}
	if _, err := r.Seek(0, io.SeekStart); err != nil {
		return nil, fmt.Errorf("failed to seek to start: %w", err)
	}

	reader := csv.NewReader(r)
	reader.ReuseRecord = true
	return reader, nil
}

func createCSVWriter(w io.Writer) (*csv.Writer, error) {
	if w == nil {
		return nil, fmt.Errorf("writer is nil")
	}

	writer := csv.NewWriter(w)
	writer.UseCRLF = true
	return writer, nil
}

func GetHighestTaskID(r io.ReadSeeker) (int, error) {
	reader, readerErr := createCSVReader(r)
	if readerErr != nil {
		return -1, readerErr
	}

	tasks, err := reader.ReadAll()
	if err != nil {
		return -1, err
	}

	maxId := 0
	for _, task := range tasks {
		id, err := strconv.Atoi(task[0])
		if err != nil {
			return -1, err
		}

		if id > maxId {
			maxId = id
		}
	}

	return maxId, nil
}

func NewTask(r io.ReadSeeker, task string) (*Task, error) {
	if r == nil {
		return nil, fmt.Errorf("Reader is nil")
	}

	if _, err := r.Seek(0, io.SeekStart); err != nil {
		return nil, fmt.Errorf("failed to seek to start: %w", err)
	}

	maxId, err := GetHighestTaskID(r)
	if err != nil {
		return nil, err
	}

	return &Task{ID: maxId + 1, Task: task, Done: false}, nil
}

func CreateTask(w io.Writer, task *Task) error {
	writer, writerErr := createCSVWriter(w)
	if writerErr != nil {
		return writerErr
	}

	err := writer.Write([]string{strconv.Itoa(task.ID), task.Task, strconv.FormatBool(task.Done)})
	if err != nil {
		return err
	}

	writer.Flush()
	return nil
}
