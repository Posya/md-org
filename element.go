package main

type element interface {
	IsParrent(level int, isTask bool) bool
	getTags() []string
	getN() int
}

type header struct {
	n     int
	level int
	tags  []string
}

func (h header) Equal(other header) bool {
	if h.n != other.n {
		return false
	}

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

func (h header) IsParrent(level int, isTask bool) bool {
	if isTask == true {
		return true
	}

	if level > h.level {
		return true
	}

	return false
}

func (h header) getTags() []string {
	return h.tags
}

func (h header) getN() int {
	return h.n
}

type task struct {
	n        int
	level    int
	original string
	text     string
	tags     []string
	date     string
}

func (t task) Equal(other task) bool {
	if t.n != other.n {
		return false
	}

	if t.level != other.level {
		return false
	}

	if t.original != other.original {
		return false
	}

	if t.text != other.text {
		return false
	}

	if len(t.tags) != len(other.tags) {
		return false
	}

	for i := range t.tags {
		if t.tags[i] != other.tags[i] {
			return false
		}
	}

	if t.date != other.date {
		return false
	}

	return true
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

func (t task) getTags() []string {
	return t.tags
}

func (t task) getN() int {
	return t.n
}

// kfthxvveozuriwtc
