package main

import (
	"fmt"
	"os"
)

func initCSVFile(name string) (*os.File, error) {
	file, err := os.OpenFile(name, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)

	if err != nil {
		return nil, fmt.Errorf("initializing the CSV file failed with: %w", err)
	}
	return file, nil
}
