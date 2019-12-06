package main

import (
	"fmt"

	"github.com/huderlem/adventofcode2019/intcode"
)

func part1() int {
	program := intcode.ReadProgram("input.txt")
	program[1] = 12
	program[2] = 2
	intcode.ExecuteProgram(program, nil, nil)
	return program[0]
}

func findNounAndVerb(originalProgram []int) (int, int) {
	// Simply try all combinations of noun and verb values
	// until we find the correct answer.
	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			program := make([]int, len(originalProgram))
			copy(program, originalProgram)
			program[1] = noun
			program[2] = verb
			intcode.ExecuteProgram(program, nil, nil)
			if program[0] == 19690720 {
				return noun, verb
			}
		}
	}

	return 0, 0
}

func part2() int {
	program := intcode.ReadProgram("input.txt")
	noun, verb := findNounAndVerb(program)
	return 100*noun + verb
}

func main() {
	fmt.Println("Part 1 Answer:", part1())
	fmt.Println("Part 2 Answer:", part2())
}
