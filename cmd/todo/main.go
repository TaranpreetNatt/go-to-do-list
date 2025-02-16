package main

import (
	"fmt"
	"os"

	cli "github.com/taranpreetnatt/todo/cmd/todo/cli"
)

func main() {
	file, initCsvErr := initCsvFile("todo.csv")
	defer file.Close()

	if initCsvErr != nil {
		fmt.Errorf("Error creating tasks: %w", initCsvErr)
	}
	_, err := cli.GetArgs(os.Args, file)
	if err != nil {
		fmt.Errorf("Error in the GetArgs function in main: %w", err)
	}
}
