package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHeaderEqual(t *testing.T) {
	ins := []header{
		header{12, 1, "", []string{"tag1", "tag2", "tag3"}},
	}

	exp := []header{
		header{12, 1, "", []string{"tag1", "tag2", "tag3"}},
	}

	if len(ins) != len(exp) {
		t.Fatal("Error in unit test: ins and exp has different length!")
	}

	for i := range exp {
		assert.True(t, exp[i].Equal(ins[i]))
	}
}

func TestIsParent(t *testing.T) {
	ins := []element{
		header{12, 1, "", []string{}},
		header{12, 1, "", []string{}},
		header{12, 1, "", []string{}},
		header{12, 1, "", []string{}},
		task{12, 1, false, 0, "asdf", []string{}, ""},
		task{12, 1, false, 0, "asdf", []string{}, ""},
		task{12, 1, false, 0, "asdf", []string{}, ""},
		task{12, 1, false, 0, "asdf", []string{}, ""},
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
