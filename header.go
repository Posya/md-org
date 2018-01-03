package main

import (
	"errors"
	"regexp"
	"strings"
)

type header struct {
	original   string
	lineNumber int
	text       string
	level      int
	tags       []string
}

var parseRegex, tagsRegex *regexp.Regexp

func init() {
	parseRegex = regexp.MustCompile(`^\s*(#+)\s+(\w.+?)(\B[#@]\w.*)?$`)
	tagsRegex = regexp.MustCompile(`^#\w+$`)
}

func parseHeader(s string) (*header, error) {
	var h header

	m := parseRegex.FindStringSubmatch(s)

	if m == nil {
		return nil, nil
	}

	h.original = s
	h.level = len(m[1])
	h.text = m[2]
	h.tags = strings.Fields(m[3])
	for _, s := range h.tags {
		if !tagsRegex.MatchString(s) {
			return nil, errors.New("parseHeader: error in tag '" + s + "'")
		}
	}

	return &h, nil
}
