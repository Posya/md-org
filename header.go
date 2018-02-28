package main

type header struct {
	n      int
	level  int
	parent int
	text   string
	tags   []string
}

func (h header) Equal(other header) bool {
	if h.n != other.n {
		return false
	}

	if h.level != other.level {
		return false
	}

	if h.parent != other.parent {
		return false
	}

	if h.text != other.text {
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

func (h header) IsParent(level int, isTask bool) bool {
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
