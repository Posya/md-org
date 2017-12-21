package main

import (
	"fmt"
	"logit"

	"action"
)

func main() {
	fmt.Println("Hello, World!")
}

func do(act action.Action) error {
	switch act {
	case action.List:
		logit.Trace("This is list action")
	case action.Check:
		logit.Trace("This in check action")

	}

	return nil
}
