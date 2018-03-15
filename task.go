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

func (t task) Equal(other task) bool {
	if t.n != other.n {
		return false
	}

	if t.level != other.level {
		return false
	}

	if t.parent != other.parent {
		return false
	}

	if t.done != other.done {
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

func (t task) IsParent(level int, isTask bool) bool {
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

func (h task) FilterByTag(tag string) bool {
	panic("")
}

func (h task) FilterByDate(from, to string) bool {
	panic("")
}
