package main

import (
	"errors"
	"io"
	"regexp"
	"time"
)

type task struct {
	n        int
	original string
	text     string
	tags     []string
	date     time.Time
}

var taskOrHeaderRegexp *regexp.Regexp

func init() {
	taskOrHeaderRegexp = regexp.MustCompile(`^\s*(#+|-\s+\[[ xX]\])`)
}

func getParseHeader() func(s string) tags []string {

}

func getParseTask() func(s string) tasks []task {

}

func parse(getNext func() (string, error)) ([]task, error) {
	var tasks []task
	var err error

	parseHeader := getParseHeader() 
	parseTask := getParseTask() 

	for s, err := getNext(); err != nil; s, err = getNext() {
		m := taskOrHeaderRegexp.FindString(s)
		switch {
		case len(m) == 0:
			continue
		case m[0] == '-':
			parseTask(s)
		case m[0] == '#':
			parseHeader(s)
		default:
			return nil, errors.New("Can't parse line: " + s)
		}
	}

	if err != io.EOF {
		return nil, err
	}

	return tasks, nil
}
