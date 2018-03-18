package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuild(t *testing.T) {
	elem := []element{
		header{n: 1, level: 1, parent: 0, text: "Заголовок 1", tags: []string{}},
		task{n: 2, level: 3, parent: 1, done: false, text: "Задача 1.1", tags: []string{}, date: ""},
		task{n: 3, level: 3, parent: 1, done: true, text: "Задача 1.2", tags: []string{}, date: ""},
		header{n: 4, level: 2, parent: 1, text: "Заголовок 1.1", tags: []string{}},
		task{n: 5, level: 6, parent: 4, done: false, text: "Задача 1.1.1", tags: []string{}, date: ""},
		task{n: 6, level: 6, parent: 4, done: true, text: "Задача 1.1.2", tags: []string{}, date: ""},
		header{n: 7, level: 1, parent: 0, text: "Заголовок 2 #header_tag", tags: []string{"#header_tag"}},
		task{n: 8, level: 3, parent: 7, done: false, text: "Задача 2.1 #task1_tag", tags: []string{"#task1_tag", "#header_tag"}, date: ""},
		task{n: 9, level: 3, parent: 7, done: true, text: "Задача 2.2 #task2_tag", tags: []string{"#task2_tag", "#header_tag"}, date: ""},
	}

	ob := NewOutBuilder(elem)
	exp := []string{
		"Заголовок 1",
		"\t[ ] Задача 1.1",
		"\t[X] Задача 1.2",
		"\tЗаголовок 1.1",
		"\t\t[ ] Задача 1.1.1",
		"\t\t[X] Задача 1.1.2",
		"Заголовок 2 #header_tag",
		"\t[ ] Задача 2.1 #task1_tag",
		"\t[X] Задача 2.2 #task2_tag",
	}
	assert.Equal(t, exp, ob.Indent().Build())

	ob = NewOutBuilder(elem)
	exp = []string{
		"Заголовок 1",
		"[ ] Задача 1.1",
		"[X] Задача 1.2",
		"Заголовок 1.1",
		"[ ] Задача 1.1.1",
		"[X] Задача 1.1.2",
		"Заголовок 2 #header_tag",
		"[ ] Задача 2.1 #task1_tag",
		"[X] Задача 2.2 #task2_tag",
	}
	assert.Equal(t, exp, ob.Build())

	ob = NewOutBuilder(elem)
	exp = []string{
		"Заголовок 1",
		"[ ] Задача 1.1",
		"[X] Задача 1.2",
		"Заголовок 1.1",
		"[ ] Задача 1.1.1",
		"[X] Задача 1.1.2",
		"Заголовок 2 #header_tag\t#header_tag",
		"[ ] Задача 2.1 #task1_tag\t#task1_tag #header_tag",
		"[X] Задача 2.2 #task2_tag\t#task2_tag #header_tag",
	}
	assert.Equal(t, exp, ob.ShowAllTags().Build())
}
