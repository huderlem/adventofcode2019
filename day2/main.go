package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/huderlem/adventofcode2019/util"
)

func readProgram() []int {
	rawIntcode := util.ReadFileString("input.txt")
	intcodes := strings.Split(rawIntcode, ",")
	program := make([]int, len(intcodes))
	for i, code := range intcodes {
		var err error
		program[i], err = strconv.Atoi(code)
		if err != nil {
			panic(err)
		}
	}
	return program
}

// Intcode operators.
const (
	opAdd        = 1
	opMultiply   = 2
	opTerminator = 99
)

func executeIntcodeProgram(program []int) {
	pc := 0
	for {
		switch program[pc] {
		case opAdd:
			program[program[pc+3]] = program[program[pc+1]] + program[program[pc+2]]
			pc += 4
		case opMultiply:
			program[program[pc+3]] = program[program[pc+1]] * program[program[pc+2]]
			pc += 4
		case opTerminator:
			return
		}
	}
}

func part1() int {
	program := readProgram()
	program[1] = 12
	program[2] = 2
	executeIntcodeProgram(program)
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
			executeIntcodeProgram(program)
			if program[0] == 19690720 {
				return noun, verb
			}
		}
	}

	return 0, 0
}

func part2() int {
	program := readProgram()
	noun, verb := findNounAndVerb(program)
	return 100*noun + verb
}

func main() {
	fmt.Println("Part 1 Answer:", part1())
	fmt.Println("Part 2 Answer:", part2())
}
