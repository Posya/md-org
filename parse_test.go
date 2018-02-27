package main

import (
	"io"
	"testing"

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

// func TestParse(t *testing.T) {
// 	ins := [][]string{
// 		[]string{"Первая строка", "# Заголовок 1", "- [ ] First task"},
// 	}

// 	exp := [][]task{
// 		[]task{
// 			task{3, 1, "- [ ] First task", "First task", []string{}, ""},
// 		},
// 	}

// 	if len(ins) != len(exp) {
// 		t.Fatal("Error in unit test: ins and exp has different length!")
// 	}

// 	for i := range exp {
// 		v, err := parse(getNextFunc(ins[i]))
// 		assert.NoError(t, err)
// 		assert.Equal(t, exp[i], v)
// 	}

// }

// func TestParseHeader(t *testing.T) {
// 	con := []context{
// 		context{
// 			[]header{
// 				header{
// 					1, []string{"#tag1", "#tag2", "#tag3"},
// 				},
// 			},
// 		},
// 		context{
// 			[]header{
// 				header{
// 					1, []string{"#tag11", "#tag12", "#tag13"},
// 				},
// 				header{
// 					2, []string{"#tag21", "#tag22", "#tag23"},
// 				},
// 				header{
// 					4, []string{"#tag41", "#tag42", "#tag43"},
// 				},
// 			},
// 		},
// 	}

// 	ins := []string{
// 		"# Header 1 #tag11, #tag12, #tag13",
// 		"### Заголовок 1 #тег_1, #ещёТег #и_последний тег",
// 	}

// 	exp := []context{
// 		context{
// 			[]header{
// 				header{
// 					1, []string{"#tag11", "#tag12", "#tag13"},
// 				},
// 			},
// 		},
// 		context{
// 			[]header{
// 				header{
// 					1, []string{"#tag11", "#tag12", "#tag13"},
// 				},
// 				header{
// 					2, []string{"#tag21", "#tag22", "#tag23"},
// 				},
// 				header{
// 					3, []string{"#тег_1", "#ещёТег", "#и_последний"},
// 				},
// 			},
// 		},
// 	}

// 	if len(ins) != len(exp) && len(con) != len(exp) {
// 		t.Fatal("Error in unit test: ins, exp and con has different length!")
// 	}

// 	for i := range exp {
// 		v, err := parseHeader(con[i], ins[i])
// 		assert.NoError(t, err)
// 		if !exp[i].Equal(v) {
// 			t.Error(
// 				"Not equal:\n",
// 				fmt.Sprintf("Expected: %q\n", exp[i]),
// 				fmt.Sprintf("Actual: %q\n", v),
// 			)
// 		}
// 	}
// }

func TestCheckDate(t *testing.T) {
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
		assert.Equal(t, exp[i], checkDate(ins[i]))
	}
}
