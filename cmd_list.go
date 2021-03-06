package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"text/tabwriter"
)

type cmdList struct {
	Verbose  bool   `short:"v" long:"verbose" env:"MDORG_VERBOSE" description:"show all tags (local and inherited)"`
	NoIndent bool   `short:"i" long:"noindent" env:"MDORG_NOINDENT" description:"print list without indents"`
	SortBy   string `short:"S" long:"sortby" choice:"none" choice:"date" choice:"done" env:"MDORG_SORT" default:"none" description:"sort by"`
	Filter   string `short:"F" long:"filter" choice:"all" choice:"task" choice:"done" choice:"notdone" env:"MDORG_FILTER" default:"all" description:"filter by"`
}

func (cl *cmdList) Execute(args []string) error {
	dir, err := DirFromOptsOrCurrent()
	if err != nil {
		return err
	}
	sort.Strings(dir)

	if cl.Filter == "all" && cl.SortBy != "none" {
		cl.Filter = "task"
	}

	for _, file := range dir {
		lines, err := ReadFile(file)
		if err != nil {
			return err
		}

		elements, err := parse(lines)
		if err != nil {
			return err
		}

		switch cl.Filter {
		case "all":
		case "task":
			elements = filterTasks(elements, "task")
		case "done":
			elements = filterTasks(elements, "done")
		case "notdone":
			elements = filterTasks(elements, "notdone")
		default:
			panic("Something goes wrong. cl.Filter = " + cl.Filter)
		}

		switch cl.SortBy {
		case "none":
		case "date":
			elements = sortTasks(elements, "date")
		case "done":
			elements = sortTasks(elements, "done")
		default:
			panic("Something goes wrong. cl.SortBy = " + cl.SortBy)
		}

		if len(elements) == 0 {
			continue
		}

		fmt.Println("File: ", file)
		fmt.Println()

		out := NewOutBuilder(elements)
		if cl.Verbose {
			out = out.ShowAllTags()
		}
		if !cl.NoIndent {
			out = out.Indent()
		}

		w := tabwriter.NewWriter(os.Stdout, 0, 8, 1, '\t', 0)

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
