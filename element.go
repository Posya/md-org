package main

import (
	"sort"
)

type element interface {
	IsParent(level int, isTask bool) bool
	getParent() int
	getTags() []string
	getN() int
	HasTag(tag string) bool
	Between(from, to string) bool
}

func filterElements(el []element, f func(element) bool) []element {
	res := []element{}
	for _, e := range el {
		if f(e) {
			res = append(res, e)
		}
	}
	return res
}

func filterTasks(el []element, what string) []element {
	if what == "all" {
		return el
	}
	res := []element{}
	for _, e := range el {
		if v, ok := e.(task); ok {
			switch what {
			case "task":
				res = append(res, v)
			case "done":
				if v.done {
					res = append(res, v)
				}
			case "notdone":
				if !v.done {
					res = append(res, v)
				}
			default:
				panic("Can't filter by " + what)
			}
		}
	}
	return res
}

func sortTasks(el []element, what string) []element {
	if what == "none" {
		return el
	}

	var sortFunc func(a, b int) bool

	switch what {
	case "date":
		sortFunc = func(a, b int) bool {
			x, ok1 := el[a].(task)
			y, ok2 := el[b].(task)
			if ok1 && ok2 {
				return x.date < y.date
			}
			return false
		}
	case "done":
		sortFunc = func(a, b int) bool {
			x, ok1 := el[a].(task)
			y, ok2 := el[b].(task)
			if ok1 && ok2 {
				return x.done && !y.done
			}
			return false
		}
	default:
		panic("Can't sort by " + what)
	}

	sort.Slice(el, sortFunc)

	return el
}

func filterBetveen(el []element, from, to string) []element {
	res := []element{}
	for _, e := range el {
		if v, ok := e.(task); ok {
			if v.Between(from, to) {
				res = append(res, v)
			}
		}
	}
	return res
}

func getIndents(el []element) map[int]int {
	res := map[int]int{}
	for _, e := range el {
		p := e.getParent()
		if p == 0 {
			res[e.getN()] = 0
		} else {
			res[e.getN()] = res[p] + 1
		}
	}
	return res
}
