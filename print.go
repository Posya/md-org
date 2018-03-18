package main

import (
	"fmt"
	"strings"
)

// OutBuilder is struct to build md-org output
type OutBuilder struct {
	_elements    []element
	_showAllTags bool
	_Indent      bool
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

// Indent swiches indents on
func (ob *OutBuilder) Indent() OutBuilder {
	ob._Indent = true
	return *ob
}

// Build builds result slice
func (ob OutBuilder) Build() []string {
	res := []string{}
	indents := getIndents(ob._elements, ob._Indent)

	for _, el := range ob._elements {
		line := ""
		tags := ""
		if ob._showAllTags && len(el.getTags()) > 0 {
			tags = "\t" + strings.Join(el.getTags(), " ")
		}

		switch v := el.(type) {
		case header:
			line = fmt.Sprintf("%s%s %s%s", strings.Repeat("\t", indents[v.n]), strings.Repeat("#", v.level), v.text, tags)
		case task:
			isDone := " "
			if v.done {
				isDone = "X"
			}
			line = fmt.Sprintf("%s- [%s] %s%s", strings.Repeat("\t", indents[v.n]), isDone, v.text, tags)
		default:
			panic("Something goes wrong: element has to be task or header only")
		}

		res = append(res, line)
	}
	return res
}
