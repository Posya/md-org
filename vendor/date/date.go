package date

import (
	"errors"
	"strconv"
	"time"
)

type date struct {
	original string
	from     time.Time
	to       *time.Time
	rep      *repetition
}

type repetition struct {
	original string
	value    int
	interval unit
}

type unit int

const (
	h unit = iota
	d
	w
	m
	y
)

var location = time.Now().Location()

func parseRepetition(val, interval string) (*repetition, error) {
	var res repetition

	res.original = "+" + val + interval

	i, err := strconv.Atoi(val)
	if err != nil {
		return nil, err
	}
	if i <= 0 {
		return nil, errors.New("parseRepetition: value can`t be negative: " + val)
	}
	res.value = i

	switch interval {
	case "h":
		res.interval = h
	case "d":
		res.interval = d
	case "w":
		res.interval = w
	case "m":
		res.interval = m
	case "y":
		res.interval = y
	default:
		return nil, errors.New("parseRepetition: unknown interval: " + interval)
	}

	return &res, nil
}

func (r repetition) equal(r1 repetition) bool {
	if r.original != r1.original ||
		r.value != r1.value ||
		r.interval != r1.interval {
		return false
	}
	return true
}

func parse(dateFrom, timeFrom, dateTo, timeTo, repeatVal, repeatUnit string) (*date, error) {
	var d date
	var err error
	d.original = combineResults(dateFrom, timeFrom, dateTo, timeTo, repeatVal, repeatUnit)

	from, err := parseOneDateTime(dateFrom, timeFrom)
	if err != nil {
		return nil, err
	}
	d.from = *from

	if dateTo != "" {
		to, err := parseOneDateTime(dateTo, timeTo)
		if err != nil {
			return nil, err
		}
		d.to = to
	}

	if repeatVal != "" && repeatUnit != "" {
		rep, err := parseRepetition(repeatVal, repeatUnit)
		if err != nil {
			return nil, err
		}
		d.rep = rep
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

func newDate(from time.Time, to, rep interface{}) date {
	var d date

	d.from = from

	switch v := to.(type) {
	case time.Time:
		d.to = &v
	default:
		d.to = nil
	}

	switch v := rep.(type) {
	case repetition:
		d.rep = &v
	default:
		d.rep = nil
	}

	d.original = d.print()

	return d
}

func (d date) equal(d1 date) bool {
	if d.original != d1.original {
		return false
	}
	if !d.from.Equal(d1.from) {
		return false
	}

	if d.to == nil || d1.to == nil {
		if d.to != d1.to {
			return false
		}
	} else {
		if !d.to.Equal(*d1.to) {
			return false
		}
	}

	if d.rep == nil || d1.rep == nil {
		if d.rep != d1.rep {
			return false
		}
	} else {
		if !d.rep.equal(*d1.rep) {
			return false
		}
	}

	return true
}

func (d date) print() string {
	var from, to, rep string

	if d.from.Hour() == 0 && d.from.Minute() == 0 {
		from = d.from.Format("2.01.2006")
	} else {
		from = d.from.Format("2.01.2006_15:04")
	}

	if d.to != nil {
		if d.to.Hour() == 0 && d.to.Minute() == 0 {
			to = d.to.Format("-2.01.2006")
		} else {
			to = d.to.Format("-2.01.2006_15:04")
		}
	}

	if d.rep != nil {
		rep = d.rep.original
	}

	res := "@" + from + to + rep

	return res
}
