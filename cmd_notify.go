package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"text/tabwriter"
	"time"
)

type cmdNotify struct {
	Verbose bool `short:"v" long:"verbose" description:"verbose output"`
	Range   int  `short:"r" long:"range" default:"5" description:"+/- minutes from now"`
}

func (cn *cmdNotify) Execute(args []string) error {
	dir, err := DirFromOptsOrCurrent()
	if err != nil {
		return err
	}
	sort.Strings(dir)

	t := time.Now().Truncate(time.Minute)
	t1 := t.Add(time.Minute * time.Duration((-1)*cn.Range))
	t2 := t.Add(time.Minute * time.Duration(cn.Range))
	t1str := t1.Format("2006.01.02 15:04")
	t2str := t2.Format("2006.01.02 15:04")

	for _, file := range dir {
		lines, err := ReadFile(file)
		if err != nil {
			return err
		}

		elements, err := parse(lines)
		if err != nil {
			return err
		}

		elements = filterElements(elements, func(el element) bool {
			v, ok := el.(task)
			if ok {
				return len(v.date) > 10
			}
			return false
		})

		elements = filterBetveen(elements, t1str, t2str)

		elements = filterTasks(elements, "notdone")

		elements = sortTasks(elements, "date")

		if len(elements) == 0 {
			continue
		}

		taskFuncs := []printFunc{
			func(el element) (s string, skip bool) {
				if v, ok := el.(task); ok {
					return v.date, false
				}
				panic("Wrong type: had to be task")
			},
			func(el element) (s string, skip bool) {
				return file, false
			},
			func(el element) (s string, skip bool) {
				if v, ok := el.(task); ok {
					return strconv.Itoa(v.n), false
				}
				panic("Wrong type: had to be task")
			},
			func(el element) (s string, skip bool) {
				if v, ok := el.(task); ok {
					isDone := "- [ ] "
					if v.done {
						isDone = "- [X] "
					}
					return isDone + v.text, false
				}
				panic("Wrong type: had to be task")
			},
		}

		w := tabwriter.NewWriter(os.Stdout, 0, 8, 1, '\t', 0)

		for _, el := range elements {
			l := prepareToPrint(el, []printFunc{}, taskFuncs)
			fmt.Fprintln(w, strings.Join(l, "\t"))
		}

		err = w.Flush()
		if err != nil {
			return err
		}
	}

	return nil
}
