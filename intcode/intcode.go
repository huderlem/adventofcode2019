package intcode

import (
	"strconv"
	"strings"

	"github.com/huderlem/adventofcode2019/util"
)

// Intcode operators.
const (
	opAdd        = 1
	opMultiply   = 2
	opInput      = 3
	opOutput     = 4
	opJumpTrue   = 5
	opJumpFalse  = 6
	opLessThan   = 7
	opEquals     = 8
	opTerminator = 99
)

// Parameter modes.
const (
	modePosition  = 0
	modeImmediate = 1
)

type intcodeInput func() int
type intcodeOutput func(int)

func getParamMode(index int, modes []int) int {
	if index < len(modes) {
		return modes[index]
	}
	return modePosition
}

func readParam(address, mode int, program []int) int {
	switch mode {
	case modePosition:
		return program[program[address]]
	case modeImmediate:
		return program[address]
	}
	panic("Invalid parameter mode")
}

// ExecuteProgram executes the given intcode program.
func ExecuteProgram(program []int, inputHandler intcodeInput, outputHandler intcodeOutput) {
	if inputHandler == nil {
		inputHandler = func() int { return 1 }
	}
	if outputHandler == nil {
		outputHandler = func(int) {}
	}

	pc := 0
	for {
		opcode, parameterModes := getOpcodeInfo(program[pc])
		switch opcode {
		case opAdd:
			p1 := readParam(pc+1, getParamMode(0, parameterModes), program)
			p2 := readParam(pc+2, getParamMode(1, parameterModes), program)
			program[program[pc+3]] = p1 + p2
			pc += 4
		case opMultiply:
			p1 := readParam(pc+1, getParamMode(0, parameterModes), program)
			p2 := readParam(pc+2, getParamMode(1, parameterModes), program)
			program[program[pc+3]] = p1 * p2
			pc += 4
		case opInput:
			program[program[pc+1]] = inputHandler()
			pc += 2
		case opOutput:
			p1 := readParam(pc+1, getParamMode(0, parameterModes), program)
			outputHandler(p1)
			pc += 2
		case opJumpTrue:
			p1 := readParam(pc+1, getParamMode(0, parameterModes), program)
			p2 := readParam(pc+2, getParamMode(1, parameterModes), program)
			if p1 != 0 {
				pc = p2
			} else {
				pc += 3
			}
		case opJumpFalse:
			p1 := readParam(pc+1, getParamMode(0, parameterModes), program)
			p2 := readParam(pc+2, getParamMode(1, parameterModes), program)
			if p1 == 0 {
				pc = p2
			} else {
				pc += 3
			}
		case opLessThan:
			p1 := readParam(pc+1, getParamMode(0, parameterModes), program)
			p2 := readParam(pc+2, getParamMode(1, parameterModes), program)
			p3 := program[pc+3]
			if p1 < p2 {
				program[p3] = 1
			} else {
				program[p3] = 0
			}
			pc += 4
		case opEquals:
			p1 := readParam(pc+1, getParamMode(0, parameterModes), program)
			p2 := readParam(pc+2, getParamMode(1, parameterModes), program)
			p3 := program[pc+3]
			if p1 == p2 {
				program[p3] = 1
			} else {
				program[p3] = 0
			}
			pc += 4
		case opTerminator:
			return
		}
	}
}

func getOpcodeInfo(rawOpcode int) (int, []int) {
	opcode := rawOpcode % 100
	parameterModes := []int{}
	modeVals := rawOpcode / 100
	for modeVals > 0 {
		parameterModes = append(parameterModes, modeVals%10)
		modeVals /= 10
	}

	return opcode, parameterModes
}

// ReadProgram parses an intcode program from a file.
func ReadProgram(filepath string) []int {
	rawIntcode := util.ReadFileString(filepath)
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
