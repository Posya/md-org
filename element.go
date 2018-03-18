package main

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

func getIndents(el []element, on bool) map[int]int {
	res := map[int]int{}
	for _, e := range el {
		p := e.getParent()
		if !on || p == 0 {
			res[e.getN()] = 0
		} else {
			res[e.getN()] = res[p] + 1
		}
	}
	return res
}
