package tasks

import (
	"encoding/csv"
	"os"
	"testing"
)

func createTempCSV(t *testing.T, data [][]string) *os.File {
	file, err := os.CreateTemp("", "task_*")

	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}

	tempFileWriter := csv.NewWriter(file)

	writeError := tempFileWriter.WriteAll(data)

	if writeError != nil {
		t.Fatalf("Error writing data to the file: %v", writeError)
	}

	file.Close()
	return file
}
