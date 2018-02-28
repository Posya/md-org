package main

type element interface {
	IsParent(level int, isTask bool) bool
	getTags() []string
	getN() int
}
