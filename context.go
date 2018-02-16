package main

type context struct {
	headers []header
}

type header struct {
	level int
	tags  []string
}

type task struct {
	n        int
	original string
	text     string
	tags     []string
	date     string
}

func (orig header) Equal(other header) bool {
	if orig.level != other.level {
		return false
	}

	if len(orig.tags) != len(other.tags) {
		return false
	}

	for i := range orig.tags {
		if orig.tags[i] != other.tags[i] {
			return false
		}
	}

	return true
}

func newContext() context {
	var con context

	return con
}

func (orig context) Equal(other context) bool {
	if len(orig.headers) != len(other.headers) {
		return false
	}

	for i := range orig.headers {
		if !orig.headers[i].Equal(other.headers[i]) {
			return false
		}
	}

	return true
}
