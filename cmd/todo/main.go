package main

import (
	"log"
	"os"

	cli "github.com/taranpreetnatt/todo/cmd/todo/cli"
)

func main() {
	file, initCSVErr := initCSVFile("todo.csv")

	if initCSVErr != nil {
		log.Fatal(initCSVErr)
	}
	_, err := cli.GetArgs(os.Args, file)
	if err != nil {
		log.Fatal(err)
	}
}
