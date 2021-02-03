package main

import (
	"async-book-shelf/src/cmd"
	"async-book-shelf/src/failure"
	"os"
)

const (
	mode    = 1
	content = 2
)

func getArg(index int) string {
	if len(os.Args) <= index {
		return ""
	}
	return os.Args[index]
}

func main() {
	mode := getArg(mode)
	content := getArg(content)

	actions := map[string]func(string){
		"insert": cmd.Insert,
	}

	action, ok := actions[mode]

	if !ok {
		failure.Fail("Please provide a valid option: [insert]")
	}

	action(content)
}
