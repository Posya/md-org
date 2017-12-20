package main

import (
	"fmt"
)

type Action int

const (
	list Action = iota
	check
)

func main() {
	fmt.Println("Hello, World!")
}

func do(act Action) error {
	switch act {
	case list:

	}
}
