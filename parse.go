package main

import (
	"errors"
	"io"
	"regexp"
)

var (
	taskOrHeaderRegexp *regexp.Regexp
	headerRegexp       *regexp.Regexp
	taskRegexp         *regexp.Regexp
	tagsRegexp         *regexp.Regexp
	dateRegexp         *regexp.Regexp
)

func init() {
	taskOrHeaderRegexp = regexp.MustCompile(`^\s*(#+|-\s+\[[ xX]\])`)
	headerRegexp = regexp.MustCompile(`^\s*(#+)\s+([\p{L}\d_].*)$`)
	taskRegexp = regexp.MustCompile(`^(\s*)-\s+\[([ xXхХvV])\]\s+([\p{L}\d_].*)$`)
	// TODO: Check for # in the middle
	tagsRegexp = regexp.MustCompile(`#[\p{L}\d_]+`)
	// TODO: Check for a dete in the middle
	dateRegexp = regexp.MustCompile(`!\((\d{4}-\d{2}-\d{2})(?: (\d{2}:\d{2}))?\)`)
}

func parseHeader(con context, s string) (context, error) {
	m := headerRegexp.FindStringSubmatch(s)
	if len(m) < 3 {
		return context{}, errors.New("Can't parse header (len(m)<3): " + s)
	}
	if len(m[1]) < 1 || len(m[2]) < 1 {
		return context{}, errors.New("Can't parse header (len(m[1])<1 || len(m[2])<1): " + s)
	}
	headerLevel := len(m[1])
	headerText := m[2]

	headerTags := tagsRegexp.FindAllString(headerText, -1)

	for i := range con.headers {
		if con.headers[i].level >= headerLevel {
			con.headers = con.headers[:i+1]
			con.headers[i].level = headerLevel
			con.headers[i].tags = headerTags
			return con, nil
		}
	}
	con.headers = append(con.headers, header{headerLevel, headerTags})
	return con, nil
}

func parseTask(con context, s string) (task, context, error) {
	m := taskRegexp.FindStringSubmatch(s)
	if len(m) < 4 {
		return task{}, context{}, errors.New("Can't parse task (len(m)<4): " + s)
	}
	if len(m[2]) < 1 || len(m[3]) < 1 {
		return task{}, context{}, errors.New("Can't parse task (len(m[2])<1||len(m[3])<1): " + s)
	}

	// taskIndent := len(m[1])
	// taskDone := m[2] != " "
	// taskText := m[3]

	// taskTags := tagsRegexp.FindAllString(taskText, -1)
	// taskDate := dateRegexp.FindString(taskText)

	panic("")

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
			// task, con, err := parseTask(con, s)
			// tasks = append(tasks, task)
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
