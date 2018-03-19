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

	test := map[string]element{
		"header without tags": header{0, 0, 0, "Header 1", []string{}},
		"header with tags":    header{0, 0, 0, "Header 1", []string{"#Tag1", "#Tag2"}},

		"task without tags": task{0, 0, 0, false, "Task1", []string{}, ""},
		"task with tags":    task{0, 0, 0, false, "Task1", []string{"#Tag1", "#Tag2"}, ""},
	}

	assert.Equal(t, false, test["header without tags"].HasTag("#TestTag"))
	assert.Equal(t, false, test["header with tags"].HasTag("#WrongTag"))
	assert.Equal(t, true, test["header with tags"].HasTag("#Tag1"))
	assert.Equal(t, true, test["header without tags"].HasTag(""))
	assert.Equal(t, true, test["header with tags"].HasTag(""))
	assert.Equal(t, true, test["header with tags"].HasTag("Tag1"))

	assert.Equal(t, false, test["task without tags"].HasTag("#TestTag"))
	assert.Equal(t, false, test["task with tags"].HasTag("#WrongTag"))
	assert.Equal(t, true, test["task with tags"].HasTag("#Tag1"))
	assert.Equal(t, true, test["task without tags"].HasTag(""))
	assert.Equal(t, true, test["task with tags"].HasTag(""))
	assert.Equal(t, true, test["task with tags"].HasTag("Tag1"))
}

func TestFilterTasks(t *testing.T) {
	test := []element{
		header{n: 1},
		task{n: 2, done: false},
		task{n: 3, done: true},
		header{n: 4},
		task{n: 5, done: true},
		task{n: 6, done: true},
	}

	assert.Equal(t, 6, len(filterTasks(test, "all")))
	assert.Equal(t, 4, len(filterTasks(test, "task")))
	assert.Equal(t, 3, len(filterTasks(test, "done")))
	assert.Equal(t, 1, len(filterTasks(test, "notdone")))
}

func TestSortTasks(t *testing.T) {
	test := []element{
		task{n: 2, done: false, date: "2018-03-18"},
		task{n: 3, done: true, date: "2018-03-17"},
		task{n: 5, done: true, date: "2018-03-16 12:00"},
		task{n: 6, done: true, date: "2018-03-16"},
	}

	assert.Equal(t, test, sortTasks(test, "none"))

	res := sortTasks(test, "done")
	assert.Equal(t, 2, res[3].getN())

	res = sortTasks(test, "date")
	assert.Equal(t, 6, res[0].getN())
	assert.Equal(t, 5, res[1].getN())
	assert.Equal(t, 3, res[2].getN())
	assert.Equal(t, 2, res[3].getN())
}
