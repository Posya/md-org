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

func TestFilterByTag(t *testing.T) {

	ins := []element{
		header{0, 0, 0, "Header 1", []string{"#Tag1", "#Tag2"}},
		task{0, 0, 0, false, "Task1", []string{"#Tag1", "#Tag2"}, ""},
		header{0, 0, 0, "Header 1", []string{"#Tag1", "#Tag2"}},
		task{0, 0, 0, false, "Task1", []string{"#Tag1", "#Tag2"}, ""},
		header{0, 0, 0, "Header 1", []string{"#Tag1", "#Tag2"}},
		task{0, 0, 0, false, "Task1", []string{"#Tag1", "#Tag2"}, ""},
		header{0, 0, 0, "Header 1", []string{}},
		task{0, 0, 0, false, "Task1", []string{}, ""},
	}

	tags := []string{
		"#Tag1",
		"#Tag1",
		"",
		"",
		"#Tag3",
		"#Tag3",
		"#Tag1",
		"#Tag1",
	}

	exp := []bool{
		true,
		true,
		false,
		false,
		false,
		false,
		false,
		false,
	}

	if len(ins) != len(exp) || len(tags) != len(exp) {
		t.Fatal("Error in unit test: ins and exp has different length!")
	}

	for i := range exp {
		assert.Equal(t, exp[i], ins[i].FilterByTag(tags[i]))
	}

}

func TestFilterByDate(t *testing.T) {

	ins := []element{
		header{0, 0, 0, "Header 1", []string{"#Tag1", "#Tag2"}},
		task{0, 0, 0, false, "Task1", []string{"#Tag1", "#Tag2"}, ""},
	}

	from := []string{}

	to := []string{}

	exp := []bool{}

	if len(ins) != len(exp) || len(from) != len(exp) || len(to) != len(exp) {
		t.Fatal("Error in unit test: ins and exp has different length!")
	}

	for i := range exp {
		assert.Equal(t, exp[i], ins[i].FilterByDate(from[i], to[i]))
	}

}
