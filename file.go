package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func ReadFile(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func WriteFile(lines []string, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for _, line := range lines {
		fmt.Fprintln(w, line)
	}
	return w.Flush()
}

func GetDir() ([]string, error) {
	ret := []string{}

	files, err := ioutil.ReadDir("./*.md-org")
	if err != nil {
		return nil, err
	}
	for _, file := range files {
		ret = append(ret, file.Name())
	}

	return ret, nil
}

func DirFromOptsOrCurrent() ([]string, error) {
	if len(opts.File) > 0 {
		return opts.File, nil
	}
	return GetDir()
}
