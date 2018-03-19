package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuild(t *testing.T) {
	elem := []element{
		header{n: 1, level: 1, parent: 0, text: "Заголовок 1", tags: []string{}},
		task{n: 2, level: 3, parent: 1, done: false, text: "Задача 1.1 !(2018-03-18)", tags: []string{}, date: "2018-03-18"},
		task{n: 3, level: 3, parent: 1, done: true, text: "Задача 1.2 !(2018-03-18 12:00)", tags: []string{}, date: "2018-03-18 12:00"},
		header{n: 4, level: 2, parent: 1, text: "Заголовок 1.1", tags: []string{}},
		task{n: 5, level: 6, parent: 4, done: false, text: "Задача 1.1.1", tags: []string{}, date: ""},
		task{n: 6, level: 6, parent: 4, done: true, text: "Задача 1.1.2", tags: []string{}, date: ""},
		header{n: 7, level: 1, parent: 0, text: "Заголовок 2 #header_tag", tags: []string{"#header_tag"}},
		task{n: 8, level: 3, parent: 7, done: false, text: "Задача 2.1 #task1_tag", tags: []string{"#task1_tag", "#header_tag"}, date: ""},
		task{n: 9, level: 3, parent: 7, done: true, text: "Задача 2.2 #task2_tag", tags: []string{"#task2_tag", "#header_tag"}, date: ""},
	}

	ob := NewOutBuilder(elem)
	exp := [][]string{
		[]string{"N", "Text"},
		[]string{"1", "# Заголовок 1"},
		[]string{"2", "  - [ ] Задача 1.1 !(2018-03-18)"},
		[]string{"3", "  - [X] Задача 1.2 !(2018-03-18 12:00)"},
		[]string{"4", "  ## Заголовок 1.1"},
		[]string{"5", "    - [ ] Задача 1.1.1"},
		[]string{"6", "    - [X] Задача 1.1.2"},
		[]string{"7", "# Заголовок 2 #header_tag"},
		[]string{"8", "  - [ ] Задача 2.1 #task1_tag"},
		[]string{"9", "  - [X] Задача 2.2 #task2_tag"},
	}
	assert.Equal(t, exp, ob.Indent().Build())

	ob = NewOutBuilder(elem)
	exp = [][]string{
		[]string{"N", "Text"},
		[]string{"1", "# Заголовок 1"},
		[]string{"2", "- [ ] Задача 1.1 !(2018-03-18)"},
		[]string{"3", "- [X] Задача 1.2 !(2018-03-18 12:00)"},
		[]string{"4", "## Заголовок 1.1"},
		[]string{"5", "- [ ] Задача 1.1.1"},
		[]string{"6", "- [X] Задача 1.1.2"},
		[]string{"7", "# Заголовок 2 #header_tag"},
		[]string{"8", "- [ ] Задача 2.1 #task1_tag"},
		[]string{"9", "- [X] Задача 2.2 #task2_tag"},
	}
	assert.Equal(t, exp, ob.Build())

	ob = NewOutBuilder(elem)
	exp = [][]string{
		[]string{"N", "Text", "Date", "Tags"},
		[]string{"1", "# Заголовок 1", "", ""},
		[]string{"2", "- [ ] Задача 1.1 !(2018-03-18)", "2018-03-18", ""},
		[]string{"3", "- [X] Задача 1.2 !(2018-03-18 12:00)", "2018-03-18 12:00", ""},
		[]string{"4", "## Заголовок 1.1", "", ""},
		[]string{"5", "- [ ] Задача 1.1.1", "", ""},
		[]string{"6", "- [X] Задача 1.1.2", "", ""},
		[]string{"7", "# Заголовок 2 #header_tag", "", "#header_tag"},
		[]string{"8", "- [ ] Задача 2.1 #task1_tag", "", "#task1_tag #header_tag"},
		[]string{"9", "- [X] Задача 2.2 #task2_tag", "", "#task2_tag #header_tag"},
	}
	assert.Equal(t, exp, ob.ShowAllTags().Build())
}
