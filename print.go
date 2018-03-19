package main

import (
	"strconv"
	"strings"
)

const indentString = "  "

// OutBuilder is struct to build md-org output
type OutBuilder struct {
	_elements []element
	_verbose  bool
	_indent   bool
}

// NewOutBuilder returns new OutBuilder struct
func NewOutBuilder(elements []element) OutBuilder {
	res := OutBuilder{
		_elements: elements,
		_verbose:  false,
	}
	return res
}

// ShowAllTags comand Build to add all tags (local and inherited)
func (ob *OutBuilder) ShowAllTags() OutBuilder {
	ob._verbose = true
	return *ob
}

// Indent swiches indents on
func (ob *OutBuilder) Indent() OutBuilder {
	ob._indent = true
	return *ob
}

// Build builds result slice
func (ob OutBuilder) Build() [][]string {
	res := [][]string{}
	indents := map[int]int{}

	if ob._indent {
		indents = getIndents(ob._elements)
	}

	if ob._verbose {
		res = append(res, []string{"N", "Text", "Date", "Tags"})
	} else {
		res = append(res, []string{"N", "Text"})
	}

	for _, el := range ob._elements {

		line := []string{}

		tags := ""
		if ob._verbose {
			tags = strings.Join(el.getTags(), " ")
		}

		switch v := el.(type) {
		case header:
			line = append(line, strconv.Itoa(v.n))

			indent := ""
			if ob._indent {
				indent = strings.Repeat(indentString, indents[v.n])
			}

			line = append(line, indent+strings.Repeat("#", v.level)+" "+v.text)

			if ob._verbose {
				line = append(line, "")
				line = append(line, tags)
			}
		case task:
			line = append(line, strconv.Itoa(v.n))

			indent := ""
			if ob._indent {
				indent = strings.Repeat(indentString, indents[v.n])
			}

			isDone := "- [ ] "
			if v.done {
				isDone = "- [X] "
			}

			line = append(line, indent+isDone+v.text)

			if ob._verbose {
				line = append(line, v.date)
				line = append(line, tags)
			}
		default:
			panic("Something goes wrong: element has to be task or header only")
		}

		res = append(res, line)
	}
	return res
}
