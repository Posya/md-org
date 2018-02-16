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

var (
	taskOrHeaderRegexp *regexp.Regexp
	headerRegexp       *regexp.Regexp
	tagsRegexp         *regexp.Regexp
)

func init() {
	taskOrHeaderRegexp = regexp.MustCompile(`^\s*(#+|-\s+\[[ xX]\])`)
	headerRegexp = regexp.MustCompile(`^\s*(#+)\s+(\w.*)$`)
	tagsRegexp = regexp.MustCompile(`#[\p{L}\d_]+`)
}

func parseHeader(con context, s string) (context, error) {
	m := headerRegexp.FindStringSubmatch(s)
	if len(m[1]) < 1 || len(m[2]) < 1 {
		return context{}, errors.New("Can't parse header: " + s)
	}
	headerLevel := len(m[1])
	headerText := m[2]

	headerTags := tagsRegexp.FindAllString(headerText, -1)

	for i := range con.headers {
		if con.headers[i].level > headerLevel {
			con.headers = con.headers[:i]
			con.headers[i].level = headerLevel
			con.headers[i].tags = headerTags
			return con, nil
		}
	}
	con.headers = append(con.headers, header{headerLevel, headerTags})
	return con, nil
}

func parse(getNext func() (string, error)) ([]task, error) {
	var tasks []task
	var err error
	var con = newContext()

	for s, err := getNext(); err != nil; s, err = getNext() {
		m := taskOrHeaderRegexp.FindString(s)
		switch {
		case len(m) == 0:
			continue
		case m[0] == '-':
			//TODO: parseTask(s)
		case m[0] == '#':
			con, err = parseHeader(con, s)
			if err != nil {
				return nil, err
			}
		default:
			return nil, errors.New("Can't parse line: " + s)
		}
	}

	if err != io.EOF {
		return nil, err
	}

	return tasks, nil
}
