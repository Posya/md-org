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

	headerFuncs := []printFunc{
		func(el element) (s string, skip bool) {
			if v, ok := el.(header); ok {
				return strconv.Itoa(v.n), false
			}
			panic("Wrong type: had to be header")
		},
		func(el element) (s string, skip bool) {
			if v, ok := el.(header); ok {
				indent := ""
				if ob._indent {
					indent = strings.Repeat(indentString, indents[v.n])
				}
				return indent + strings.Repeat("#", v.level) + " " + v.text, false
			}
			panic("Wrong type: had to be header")
		},
		func(el element) (s string, skip bool) {
			if _, ok := el.(header); ok {
				return "", !ob._verbose
			}
			panic("Wrong type: had to be header")
		},
		func(el element) (s string, skip bool) {
			if v, ok := el.(header); ok {
				if ob._verbose {
					return strings.Join(v.tags, " "), false
				}
				return "", true
			}
			panic("Wrong type: had to be header")
		},
	}
	taskFuncs := []printFunc{
		func(el element) (s string, skip bool) {
			if v, ok := el.(task); ok {
				return strconv.Itoa(v.n), false
			}
			panic("Wrong type: had to be task")
		},
		func(el element) (s string, skip bool) {
			if v, ok := el.(task); ok {
				indent := ""
				if ob._indent {
					indent = strings.Repeat(indentString, indents[v.n])
				}

				isDone := "- [ ] "
				if v.done {
					isDone = "- [X] "
				}

				return indent + isDone + v.text, false
			}
			panic("Wrong type: had to be task")
		},
		func(el element) (s string, skip bool) {
			if v, ok := el.(task); ok {
				return v.date, !ob._verbose
			}
			panic("Wrong type: had to be task")
		},
		func(el element) (s string, skip bool) {
			if v, ok := el.(task); ok {
				if ob._verbose {
					return strings.Join(v.tags, " "), false
				}
				return "", true
			}
			panic("Wrong type: had to be task")
		},
	}

	for _, el := range ob._elements {
		res = append(res, prepareToPrint(el, headerFuncs, taskFuncs))
	}
	return res
}

type printFunc func(element) (s string, skip bool)

func prepareToPrint(el element, headerFuncs, taskFuncs []printFunc) []string {
	res := []string{}
	var funcsToRun []printFunc

	switch el.(type) {
	case header:
		funcsToRun = headerFuncs
	case task:
		funcsToRun = taskFuncs
	default:
		panic("Something goes wring: Found new type of element")
	}
	for _, pf := range funcsToRun {
		s, skip := pf(el)
		if !skip {
			res = append(res, s)
		}
	}

	return res
}
