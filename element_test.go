package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsParent(t *testing.T) {
	ins := []element{
		header{12, 1, 0, "", []string{}},
		header{12, 1, 0, "", []string{}},
		header{12, 1, 0, "", []string{}},
		header{12, 1, 0, "", []string{}},
		task{12, 1, 0, false, "asdf", []string{}, ""},
		task{12, 1, 0, false, "asdf", []string{}, ""},
		task{12, 1, 0, false, "asdf", []string{}, ""},
		task{12, 1, 0, false, "asdf", []string{}, ""},
	}

	level := []int{
		2,
		1,
		2,
		1,
		2,
		1,
		2,
		1,
	}

	isTask := []bool{
		false,
		false,
		true,
		true,
		false,
		false,
		true,
		true,
	}

	exp := []bool{
		true,
		false,
		true,
		true,
		false,
		false,
		true,
		false,
	}

	if len(ins) != len(exp) || len(level) != len(exp) || len(isTask) != len(exp) {
		t.Fatal("Error in unit test: ins and exp has different length!")
	}

	for i := range exp {
		assert.Equal(t, exp[i], ins[i].IsParent(level[i], isTask[i]))
	}

}
