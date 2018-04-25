package main

import (
	"os"
	"fmt"

	"github.com/jessevdk/go-flags"
)

var opts struct {
	File  []string `short:"f" long:"file" description:"files to parse"`
	Debug bool     `long:"debug" description:"debug flag"`
}

func main() {

	fmt.Println()
	parser := flags.NewParser(&opts, flags.Default)
	parser.AddCommand("list", "shows list of tasks and headers", "", &cmdList{})
	parser.AddCommand("agenda", "shows agenda", "", &cmdAgenda{})
	parser.AddCommand("add", "adds new task", "", &cmdAdd{})
	parser.AddCommand("done", "mark task as done / not done", "", &cmdDone{})
	parser.AddCommand("archive", "archive marked tasks", "", &cmdArchive{})
	parser.AddCommand("notify", "show notifications", "", &cmdNotify{})

	_, err := parser.Parse()
	if err != nil {
		os.Exit(1)
	}
}
