package main

type context struct {
	headers []header
}

type header struct {
	level int
	tags  []string
}

func newContext() context {
	var con context

	return con
}
