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
	exp := []string{
		"N\tText\t",
		"1\t# Заголовок 1\t",
		"2\t  - [ ] Задача 1.1 !(2018-03-18)\t",
		"3\t  - [X] Задача 1.2 !(2018-03-18 12:00)\t",
		"4\t  ## Заголовок 1.1\t",
		"5\t    - [ ] Задача 1.1.1\t",
		"6\t    - [X] Задача 1.1.2\t",
		"7\t# Заголовок 2 #header_tag\t",
		"8\t  - [ ] Задача 2.1 #task1_tag\t",
		"9\t  - [X] Задача 2.2 #task2_tag\t",
	}
	assert.Equal(t, exp, ob.Indent().Build())

	ob = NewOutBuilder(elem)
	exp = []string{
		"N\tText\t",
		"1\t# Заголовок 1\t",
		"2\t- [ ] Задача 1.1 !(2018-03-18)\t",
		"3\t- [X] Задача 1.2 !(2018-03-18 12:00)\t",
		"4\t## Заголовок 1.1\t",
		"5\t- [ ] Задача 1.1.1\t",
		"6\t- [X] Задача 1.1.2\t",
		"7\t# Заголовок 2 #header_tag\t",
		"8\t- [ ] Задача 2.1 #task1_tag\t",
		"9\t- [X] Задача 2.2 #task2_tag\t",
	}
	assert.Equal(t, exp, ob.Build())

	ob = NewOutBuilder(elem)
	exp = []string{
		"N\tText\tTags",
		"1\t# Заголовок 1\t",
		"2\t- [ ] Задача 1.1 !(2018-03-18)\t",
		"3\t- [X] Задача 1.2 !(2018-03-18 12:00)\t",
		"4\t## Заголовок 1.1\t",
		"5\t- [ ] Задача 1.1.1\t",
		"6\t- [X] Задача 1.1.2\t",
		"7\t# Заголовок 2 #header_tag\t#header_tag",
		"8\t- [ ] Задача 2.1 #task1_tag\t#task1_tag #header_tag",
		"9\t- [X] Задача 2.2 #task2_tag\t#task2_tag #header_tag",
	}
	assert.Equal(t, exp, ob.ShowAllTags().Build())
}
