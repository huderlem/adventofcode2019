package util

import (
	"bufio"
	"io/ioutil"
	"os"
	"strconv"
)

// ReadFileString reads a file as a string.
func ReadFileString(filepath string) string {
	bytes, err := ioutil.ReadFile(filepath)
	if err != nil {
		panic(err.Error())
	}
	return string(bytes)
}

// ReadFileLines reads a file as a string split into separate lines.
func ReadFileLines(filepath string) []string {
	f, err := os.Open(filepath)
	if err != nil {
		panic(err.Error())
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if scanner.Err() != nil {
		panic(scanner.Err())
	}
	return lines
}

// ReadFileInts reads a file as an line-separated list of integers.
func ReadFileInts(filepath string) []int {
	rawVals := ReadFileLines(filepath)
	vals := make([]int, len(rawVals))
	for i, rawVal := range rawVals {
		intVal, err := strconv.ParseInt(rawVal, 0, 64)
		if err != nil {
			panic(err)
		}
		// Note, casting int64 to int will truncate the value on some systems.
		// We won't worry about it here, since this code isn't intended to be
		// used for any serious workload. This makes it more convenient to
		// consume input data from Advent of code.
		vals[i] = int(intVal)
	}
	return vals
}
