package main

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	ins := [][]string{
		[]string{
			"Первая строка",
			"# Заголовок 1",
			"- [ ] First task",
		},
		[]string{
			"asdfasdfasdf asdfasdfasdf",
			"# Заголовок 1",
			"asdfasdfasdf asdfasdfasdf",
			" # Заголовок 2",
			"asdfasdfasdf asdfasdfasdf",
			"## Заголовок 2.1",
			"## Заголовок 2.2",
			"asdfasdfasdf asdfasdfasdf",
			"- [ ] First task 1",
			"	- [ ] First task 1.1",
			"	- [ ] First task 1.2",
			"asdfasdfasdf asdfasdfasdf",
			"		- [ ] First task 1.2.1",
			" # Заголовок 3 #tag1",
			"## Заголовок 3.1 #tag1_1 ",
			"asdfasdfasdf asdfasdfasdf",
			"		- [ ] First task 3 #task_tag_1",
			"- [ ] First task 4",
			"	- [ ] First task 4.1 !(2018.03.01) сделать.",
			"asdfasdfasdf asdfasdfasdf",
			"	- [ ] First !(2018-03-20 14:45) task 4.2",
		},
	}

	exp := [][]element{
		[]element{
			header{2, 1, 0, "Заголовок 1", []string{}},
			task{3, 0, 2, false, "First task", []string{}, ""},
		},
		[]element{
			header{2, 1, 0, "Заголовок 1", []string{}},
			header{4, 1, 0, "Заголовок 2", []string{}},
			header{6, 2, 4, "Заголовок 2.1", []string{}},
			header{7, 2, 4, "Заголовок 2.2", []string{}},
			task{9, 0, 7, false, "First task 1", []string{}, ""},
			task{10, 1, 9, false, "First task 1.1", []string{}, ""},
			task{11, 1, 9, false, "First task 1.2", []string{}, ""},
			task{13, 2, 11, false, "First task 1.2.1", []string{}, ""},
			header{14, 1, 0, "Заголовок 3 #tag1", []string{"#tag1"}},
			header{15, 2, 14, "Заголовок 3.1 #tag1_1 ", []string{"#tag1_1", "#tag1"}},
			task{17, 2, 15, false, "First task 3 #task_tag_1", []string{"#task_tag_1", "#tag1_1", "#tag1"}, ""},
			task{18, 0, 15, false, "First task 4", []string{"#tag1_1", "#tag1"}, ""},
			task{19, 1, 18, false, "First task 4.1 !(2018.03.01) сделать.", []string{"#tag1_1", "#tag1"}, "2018.03.01"},
			task{21, 1, 18, false, "First !(2018-03-20 14:45) task 4.2", []string{"#tag1_1", "#tag1"}, "2018.03.20 14:45"},
		},
	}

	if len(ins) != len(exp) {
		t.Fatal("Error in unit test: ins and exp has different length!")
	}

	for i := range exp {
		v, err := parse(ins[i])
		assert.NoError(t, err)
		assert.Equal(t, exp[i], v)
	}
}

func TestParseHeader(t *testing.T) {
	ins := []string{
		"# Header 1 #tag11, #tag12, #tag13",
		"### Заголовок 1 #тег_1, #ещёТег #и_последний тег",
		" # Заголовок с пробелом и с #тегами",
	}

	exp := []header{
		header{0, 1, 0, "Header 1 #tag11, #tag12, #tag13", []string{"#tag11", "#tag12", "#tag13"}},
		header{0, 3, 0, "Заголовок 1 #тег_1, #ещёТег #и_последний тег", []string{"#тег_1", "#ещёТег", "#и_последний"}},
		header{0, 1, 0, "Заголовок с пробелом и с #тегами", []string{"#тегами"}},
	}

	if len(ins) != len(exp) {
		t.Fatal("Error in unit test: ins, exp and con has different length!")
	}

	for i := range exp {
		v, err := parseHeader(ins[i])
		assert.NoError(t, err)
		assert.Equal(t, exp[i], v)
	}
}

func TestParseHeaderError(t *testing.T) {
	_, err := parseHeader("#")
	assert.Error(t, err)
}

func TestParseTask(t *testing.T) {
	ins := []string{
		"- [ ] Task 1",
	}

	exp := []task{
		task{0, 0, 0, false, "Task 1", []string{}, ""},
	}

	if len(ins) != len(exp) {
		t.Fatal("Error in unit test: ins, exp and con has different length!")
	}

	for i := range exp {
		v, err := parseTask(ins[i])
		assert.NoError(t, err)
		assert.Equal(t, exp[i], v)
	}
}

func TestParseTaskError(t *testing.T) {
	_, err := parseTask("- [ ] Task !(2018-02-30)")
	assert.Error(t, err)
}

func TestDateIsCorrect(t *testing.T) {
	ins := []string{
		"2018.11.05",
		"2018.11.5",
		"2018.11.01 17:20",
	}

	exp := []bool{
		true,
		false,
		true,
	}

	if len(ins) != len(exp) {
		t.Fatal("Error in unit test: ins and exp has different length!")
	}

	for i := range exp {
		assert.Equal(t, exp[i], dateIsCorrect(ins[i]), "error in date (must be "+strconv.FormatBool(exp[i])+"): "+ins[i])
	}
}
