package main

import (
	"fmt"
	"log"
	"os"
)

type action int

var trace, info *log.Logger

const (
	list action = iota
	check
)

func init() {
	trace = log.New(os.Stdout, "Trace: ", log.Ldate|log.Ltime|log.Lshortfile)
	info = log.New(os.Stdout, "Info: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func main() {
	fmt.Println("Hello, World!")
}

func do(act action) error {
	switch act {
	case list:
		trace.Println("This is list action")
	case check:
		trace.Println("This in check action")

	}

	return nil
}
