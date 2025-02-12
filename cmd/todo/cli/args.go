package cli

import (
	"flag"
	"os"
)

func GetArgs(args []string, file *os.File) (string, error) {
	os.Args = args

	// createValue := flag.String("create", "", "Create a task")

	flag.Func("create", "Create a task", func(s string) error {
		return nil
	})

	flag.Parse()

	// if *createValue == "" {
	// 	return "", errors.New("Cannot create a task with an empty string")
	// }
	//
	// return *createValue, nil
	return "", nil
}
