package main

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"
	"time"
)

type cmdAgenda struct {
	Verbose bool   `short:"v" long:"verbose" env:"MDORG_VERBOSE" description:"show all tags (local and inherited)"`
	From    string `short:"a" long:"from" description:"from date"`
	To      string `short:"b" long:"to" description:"to date"`
}

func (ca *cmdAgenda) Execute(args []string) error {
	dir, err := DirFromOptsOrCurrent()
	if err != nil {
		return err
	}

	if ca.To == "" {
		t := time.Now().AddDate(0, 0, 1)
		ca.To = t.Format("2006.01.02")
	}

	for _, file := range dir {
		fmt.Println("File: ", file)
		fmt.Println()

		w := tabwriter.NewWriter(os.Stdout, 0, 8, 1, '\t', 0)

		lines, err := ReadFile(file)
		if err != nil {
			return err
		}

		elements, err := parse(lines)
		if err != nil {
			return err
		}

		elements = filterBetveen(elements, ca.From, ca.To)

		elements = filterTasks(elements, "notdone")

		elements = sortTasks(elements, "date")

		out := NewOutBuilder(elements)
		if ca.Verbose {
			out = out.ShowAllTags()
		}

		for _, l := range out.Build() {
			fmt.Fprintln(w, strings.Join(l, "\t"))
		}

		err = w.Flush()
		if err != nil {
			return err
		}

		fmt.Println()
	}

	return nil
}
