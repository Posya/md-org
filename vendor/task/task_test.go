package task

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseTask(t *testing.T) {
	var in string
	in = "- [ ] Task 1"
	h, err := parseTask(in)
	assert.Nil(t, err)
	if assert.NotNil(t, h) {
		assert.Equal(t, in, h.original)
		assert.Equal(t, 0, h.level)
		assert.Equal(t, "Task 1", h.text)
		assert.Equal(t, 0, len(h.tags))
	}

}

func TestParseWrongTask(t *testing.T) {

	headers := []string{
		"- [] Task",
		"- [z] Task",
	}

	for _, in := range headers {
		h, err := parseTask(in)
		assert.NotNil(t, err)
		assert.Nil(t, h)
	}
}

func TestParseNotATask(t *testing.T) {

	headers := []string{}

	for _, in := range headers {
		h, err := parseTask(in)
		assert.Nil(t, err)
		assert.Nil(t, h)
	}
}

func TestParseDate(t *testing.T) {
	ins := [][]string{
		{"@10.01.2018_15:00+1w"},
		{"@10.01.2018+1w"},
		{"@10.01.2018_15:00-17.01.2018_17:00+1w"},
		// Wrong dated
		{"@10.01.2018 15:00"},
	}

	exp := []string{
		"@10.01.2018_15:00+1w",
		"@10.01.2018+1w",
		"@10.01.2018_15:00-17.01.2018_17:00+1w",
		// Wrong dated
		"",
	}

	if len(ins) != len(exp) {
		t.Fatal("Error in unit test: ins and exp has different length!")
	}

	for i := range exp {
		assert.Equal(t, exp[i], parseDates(ins[i]))
	}

}
