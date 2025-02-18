package main

import (
	"fmt"
	"os"

	cli "github.com/taranpreetnatt/todo/cmd/todo/cli"
	tasks "github.com/taranpreetnatt/todo/internal/tasks"
)

const csvFileName = "todo.csv"

func main() {
	file, iniFileErr := initFile(csvFileName)
	defer file.Close()

	if iniFileErr != nil {
		fmt.Errorf("Error creating tasks: %w", iniFileErr)
	}

	fileInfo, fileStatErr := file.Stat()
	if fileStatErr != nil {
		fmt.Errorf("Error getting fileInfo: %v", fileStatErr)
	}

	if fileInfo.Size() == 0 {
		initCsvErr := tasks.InitCsvFile(file)
		if initCsvErr != nil {
			fmt.Errorf("%v", initCsvErr)
		}
	}
	err := cli.GetArgs(os.Args, file)
	if err != nil {
		fmt.Errorf("Error in the GetArgs function in main: %w", err)
	}
}
