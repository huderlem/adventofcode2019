package main

import (
	"fmt"

	"github.com/huderlem/adventofcode2019/intcode"
)

func part1() int {
	program := intcode.ReadProgram("input.txt")
	result := 0
	intcode.ExecuteProgram(program, func() int {
		return 1
	}, func(val int) {
		result = val
	})
	return result
}

func part2() int {
	program := intcode.ReadProgram("input.txt")
	result := 0
	intcode.ExecuteProgram(program, func() int {
		return 2
	}, func(val int) {
		result = val
	})
	return result
}

func main() {
	fmt.Println("Part 1 Answer:", part1())
	fmt.Println("Part 2 Answer:", part2())
}
