package main

import (
	"fmt"
	"sort"
)

type cmdArchive struct {
	Verbose     bool   `short:"v" long:"verbose" description:"show more"`
	ArchiveSign string `short:"s" long:"sign" default:"@" description:"sign to archive line. Has to be just one and first symbol of the line."`
	ArchiveFile string `short:"a" long:"archive" default:"./archive.org.md" description:"file name to archive"`
}

func (ca *cmdArchive) Execute(args []string) error {
	dir, err := DirFromOptsOrCurrent()
	if err != nil {
		return err
	}
	sort.Strings(dir)

	for _, file := range dir {

		lines, err := ReadFile(file)
		if err != nil {
			return err
		}

		keepLines := []string{}
		archiveLines := []string{}

		for _, line := range lines {
			if len(line) > 0 && line[0] == '@' {
				archiveLines = append(archiveLines, line)
			} else {
				keepLines = append(keepLines, line)
			}
		}

		err = AppendFile(archiveLines, "./archive.org.md")
		if err != nil {
			return err
		}
		err = WriteFile(keepLines, file)
		if err != nil {
			return err
		}

		if ca.Verbose {
			fmt.Printf("%s: %d lines was archived\n", file, len(archiveLines))
		}
	}

	return nil
}
