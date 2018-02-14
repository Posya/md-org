package main

import (
	"io"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func getNextFunc(lines []string) func() (string, error) {
	i := 0
	var next = func() (string, error) {
		if len(lines) == i {
			return "", io.EOF
		}

		ret := lines[i]
		i++
		return ret, nil
	}
	return next
}

func TestParse(t *testing.T) {
	ins := [][]string{
		[]string{"Первая строка", "# Заголовок 1", "- [ ] First task"},
	}

	exp := [][]task{
		[]task{
			task{3, "- [ ] First task", "First task", []string{}, time.Time{}},
		},
	}

	if len(ins) != len(exp) {
		t.Fatal("Error in unit test: ins and exp has different length!")
	}

	for i := range exp {
		v, err := parse(getNextFunc(ins[i]))
		assert.NoError(t, err)
		assert.Equal(t, exp[i], v)
	}

}
