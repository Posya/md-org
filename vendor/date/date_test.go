package date

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

//var location = time.Now().Location()

func TestParseRepetition(t *testing.T) {
	ins := [][2]string{
		{"5", "d"},
		{"3", "w"},
		{"10", "y"},
	}

	exp := []Repetition{
		Repetition{"+5d", 5, d},
		Repetition{"+3w", 3, w},
		Repetition{"+10y", 10, y},
	}

	if len(ins) != len(exp) {
		t.Fatal("Error in unit test: ins and exp has different length!")
	}

	for i := range exp {
		v, err := parseRepetition(ins[i][0], ins[i][1])
		assert.NoError(t, err)
		assert.Equal(t, exp[i], *v)
	}
}

func TestParseRepetitionErrors(t *testing.T) {
	ins := [][2]string{
		{"5", "D"},
		{"-3", "w"},
		{"10", "asd"},
		{"1", ""},
	}

	for i := range ins {
		_, err := parseRepetition(ins[i][0], ins[i][1])
		assert.Error(t, err, "in line: "+ins[i][0]+"; "+ins[i][1])
	}
}

func TestCombineResults(t *testing.T) {
	ins := [][6]string{
		{"10.01.2018", "15:00", "", "", "", ""},
		{"10.01.2018", "15:00", "17.01.2018", "", "", ""},
		{"10.01.2018", "15:00", "17.01.2018", "17:00", "", ""},
		{"10.01.2018", "15:00", "17.01.2018", "17:00", "1", "w"},
		{"10.01.2018", "15:00", "", "", "1", "w"},
	}

	exp := []string{
		"@10.01.2018_15:00",
		"@10.01.2018_15:00-17.01.2018",
		"@10.01.2018_15:00-17.01.2018_17:00",
		"@10.01.2018_15:00-17.01.2018_17:00+1w",
		"@10.01.2018_15:00+1w",
	}

	if len(ins) != len(exp) {
		t.Fatal("Error in unit test: ins and exp has different length!")
	}

	for i := range exp {
		assert.Equal(t, exp[i], combineResults(ins[i][0], ins[i][1], ins[i][2], ins[i][3], ins[i][4], ins[i][5]))
	}
}

func TestParseOneDateTime(t *testing.T) {

	ins := [][2]string{
		{"1.01.2018", "15:00"},
		{"12.01.2018", ""},
	}

	exp := []time.Time{
		time.Date(2018, time.January, 1, 15, 0, 0, 0, location),
		time.Date(2018, time.January, 12, 0, 0, 0, 0, location),
	}

	if len(ins) != len(exp) {
		t.Fatal("Error in unit test: ins and exp has different length!")
	}

	for i := range exp {
		v, err := parseOneDateTime(ins[i][0], ins[i][1])
		assert.NoError(t, err)
		assert.Equal(t, exp[i], *v)
	}
}

func TestParseOneDateTimeErrors(t *testing.T) {
	ins := [][2]string{
		{"1.01.18", "15:00"},
		{"", ""},
		{"31.02.2018", ""},
		{"1.02.2018", "25:00"},
	}

	for i := range ins {
		_, err := parseOneDateTime(ins[i][0], ins[i][1])
		assert.Error(t, err, "in line: "+ins[i][0]+"; "+ins[i][1])
	}
}

func TestParse(t *testing.T) {
	ins := [][6]string{
		{"12.01.2018"},
		{"12.01.2018", "15:00", "13.01.2018", "16:00", "1", "w"},
	}

	exp := []Date{
		NewDate(
			time.Date(2018, time.January, 12, 0, 0, 0, 0, location),
			nil,
			nil,
		),
		NewDate(
			time.Date(2018, time.January, 12, 15, 0, 0, 0, location),
			time.Date(2018, time.January, 13, 16, 0, 0, 0, location),
			Repetition{"+1w", 1, w},
		),
	}

	if len(ins) != len(exp) {
		t.Fatal("Error in unit test: ins and exp has different length!")
	}

	for i := range exp {
		// parse(dateFrom, timeFrom, dateTo, timeTo, repeatVal, repeatUnit string) (*date, error)
		v, err := Parse(ins[i][0], ins[i][1], ins[i][2], ins[i][3], ins[i][4], ins[i][5])
		assert.NoError(t, err)
		assert.True(t, exp[i].Equal(*v), "values: expected: ", exp[i].Print(), "real: ", v.Print())
	}
}

func TestDateEqual(t *testing.T) {
	ins := [][2]Date{
		{ // 1
			NewDate(
				time.Date(2018, time.January, 12, 0, 0, 0, 0, location),
				nil,
				nil,
			),
			NewDate(
				time.Date(2018, time.January, 12, 0, 0, 0, 0, location),
				nil,
				nil,
			),
		},
		{ // -1
			NewDate(
				time.Date(2018, time.January, 12, 0, 0, 0, 0, location),
				nil,
				nil,
			),
			NewDate(
				time.Date(2018, time.January, 12, 0, 0, 0, 0, location),
				time.Date(2018, time.January, 12, 0, 0, 0, 0, location),
				nil,
			),
		},
		{ // 2
			NewDate(
				time.Date(2018, time.January, 12, 0, 0, 0, 0, location),
				time.Date(2018, time.January, 12, 0, 0, 0, 0, location),
				nil,
			),
			NewDate(
				time.Date(2018, time.January, 12, 0, 0, 0, 0, location),
				time.Date(2018, time.January, 12, 0, 0, 0, 0, location),
				nil,
			),
		},
		{ // -2
			NewDate(
				time.Date(2018, time.January, 12, 0, 0, 0, 0, location),
				time.Date(2018, time.January, 12, 0, 0, 0, 0, location),
				nil,
			),
			NewDate(
				time.Date(2018, time.January, 12, 0, 0, 0, 0, location),
				time.Date(2018, time.January, 13, 0, 0, 0, 0, location),
				nil,
			),
		},
		{ // 3
			NewDate(
				time.Date(2018, time.January, 12, 0, 0, 0, 0, location),
				time.Date(2018, time.January, 12, 0, 0, 0, 0, location),
				Repetition{"+1w", 1, w},
			),
			NewDate(
				time.Date(2018, time.January, 12, 0, 0, 0, 0, location),
				time.Date(2018, time.January, 12, 0, 0, 0, 0, location),
				Repetition{"+1w", 1, w},
			),
		},
		{ // -3
			NewDate(
				time.Date(2018, time.January, 12, 0, 0, 0, 0, location),
				time.Date(2018, time.January, 12, 0, 0, 0, 0, location),
				Repetition{"+1w", 1, w},
			),
			NewDate(
				time.Date(2018, time.January, 12, 0, 0, 0, 0, location),
				time.Date(2018, time.January, 12, 0, 0, 0, 0, location),
				nil,
			),
		},
		{ // 4
			NewDate(
				time.Date(2018, time.January, 12, 0, 0, 0, 0, location),
				nil,
				Repetition{"+1w", 1, w},
			),
			NewDate(
				time.Date(2018, time.January, 12, 0, 0, 0, 0, location),
				nil,
				Repetition{"+1w", 1, w},
			),
		},
		{ // -4
			NewDate(
				time.Date(2018, time.January, 12, 0, 0, 0, 0, location),
				nil,
				Repetition{"+1w", 1, w},
			),
			NewDate(
				time.Date(2018, time.January, 12, 0, 0, 0, 0, location),
				nil,
				nil,
			),
		},
	}

	exp := []bool{
		true, // 1
		false,
		true, // 2
		false,
		true, // 3
		false,
		true, // 4
		false,
	}

	if len(ins) != len(exp) {
		t.Fatal("Error in unit test: ins and exp has different length!")
	}

	for i := range exp {
		a := ins[i][0]
		b := ins[i][1]

		r := a.Equal(b)
		assert.Equal(t, exp[i], r, "d1: "+a.Print(), "d2: "+b.Print())
	}
}

func TestDatePrint(t *testing.T) {

	ins := []Date{
		NewDate(
			time.Date(2018, time.January, 12, 17, 0, 0, 0, location),
			time.Date(2018, time.January, 13, 20, 0, 0, 0, location),
			Repetition{"+1w", 1, w},
		),
		NewDate(
			time.Date(2018, time.January, 12, 0, 0, 0, 0, location),
			time.Date(2018, time.January, 13, 0, 0, 0, 0, location),
			Repetition{"+1w", 1, w},
		),
	}

	exp := []string{
		"@12.01.2018_17:00-13.01.2018_20:00+1w",
		"@12.01.2018-13.01.2018+1w",
	}

	if len(ins) != len(exp) {
		t.Fatal("Error in unit test: ins and exp has different length!")
	}

	for i := range exp {
		assert.Equal(t, exp[i], ins[i].Print())
	}
}

func TestNewDate(t *testing.T) {

	ins := [][4]interface{}{
		{
			"@12.01.2018_17:00-13.01.2018_20:00+1w",
			time.Date(2018, time.January, 12, 17, 0, 0, 0, location),
			time.Date(2018, time.January, 13, 20, 0, 0, 0, location),
			Repetition{"+1w", 1, w},
		},
		{
			"@12.01.2018-13.01.2018+1w",
			time.Date(2018, time.January, 12, 0, 0, 0, 0, location),
			time.Date(2018, time.January, 13, 0, 0, 0, 0, location),
			Repetition{"+1w", 1, w},
		},
	}

	for _, i := range ins {
		d := NewDate(
			i[1].(time.Time),
			i[2].(time.Time),
			i[3].(Repetition),
		)

		assert.Equal(t, i[0].(string), d.Original)
		assert.True(t, i[1].(time.Time).Equal(d.From))
		assert.True(t, i[2].(time.Time).Equal(*d.To))
		assert.True(t, i[3].(Repetition).Equal(*d.Rep))
	}
}
