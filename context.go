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
	level    int
	original string
	text     string
	tags     []string
	date     string
}

func (h header) Equal(other header) bool {
	if h.level != other.level {
		return false
	}

	if len(h.tags) != len(other.tags) {
		return false
	}

	for i := range h.tags {
		if h.tags[i] != other.tags[i] {
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

type element interface {
	IsParrent(level int, isTask bool) bool
	getTags() []string
	getN() int
}

func (t task) IsParrent(level int, isTask bool) bool {
	if isTask == false {
		return false
	}

	if level > t.level {
		return true
	}

	return false
}

func (h header) IsParrent(level int, isTask bool) bool {
	if isTask == true {
		return true
	}

	if level > h.level {
		return true
	}

	return false
}
