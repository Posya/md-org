package main

import "fmt"

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

func (ob *OutBuilder) ShowAllTags() OutBuilder {
	ob._showAllTags = true
	return *ob
}

func (ob OutBuilder) Build() []string {
	res := []string{}
	for _, el := range ob._elements {
		isDone := " "
		line := ""
		switch v := el.(type) {
		case task:
			if v.done {
				isDone = "X"
				line = fmt.Sprintf("[%s]", isDone)
			}
		default:
			panic("Something goes wgong: element has to be task or header only")
		}

		res = append(res, line)
	}
}
