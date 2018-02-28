package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHeaderEqual(t *testing.T) {
	ins := [][]header{
		[]header{
			header{12, 1, 0, "", []string{"tag1", "tag2", "tag3"}},
			header{12, 1, 0, "", []string{"tag1", "tag2", "tag3"}},
		},
		[]header{
			header{12, 1, 0, "", []string{"tag1", "tag2", "tag3"}},
			header{12, 1, 0, "", []string{"tag3", "tag2", "tag3"}},
		},
		[]header{
			header{12, 1, 0, "", []string{"tag1", "tag2", "tag3"}},
			header{12, 1, 0, "a", []string{"tag1", "tag2", "tag3"}},
		},
	}

	exp := []bool{
		true,
		false,
		false,
	}

	if len(ins) != len(exp) {
		t.Fatal("Error in unit test: ins and exp has different length!")
	}

	for i := range exp {
		assert.Equal(t, exp[i], ins[i][0].Equal(ins[i][1]))
	}
}
