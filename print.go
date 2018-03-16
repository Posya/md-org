package main

// OutBuilder is struct to build md-org output
type OutBuilder struct {
	elements    []element
	showAllTags bool
}

// NewOutBuilder returns new OutBuilder struct
func NewOutBuilder(elements []element) OutBuilder {
	res := OutBuilder{
		elements:    elements,
		showAllTags: false,
	}
	return res
}
