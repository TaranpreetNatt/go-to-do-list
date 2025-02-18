package tasks

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"testing"
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

func createCSVWriter(w io.WriteSeeker) (*csv.Writer, error) {
	if w == nil {
		return nil, fmt.Errorf("writer is nil")
	}

	if _, err := w.Seek(0, io.SeekEnd); err != nil {
		return nil, fmt.Errorf("Failed to seek to start: %v", err)
	}

	writer := csv.NewWriter(w)
	writer.UseCRLF = true
	return writer, nil
}

func InitCsvFile(w io.WriteSeeker) error {
	if _, err := w.Seek(0, io.SeekStart); err != nil {
		return fmt.Errorf("failed to seek to start: %w", err)
	}

	writer, writerErr := createCSVWriter(w)
	if writerErr != nil {
		return writerErr
	}

	csvWriteErr := writer.Write([]string{"ID", "Task", "Done"})
	if csvWriteErr != nil {
		return fmt.Errorf("Error writing header to csv, %w", csvWriteErr)
	}
	writer.Flush()

	return nil
}

func GetHighestTaskID(r io.ReadSeeker) (int, error) {
	reader, readerErr := createCSVReader(r)
	if readerErr != nil {
		return -1, readerErr
	}

	if _, err := reader.Read(); err != nil {
		return -1, fmt.Errorf("failed to read header: %w", err)
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
	maxId, err := GetHighestTaskID(r)
	if err != nil {
		return nil, err
	}

	return &Task{ID: maxId + 1, Task: task, Done: false}, nil
}

func CreateTask(w io.WriteSeeker, task *Task) error {
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

func GetTasks(r io.ReadSeeker, includeHeader bool) ([][]string, error) {
	csvReader, readerErr := createCSVReader(r)
	if readerErr != nil {
		return nil, readerErr
	}

	if !includeHeader {
		if _, err := csvReader.Read(); err != nil {
			return nil, fmt.Errorf("failed to read header: %w", err)
		}
	}

	tasks, csvReadErr := csvReader.ReadAll()
	if csvReadErr != nil {
		return nil, csvReadErr
	}
	return tasks, nil
}

func CreateTempCSV(t *testing.T, data [][]string) (*os.File, error) {
	t.Helper()

	file, err := os.CreateTemp("", "task-test-*.csv")
	if err != nil {
		return nil, err
	}

	initCsvErr := InitCsvFile(file)
	if initCsvErr != nil {
		return nil, initCsvErr
	}

	writer := csv.NewWriter(file)
	csvErr := writer.WriteAll(data)

	if csvErr != nil {
		return nil, csvErr
	}

	return file, nil
}
