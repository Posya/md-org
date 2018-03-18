package main

type task struct {
	n      int
	level  int
	parent int
	done   bool
	text   string
	tags   []string
	date   string
}

func (t task) IsParent(level int, isTask bool) bool {
	if isTask == false {
		return false
	}

	if level > t.level {
		return true
	}

	return false
}

func (t task) getParent() int {
	return t.parent
}

func (t task) getTags() []string {
	return t.tags
}

func (t task) getN() int {
	return t.n
}

func (t task) HasTag(tag string) bool {
	if tag == "" {
		return true
	}

	if tag[0] != '#' {
		tag = "#" + tag
	}

	for _, ct := range t.tags {
		if ct == tag {
			return true
		}
	}
	return false
}

func (t task) Between(from, to string) bool {
	if from <= t.date && t.date <= to {
		return true
	}
	return false
}
