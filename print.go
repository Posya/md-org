package main

import (
	"fmt"
	"strings"
)

// OutBuilder is struct to build md-org output
type OutBuilder struct {
	_elements    []element
	_showAllTags bool
}

// NewOutBuilder returns new OutBuilder struct
func NewOutBuilder(elements []element) OutBuilder {
	res := OutBuilder{
		_elements:    elements,
		_showAllTags: false,
	}
	return res
}

// ShowAllTags comand Build to add all tags (local and inherited)
func (ob *OutBuilder) ShowAllTags() OutBuilder {
	ob._showAllTags = true
	return *ob
}

// Build builds result slice
func (ob OutBuilder) Build() []string {
	res := []string{}
	indent := 0
	currentParrent := 0

	for _, el := range ob._elements {
		line := ""
		if el.

		switch v := el.(type) {
		case header:
			line = fmt.Sprintf("%s%s", strings.Repeat("\t", indent), v.text)
		case task:
			isDone := " "
			if v.done {
				isDone = "X"
			}
			line = fmt.Sprintf("%s[%s] %s", strings.Repeat("\t", indent), isDone, v.text)
		default:
			panic("Something goes wrong: element has to be task or header only")
		}

		res = append(res, line)
	}
	return res
}
