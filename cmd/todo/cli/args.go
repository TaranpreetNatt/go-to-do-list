package cli

import (
	"errors"
	"flag"
	"os"
)

func GetArgs(args []string) (string, error) {
	os.Args = args

	createValue := flag.String("create", "", "Create a task")

	flag.Parse()

	if *createValue == "" {
		return "", errors.New("Cannot create a task with an empty string")
	}

	return *createValue, nil
}
