package main

type element interface {
	IsParent(level int, isTask bool) bool
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
