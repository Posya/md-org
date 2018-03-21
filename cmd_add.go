package main

import "strings"

type cmdAdd struct {
}

func (ca *cmdAdd) Execute(args []string) error {
	lines := []string{}
	line := "- [ ] " + strings.Join(args, " ")
	lines = append(lines, line)
	return AppendFile(lines, "./inbox.org.md")
}
