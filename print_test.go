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
	}
	ob := NewOutBuilder(elem)

	exp1 := []string{
		"Заголовок 1",
		"\t[ ] Задача 1.1",
		"\t[X] Задача 1.2",
		"\tЗаголовок 1.1",
		"\t\t[ ] Задача 1.1.1",
		"\t\t[X] Задача 1.1.2",
	}

	assert.Equal(t, exp1, ob.Build())
}
