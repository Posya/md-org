package main

type header struct {
	n      int
	level  int
	parent int
	text   string
	tags   []string
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

func (h header) getParent() int {
	return h.parent
}

func (h header) getTags() []string {
	return h.tags
}

func (h header) getN() int {
	return h.n
}

func (h header) HasTag(tag string) bool {
	if tag == "" {
		return true
	}

	if tag[0] != '#' {
		tag = "#" + tag
	}

	for _, ct := range h.tags {
		if ct == tag {
			return true
		}
	}
	return false
}

func (h header) Between(from, to string) bool {
	return false
}
