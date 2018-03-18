package main

import "fmt"

type cmdList struct {
	ShowAllTags bool `short:"t" long:"tags" description:"show all tags (local and inherited)"`
	NoItdent    bool `short:"i" long:"noindent" description:"print list without indents"`
}

func (cl *cmdList) Execute(args []string) error {
	dir, err := DirFromOptsOrCurrent()
	if err != nil {
		return err
	}

	for _, file := range dir {
		fmt.Println("File: ", file)
		fmt.Println()

		lines, err := ReadFile(file)
		if err != nil {
			return err
		}

		elements, err := parse(lines)
		if err != nil {
			return err
		}

		out := NewOutBuilder(elements)
		if cl.ShowAllTags {
			out = out.ShowAllTags()
		}
		if !cl.NoItdent {
			out = out.Indent()
		}
		for _, l := range out.Build() {
			fmt.Println(l)
		}

		fmt.Println()
	}

	return nil
}
