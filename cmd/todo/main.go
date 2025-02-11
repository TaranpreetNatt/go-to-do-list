package main

import (
	"log"
	"os"

	cli "github.com/taranpreetnatt/todo/cmd/todo/cli"
)

func main() {
	_, err := cli.GetArgs(os.Args)

	if err != nil {
		log.Fatal(err)
	}
}
