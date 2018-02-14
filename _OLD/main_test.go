package main

import (
	"testing"

	"action"
)

func TestMain(t *testing.T) {
	//main()
}

func TestDoList(t *testing.T) {
	err := do(action.List)
	if err != nil {
		t.Error("Error occurs: ", err)
	}
}

func TestDoCheck(t *testing.T) {
	err := do(action.Check)
	if err != nil {
		t.Error("Error occurs: ", err)
	}
}

func TestParse(t *testing.T) {
	parse("test.md-org")
}
