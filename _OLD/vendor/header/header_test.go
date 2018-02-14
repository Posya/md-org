package header

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseHeader(t *testing.T) {
	var in string
	in = "# Header 1"
	h, err := parseHeader(in)
	assert.Nil(t, err)
	if assert.NotNil(t, h) {
		assert.Equal(t, in, h.original)
		assert.Equal(t, 1, h.level)
		assert.Equal(t, "Header 1", h.text)
		assert.Equal(t, 0, len(h.tags))
	}

	in = "   	# Header 1 # asdfasdf #asdf #asdf1 #1234"
	h, err = parseHeader(in)
	assert.Nil(t, err)
	if assert.NotNil(t, h) {
		assert.Equal(t, in, h.original)
		assert.Equal(t, 1, h.level)
		assert.Equal(t, "Header 1 # asdfasdf ", h.text)
		assert.Equal(t, "#asdf", h.tags[0])
		assert.Equal(t, "#asdf1", h.tags[1])
		assert.Equal(t, "#1234", h.tags[2])
	}

	in = " ### 1. Another header"
	h, err = parseHeader(in)
	assert.Nil(t, err)
	if assert.NotNil(t, h) {
		assert.Equal(t, in, h.original)
		assert.Equal(t, 3, h.level)
		assert.Equal(t, "1. Another header", h.text)
		assert.Equal(t, 0, len(h.tags))
	}
}

func TestParseWrongHeader(t *testing.T) {

	headers := []string{
		"# Bad tags #asdf this is error",
		"# Date in tags #asdf @24.05.20018",
		"# Date in tags @24.05.20018",
	}

	for _, in := range headers {
		h, err := parseHeader(in)
		assert.NotNil(t, err)
		assert.Nil(t, h)
	}
}

func TestParseNotAHeader(t *testing.T) {

	headers := []string{
		"## # Not a header",
		"a # Not a header",
		"#Not a header",
	}

	for _, in := range headers {
		h, err := parseHeader(in)
		assert.Nil(t, err)
		assert.Nil(t, h)
	}
}
