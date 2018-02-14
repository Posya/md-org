package task

import (
	"date"
	"errors"
	"regexp"
	"strings"
)

// Task is struct for task data keeping
type Task struct {
	Original   string
	LineNumber int
	Level      int
	Text       string
	Done       bool
	Tags       []string
	Date       date.Date
}

var parseTaskRegex, tagsAndDatesRegex, dateRegex *regexp.Regexp

func init() {
	parseTaskRegex = regexp.MustCompile(`^(\s*)- \[([\w ]?)\]\s+(\w.+?)(\B[#@]\w.*)?$`)
	tagsAndDatesRegex = regexp.MustCompile(`^([@#])[\w\.:+-]+$`)
	dateRegex = regexp.MustCompile(`^@(\d{1,2}\.\d{2}(?:\.\d{4})?)(?:_(\d{1,2}:\d{2}))?(?:-(\d{1,2}\.\d{2}(?:\.\d{4})?)(?:_(\d{1,2}:\d{2}))?)?(?:\+(\d+)([hdwmy]))?$`)

}

func parseTask(s string) (*Task, error) {
	var t Task

	m := parseTaskRegex.FindStringSubmatch(s)

	if m == nil {
		return nil, nil
	}

	t.Original = s
	t.Level = len(m[1])
	switch m[2] {
	case " ":
		t.Done = false
	case "X", "x":
		t.Done = true
	default:
		return nil, errors.New("parseTask: wrong done state '" + m[2] + "'")
	}

	t.Text = m[3]
	tagsAndDates := strings.Fields(m[4])
	var dates []string
	for _, s := range tagsAndDates {
		if m := tagsAndDatesRegex.FindStringSubmatch(s); m != nil {
			switch m[1] {
			case "#":
				t.Tags = append(t.Tags, s)
			case "@":
				dates = append(dates, s)
			default:
				return nil, errors.New("parseTask: some internal error with string '" + s + "'")
			}
		} else {
			return nil, errors.New("parseTask: error in tags and dates '" + s + "'")
		}

	}

	d, err := parseDates(dates)
	if err != nil {
		return nil, err
	}
	t.Date = *d

	return &t, nil
}

func parseDates(dates []string) (*date.Date, error) {
	var result [7]string

	for _, d := range dates {

		m := dateRegex.FindStringSubmatch(d)

		if m == nil {
			return nil, errors.New("parseDates: this is not date: " + d)
		}

		for i := 1; i < 7; i++ {
			if m[i] == "" {
				continue
			}

			if result[i] != "" {
				return nil, errors.New("parseDates: error: " + m[i] + " in date: " + d)
			}

			result[i] = m[i]
		}
	}

	//result[0] = date.combineResults(result)

	return nil, nil
}
