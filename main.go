package main

import (
	"action"
	"bufio"
	"fmt"
	"os"
	"regexp"
)

var (
	reHeader = regexp.MustCompile("^\\s*#+")
	reTask   = regexp.MustCompile("^\\s*-\\s+\\[[ xX]\\]\\s+")
)

func main() {
	fmt.Println("Hello, World!")
}

func do(act action.Action) error {
	switch act {
	case action.List:
		// logit.Trace("This is list action")
	case action.Check:
		// logit.Trace("This in check action")

	}

	return nil
}

func parse(path string) error {
	inFile, err := os.Open(path)
	if err != nil {
		return err
	}
	defer inFile.Close()

	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		//s := scanner.Text()
	}

	return nil
}
