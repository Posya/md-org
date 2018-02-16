package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHeaderEqual(t *testing.T) {
	ins := []header{
		header{1, []string{"tag1", "tag2", "tag3"}},
	}

	exp := []header{
		header{1, []string{"tag1", "tag2", "tag3"}},
	}

	if len(ins) != len(exp) {
		t.Fatal("Error in unit test: ins and exp has different length!")
	}

	for i := range exp {
		assert.True(t, exp[i].Equal(ins[i]))
	}

}

func TestContentEqual(t *testing.T) {
	ins := []context{
		context{
			[]header{
				header{1, []string{"tag1", "tag2", "tag3"}},
			},
		},
		context{
			[]header{
				header{
					1, []string{"#tag11", "#tag12", "#tag13"},
				},
				header{
					2, []string{"#tag21", "#tag22", "#tag23"},
				},
				header{
					4, []string{"#tag41", "#tag42", "#tag43"},
				},
			},
		},
	}

	exp := []context{
		context{
			[]header{
				header{1, []string{"tag1", "tag2", "tag3"}},
			},
		},
		context{
			[]header{
				header{
					1, []string{"#tag11", "#tag12", "#tag13"},
				},
				header{
					2, []string{"#tag21", "#tag22", "#tag23"},
				},
				header{
					4, []string{"#tag41", "#tag42", "#tag43"},
				},
			},
		},
	}

	if len(ins) != len(exp) {
		t.Fatal("Error in unit test: ins and exp has different length!")
	}

	for i := range exp {
		assert.True(t, exp[i].Equal(ins[i]))
	}

}
