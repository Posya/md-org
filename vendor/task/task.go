package task

import (
	"errors"
	"regexp"
	"strings"
)

type task struct {
	original   string
	lineNumber int
	level      int
	text       string
	done       bool
	tags       []string
	date       string
}

var parseTaskRegex, tagsAndDatesRegex, dateRegex *regexp.Regexp

func init() {
	parseTaskRegex = regexp.MustCompile(`^(\s*)- \[([\w ]?)\]\s+(\w.+?)(\B[#@]\w.*)?$`)
	tagsAndDatesRegex = regexp.MustCompile(`^([@#])[\w\.:+-]+$`)
	dateRegex = regexp.MustCompile(`^@(\d{1,2}\.\d{2}(?:\.\d{4})?)(?:_(\d{1,2}:\d{2}))?(?:-(\d{1,2}\.\d{2}(?:\.\d{4})?)(?:_(\d{1,2}:\d{2}))?)?(?:\+(\d+)([hdwmy]))?$`)

}

func parseTask(s string) (*task, error) {
	var t task

	m := parseTaskRegex.FindStringSubmatch(s)

	if m == nil {
		return nil, nil
	}

	t.original = s
	t.level = len(m[1])
	switch m[2] {
	case " ":
		t.done = false
	case "X", "x":
		t.done = true
	default:
		return nil, errors.New("parseTask: wrong done state '" + m[2] + "'")
	}

	t.text = m[3]
	tagsAndDates := strings.Fields(m[4])
	var dates []string
	for _, s := range tagsAndDates {
		if m := tagsAndDatesRegex.FindStringSubmatch(s); m != nil {
			switch m[1] {
			case "#":
				t.tags = append(t.tags, s)
			case "@":
				dates = append(dates, s)
			default:
				return nil, errors.New("parseTask: some internal error with string '" + s + "'")
			}
		} else {
			return nil, errors.New("parseTask: error in tags and dates '" + s + "'")
		}

	}

	t.date = parseDates(dates)

	return &t, nil
}

func parseDates(dates []string) string {
	var result [6]string

	for _, d := range dates {

		m := dateRegex.FindStringSubmatch(d)

		if m == nil {
			return ""
		}

		for i := 1; i < 6; i++ {
			if m[i] == "" {
				continue
			}

			if result[i] != "" {
				return ""
			}

			result[i] = m[i]
		}
	}

	//result[0] = combineResults(result)

	return result[0]
}
