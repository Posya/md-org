package main

import (
	"fmt"
	"log"
	"os"

	"action"
)

var trace, info *log.Logger

func init() {
	trace = log.New(os.Stdout, "Trace: ", log.Ldate|log.Ltime|log.Lshortfile)
	info = log.New(os.Stdout, "Info: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func main() {
	fmt.Println("Hello, World!")
}

func do(act action.Action) error {
	switch act {
	case action.List:
		trace.Println("This is list action")
	case action.Check:
		trace.Println("This in check action")

	}

	return nil
}
