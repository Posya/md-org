package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type cmdDone struct {
	N      int  `short:"n" required:"true" description:"number of line to mark done"`
	Undone bool `short:"u" long:"undone" description:"mark task as not done"`
}

func (cd *cmdDone) Execute(args []string) error {
	if len(opts.File) != 1 {
		return errors.New("Please specify one file to edit")
	}

	cd.N-- // First line has index 0

	lines, err := ReadFile(opts.File[0])
	if err != nil {
		return err
	}

	if cd.N >= len(lines) || cd.N < 0 {
		return errors.New("wrong number of line " + strconv.Itoa(cd.N))
	}
	line := lines[cd.N]

	t, err := parseTask(line)
	if err != nil {
		return errors.New("line " + strconv.Itoa(cd.N) + " is not task")
	}

	if !t.done && !cd.Undone {
		line = strings.Replace(line, "[ ]", "[X]", 1)
	}
	if t.done && cd.Undone {
		line = strings.Replace(line, "[X]", "[ ]", 1)
	}

	fmt.Println(line)

	lines[cd.N] = line

	err = WriteFile(lines, opts.File[0])
	if err != nil {
		return err
	}

	return nil
}
