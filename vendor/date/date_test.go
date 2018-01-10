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

	exp := []repetition{
		repetition{"+5d", 5, d},
		repetition{"+3w", 3, w},
		repetition{"+10y", 10, y},
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

	exp := []date{
		newDate(
			time.Date(2018, time.January, 12, 0, 0, 0, 0, location),
			nil,
			nil,
		),
		newDate(
			time.Date(2018, time.January, 12, 15, 0, 0, 0, location),
			time.Date(2018, time.January, 13, 16, 0, 0, 0, location),
			repetition{"+1w", 1, w},
		),
	}

	if len(ins) != len(exp) {
		t.Fatal("Error in unit test: ins and exp has different length!")
	}

	for i := range exp {
		// parse(dateFrom, timeFrom, dateTo, timeTo, repeatVal, repeatUnit string) (*date, error)
		v, err := parse(ins[i][0], ins[i][1], ins[i][2], ins[i][3], ins[i][4], ins[i][5])
		assert.NoError(t, err)
		assert.True(t, exp[i].equal(*v), "values: expected: ", exp[i].print(), "real: ", v.print())
	}
}

func TestDateEqual(t *testing.T) {
	ins := [][2]date{
		{ // 1
			newDate(
				time.Date(2018, time.January, 12, 0, 0, 0, 0, location),
				nil,
				nil,
			),
			newDate(
				time.Date(2018, time.January, 12, 0, 0, 0, 0, location),
				nil,
				nil,
			),
		},
		{ // -1
			newDate(
				time.Date(2018, time.January, 12, 0, 0, 0, 0, location),
				nil,
				nil,
			),
			newDate(
				time.Date(2018, time.January, 12, 0, 0, 0, 0, location),
				time.Date(2018, time.January, 12, 0, 0, 0, 0, location),
				nil,
			),
		},
		{ // 2
			newDate(
				time.Date(2018, time.January, 12, 0, 0, 0, 0, location),
				time.Date(2018, time.January, 12, 0, 0, 0, 0, location),
				nil,
			),
			newDate(
				time.Date(2018, time.January, 12, 0, 0, 0, 0, location),
				time.Date(2018, time.January, 12, 0, 0, 0, 0, location),
				nil,
			),
		},
		{ // -2
			newDate(
				time.Date(2018, time.January, 12, 0, 0, 0, 0, location),
				time.Date(2018, time.January, 12, 0, 0, 0, 0, location),
				nil,
			),
			newDate(
				time.Date(2018, time.January, 12, 0, 0, 0, 0, location),
				time.Date(2018, time.January, 13, 0, 0, 0, 0, location),
				nil,
			),
		},
		{ // 3
			newDate(
				time.Date(2018, time.January, 12, 0, 0, 0, 0, location),
				time.Date(2018, time.January, 12, 0, 0, 0, 0, location),
				repetition{"+1w", 1, w},
			),
			newDate(
				time.Date(2018, time.January, 12, 0, 0, 0, 0, location),
				time.Date(2018, time.January, 12, 0, 0, 0, 0, location),
				repetition{"+1w", 1, w},
			),
		},
		{ // -3
			newDate(
				time.Date(2018, time.January, 12, 0, 0, 0, 0, location),
				time.Date(2018, time.January, 12, 0, 0, 0, 0, location),
				repetition{"+1w", 1, w},
			),
			newDate(
				time.Date(2018, time.January, 12, 0, 0, 0, 0, location),
				time.Date(2018, time.January, 12, 0, 0, 0, 0, location),
				nil,
			),
		},
		{ // 4
			newDate(
				time.Date(2018, time.January, 12, 0, 0, 0, 0, location),
				nil,
				repetition{"+1w", 1, w},
			),
			newDate(
				time.Date(2018, time.January, 12, 0, 0, 0, 0, location),
				nil,
				repetition{"+1w", 1, w},
			),
		},
		{ // -4
			newDate(
				time.Date(2018, time.January, 12, 0, 0, 0, 0, location),
				nil,
				repetition{"+1w", 1, w},
			),
			newDate(
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

		r := a.equal(b)
		assert.Equal(t, exp[i], r, "d1: "+a.print(), "d2: "+b.print())
	}
}

func TestDatePrint(t *testing.T) {

	ins := []date{
		newDate(
			time.Date(2018, time.January, 12, 17, 0, 0, 0, location),
			time.Date(2018, time.January, 13, 20, 0, 0, 0, location),
			repetition{"+1w", 1, w},
		),
		newDate(
			time.Date(2018, time.January, 12, 0, 0, 0, 0, location),
			time.Date(2018, time.January, 13, 0, 0, 0, 0, location),
			repetition{"+1w", 1, w},
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
		assert.Equal(t, exp[i], ins[i].print())
	}
}

func TestNewDate(t *testing.T) {

	ins := [][4]interface{}{
		{
			"@12.01.2018_17:00-13.01.2018_20:00+1w",
			time.Date(2018, time.January, 12, 17, 0, 0, 0, location),
			time.Date(2018, time.January, 13, 20, 0, 0, 0, location),
			repetition{"+1w", 1, w},
		},
		{
			"@12.01.2018-13.01.2018+1w",
			time.Date(2018, time.January, 12, 0, 0, 0, 0, location),
			time.Date(2018, time.January, 13, 0, 0, 0, 0, location),
			repetition{"+1w", 1, w},
		},
	}

	for _, i := range ins {
		d := newDate(
			i[1].(time.Time),
			i[2].(time.Time),
			i[3].(repetition),
		)

		assert.Equal(t, i[0].(string), d.original)
		assert.True(t, i[1].(time.Time).Equal(d.from))
		assert.True(t, i[2].(time.Time).Equal(*d.to))
		assert.True(t, i[3].(repetition).equal(*d.rep))
	}
}
