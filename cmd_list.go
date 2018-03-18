package main

import (
	"fmt"
	"os"
	"text/tabwriter"
)

type cmdList struct {
	Verbose  bool `short:"v" long:"verbose" description:"show all tags (local and inherited)"`
	NoItdent bool `short:"i" long:"noindent" description:"print list without indents"`
}

func (cl *cmdList) Execute(args []string) error {
	dir, err := DirFromOptsOrCurrent()
	if err != nil {
		return err
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

		out := NewOutBuilder(elements)
		if cl.Verbose {
			out = out.ShowAllTags()
		}
		if !cl.NoItdent {
			out = out.Indent()
		}
		for _, l := range out.Build() {
			fmt.Fprintln(w, l)
		}

		err = w.Flush()
		if err != nil {
			return err
		}

		fmt.Println()
	}

	return nil
}
