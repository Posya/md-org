package main

import (
	"errors"
	"io"
	"regexp"
	"time"
)

var (
	taskOrHeaderRegexp *regexp.Regexp
	headerRegexp       *regexp.Regexp
	taskRegexp         *regexp.Regexp
	tagsRegexp         *regexp.Regexp
	dateRegexp         *regexp.Regexp
)

var location = time.Now().Location()

func init() {
	taskOrHeaderRegexp = regexp.MustCompile(`^\s*(#+|-\s+\[[ xX]\])`)
	headerRegexp = regexp.MustCompile(`^\s*(#+)\s+([\p{L}\d_].*)$`)
	taskRegexp = regexp.MustCompile(`^(\s*)-\s+\[([ xXхХvV])\]\s+([\p{L}\d_].*)$`)
	// TODO: Check for # in the middle
	tagsRegexp = regexp.MustCompile(`#[\p{L}\d_]+`)
	// TODO: Check for a dete in the middle
	dateRegexp = regexp.MustCompile(`!\((\d{4}-\d{2}-\d{2})(?: (\d{2}:\d{2}))?\)`)
}

func parseHeader(s string) (header, error) {
	m := headerRegexp.FindStringSubmatch(s)
	if len(m) < 3 {
		return header{}, errors.New("Can't parse header (len(m)<3): " + s)
	}
	if len(m[1]) < 1 || len(m[2]) < 1 {
		return header{}, errors.New("Can't parse header (len(m[1])<1 || len(m[2])<1): " + s)
	}
	headerLevel := len(m[1])
	headerText := m[2]

	headerTags := tagsRegexp.FindAllString(headerText, -1)

	return header{0, headerLevel, headerText, headerTags}, nil
}

func parseTask(s string) (task, error) {
	m := taskRegexp.FindStringSubmatch(s)
	if len(m) < 4 {
		return task{}, errors.New("Can't parse task (len(m)<4): " + s)
	}
	if len(m[2]) < 1 || len(m[3]) < 1 {
		return task{}, errors.New("Can't parse task (len(m[2])<1||len(m[3])<1): " + s)
	}

	taskLevel := len(m[1])
	taskDone := m[2] != " "
	taskText := m[3]

	taskTags := tagsRegexp.FindAllString(taskText, -1)
	taskDate := dateRegexp.FindString(taskText)

	if !checkDate(taskDate) {
		return task{}, errors.New("Can't parse task (wrong date): " + s)
	}

	return task{0, taskLevel, taskDone, 0, taskText, taskTags, taskDate}, nil
}

func checkDate(s string) bool {
	_, err1 := time.ParseInLocation("2006.01.02", s, location)
	_, err2 := time.ParseInLocation("2006.01.02 15:04", s, location)
	if err1 == nil || err2 == nil {
		return true
	}
	return false
}

func parse(getNext func() (string, error)) ([]element, error) {
	var elements []element
	var err error

	lineN := 0

	for s, err := getNext(); err != nil; s, err = getNext() {
		lineN++
		m := taskOrHeaderRegexp.FindString(s)
		switch {
		case len(m) == 0:
			continue
		case m[0] == '-':
			t, err := parseTask(s)
			if err != nil {
				return nil, err
			}
			for i := len(elements) - 1; i >= 0; i-- {
				if elements[i].IsParent(t.level, false) {
					t.tags = append(t.tags, elements[i].getTags()...)
				}
			}
			t.n = lineN
			elements = append(elements, t)
		case m[0] == '#':
			h, err := parseHeader(s)
			if err != nil {
				return nil, err
			}
			for i := len(elements) - 1; i >= 0; i-- {
				if elements[i].IsParent(h.level, false) {
					h.tags = append(h.tags, elements[i].getTags()...)
				}
			}
			h.n = lineN
			elements = append(elements, h)
		default:
			return nil, errors.New("Can't parse line: " + s)
		}
	}

	if err != io.EOF {
		return nil, err
	}

	return elements, nil
}
