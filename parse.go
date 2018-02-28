package main

import (
	"errors"
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
	if headerTags == nil {
		headerTags = []string{}
	}

	return header{0, headerLevel, 0, headerText, headerTags}, nil
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
	if taskTags == nil {
		taskTags = []string{}
	}

	taskDate := dateRegexp.FindString(taskText)

	if !dateIsCorrect(taskDate) {
		return task{}, errors.New("Can't parse task (wrong date): " + s)
	}

	return task{0, taskLevel, 0, taskDone, taskText, taskTags, taskDate}, nil
}

func dateIsCorrect(s string) bool {
	if s == "" {
		return true
	}
	_, err1 := time.ParseInLocation("2006.01.02", s, location)
	_, err2 := time.ParseInLocation("2006.01.02 15:04", s, location)
	if err1 == nil || err2 == nil {
		return true
	}
	return false
}

func parse(lines []string) ([]element, error) {
	elements := []element{}

	for lineN := range lines {
		m := taskOrHeaderRegexp.FindStringSubmatch(lines[lineN])
		if m == nil {
			continue
		}

		switch {
		case m[1][0] == '-':
			t, err := parseTask(lines[lineN])
			if err != nil {
				return nil, err
			}
			for i := len(elements) - 1; i >= 0; i-- {
				if elements[i].IsParent(t.level, true) {
					t.tags = append(t.tags, elements[i].getTags()...)
					t.parent = elements[i].getN()
					break
				}
			}
			t.n = lineN + 1
			elements = append(elements, t)
		case m[1][0] == '#':
			h, err := parseHeader(lines[lineN])
			if err != nil {
				return nil, err
			}
			for i := len(elements) - 1; i >= 0; i-- {
				if elements[i].IsParent(h.level, false) {
					h.tags = append(h.tags, elements[i].getTags()...)
					h.parent = elements[i].getN()
					break
				}
			}
			h.n = lineN + 1
			elements = append(elements, h)
		default:
			return nil, errors.New("Can't parse line: " + lines[lineN])
		}
	}

	return elements, nil
}
