package intcode

import (
	"strconv"
	"strings"

	"github.com/huderlem/adventofcode2019/util"
)

// Intcode operators.
const (
	opAdd          = 1
	opMultiply     = 2
	opInput        = 3
	opOutput       = 4
	opJumpTrue     = 5
	opJumpFalse    = 6
	opLessThan     = 7
	opEquals       = 8
	opRelativeBase = 9
	opTerminator   = 99
)

// Parameter modes.
const (
	modePosition  = 0
	modeImmediate = 1
	modeRelative  = 2
)

type intcodeInput func() int
type intcodeOutput func(int)

func getParamMode(index int, modes []int) int {
	if index < len(modes) {
		return modes[index]
	}
	return modePosition
}

func readParam(address, mode int, program []int, relativeBase int) int {
	switch mode {
	case modePosition:
		return program[program[address]]
	case modeImmediate:
		return program[address]
	case modeRelative:
		return program[relativeBase+program[address]]
	}
	panic("Invalid parameter mode")
}

func readLiteralParam(address, mode int, program []int, relativeBase int) int {
	switch mode {
	case modePosition, modeImmediate:
		return program[address]
	case modeRelative:
		return relativeBase + program[address]
	}
	panic("Invalid literal parameter mode")
}

// ExecuteProgram executes the given intcode program.
func ExecuteProgram(inputProgram []int, inputHandler intcodeInput, outputHandler intcodeOutput) {
	program := make([]int, len(inputProgram)*200)
	copy(program, inputProgram)
	if inputHandler == nil {
		inputHandler = func() int { return 1 }
	}
	if outputHandler == nil {
		outputHandler = func(int) {}
	}

	relativeBase := 0
	pc := 0
	for {
		opcode, parameterModes := getOpcodeInfo(program[pc])
		switch opcode {
		case opAdd:
			p1 := readParam(pc+1, getParamMode(0, parameterModes), program, relativeBase)
			p2 := readParam(pc+2, getParamMode(1, parameterModes), program, relativeBase)
			p3 := readLiteralParam(pc+3, getParamMode(2, parameterModes), program, relativeBase)
			program[p3] = p1 + p2
			pc += 4
		case opMultiply:
			p1 := readParam(pc+1, getParamMode(0, parameterModes), program, relativeBase)
			p2 := readParam(pc+2, getParamMode(1, parameterModes), program, relativeBase)
			p3 := readLiteralParam(pc+3, getParamMode(2, parameterModes), program, relativeBase)
			program[p3] = p1 * p2
			pc += 4
		case opInput:
			p1 := readLiteralParam(pc+1, getParamMode(0, parameterModes), program, relativeBase)
			program[p1] = inputHandler()
			pc += 2
		case opOutput:
			p1 := readParam(pc+1, getParamMode(0, parameterModes), program, relativeBase)
			outputHandler(p1)
			pc += 2
		case opJumpTrue:
			p1 := readParam(pc+1, getParamMode(0, parameterModes), program, relativeBase)
			p2 := readParam(pc+2, getParamMode(1, parameterModes), program, relativeBase)
			if p1 != 0 {
				pc = p2
			} else {
				pc += 3
			}
		case opJumpFalse:
			p1 := readParam(pc+1, getParamMode(0, parameterModes), program, relativeBase)
			p2 := readParam(pc+2, getParamMode(1, parameterModes), program, relativeBase)
			if p1 == 0 {
				pc = p2
			} else {
				pc += 3
			}
		case opLessThan:
			p1 := readParam(pc+1, getParamMode(0, parameterModes), program, relativeBase)
			p2 := readParam(pc+2, getParamMode(1, parameterModes), program, relativeBase)
			p3 := readLiteralParam(pc+3, getParamMode(2, parameterModes), program, relativeBase)
			if p1 < p2 {
				program[p3] = 1
			} else {
				program[p3] = 0
			}
			pc += 4
		case opEquals:
			p1 := readParam(pc+1, getParamMode(0, parameterModes), program, relativeBase)
			p2 := readParam(pc+2, getParamMode(1, parameterModes), program, relativeBase)
			p3 := readLiteralParam(pc+3, getParamMode(2, parameterModes), program, relativeBase)
			if p1 == p2 {
				program[p3] = 1
			} else {
				program[p3] = 0
			}
			pc += 4
		case opRelativeBase:
			p1 := readParam(pc+1, getParamMode(0, parameterModes), program, relativeBase)
			relativeBase += p1
			pc += 2
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
