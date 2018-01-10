package date

import (
	"errors"
	"strconv"
	"time"
)

// Date is struct to hold range of dates with possibility of repetition
type Date struct {
	Original string
	From     time.Time
	To       *time.Time
	Rep      *Repetition
}

// Repetition is struct to hold value and units of date repetition
type Repetition struct {
	Original string
	Value    int
	Interval Unit
}

// Unit is enum type for units of date repetition
type Unit int

const (
	h Unit = iota
	d
	w
	m
	y
)

var location = time.Now().Location()

func parseRepetition(val, interval string) (*Repetition, error) {
	var res Repetition

	res.Original = "+" + val + interval

	i, err := strconv.Atoi(val)
	if err != nil {
		return nil, err
	}
	if i <= 0 {
		return nil, errors.New("parseRepetition: value can`t be negative: " + val)
	}
	res.Value = i

	switch interval {
	case "h":
		res.Interval = h
	case "d":
		res.Interval = d
	case "w":
		res.Interval = w
	case "m":
		res.Interval = m
	case "y":
		res.Interval = y
	default:
		return nil, errors.New("parseRepetition: unknown interval: " + interval)
	}

	return &res, nil
}

// Equal compares two Repetition
func (r Repetition) Equal(r1 Repetition) bool {
	if r.Original != r1.Original ||
		r.Value != r1.Value ||
		r.Interval != r1.Interval {
		return false
	}
	return true
}

// Parse parses strings to return new Date
func Parse(dateFrom, timeFrom, dateTo, timeTo, repeatVal, repeatUnit string) (*Date, error) {
	var d Date
	var err error
	d.Original = combineResults(dateFrom, timeFrom, dateTo, timeTo, repeatVal, repeatUnit)

	from, err := parseOneDateTime(dateFrom, timeFrom)
	if err != nil {
		return nil, err
	}
	d.From = *from

	if dateTo != "" {
		to, err := parseOneDateTime(dateTo, timeTo)
		if err != nil {
			return nil, err
		}
		d.To = to
	}

	if repeatVal != "" && repeatUnit != "" {
		rep, err := parseRepetition(repeatVal, repeatUnit)
		if err != nil {
			return nil, err
		}
		d.Rep = rep
	}
	return &d, nil
}

func parseOneDateTime(d, t string) (*time.Time, error) {
	var res time.Time
	var err error
	if t == "" {
		res, err = time.ParseInLocation("2.01.2006", d, location)
	} else {
		res, err = time.ParseInLocation("2.01.2006 15:04", d+" "+t, location)
	}
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func combineResults(dateFrom, timeFrom, dateTo, timeTo, repeatVal, repeatUnit string) string {
	res := "@"
	res += dateFrom
	if timeFrom != "" {
		res += "_" + timeFrom
	}
	if dateTo != "" {
		res += "-" + dateTo
	}
	if timeTo != "" {
		res += "_" + timeTo
	}
	if repeatVal != "" {
		res += "+" + repeatVal + repeatUnit
	}
	return res
}

// NewDate is constructor for Date struct
func NewDate(from time.Time, to, rep interface{}) Date {
	var d Date

	d.From = from

	switch v := to.(type) {
	case time.Time:
		d.To = &v
	default:
		d.To = nil
	}

	switch v := rep.(type) {
	case Repetition:
		d.Rep = &v
	default:
		d.Rep = nil
	}

	d.Original = d.Print()

	return d
}

// Equal compares two Dates
func (d Date) Equal(d1 Date) bool {
	if d.Original != d1.Original {
		return false
	}
	if !d.From.Equal(d1.From) {
		return false
	}

	if d.To == nil || d1.To == nil {
		if d.To != d1.To {
			return false
		}
	} else {
		if !d.To.Equal(*d1.To) {
			return false
		}
	}

	if d.Rep == nil || d1.Rep == nil {
		if d.Rep != d1.Rep {
			return false
		}
	} else {
		if !d.Rep.Equal(*d1.Rep) {
			return false
		}
	}

	return true
}

// Print returns string representation of Date
func (d Date) Print() string {
	var from, to, rep string

	if d.From.Hour() == 0 && d.From.Minute() == 0 {
		from = d.From.Format("2.01.2006")
	} else {
		from = d.From.Format("2.01.2006_15:04")
	}

	if d.To != nil {
		if d.To.Hour() == 0 && d.To.Minute() == 0 {
			to = d.To.Format("-2.01.2006")
		} else {
			to = d.To.Format("-2.01.2006_15:04")
		}
	}

	if d.Rep != nil {
		rep = d.Rep.Original
	}

	res := "@" + from + to + rep

	return res
}
