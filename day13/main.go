package main

import (
	"fmt"
	"strings"

	"github.com/huderlem/adventofcode2019/intcode"
)

type entity struct {
	x, y int
	kind int
}

func entityDisplay(kind int) rune {
	switch kind {
	case 0:
		return ' '
	case 1:
		return '#'
	case 2:
		return 'B'
	case 3:
		return '_'
	case 4:
		return 'O'
	default:
		return '?'
	}
}

func part1() int {
	program := intcode.ReadProgram("input.txt")
	outputs := []int{}
	intcode.ExecuteProgram(program, func() int {
		return 0
	}, func(val int) {
		outputs = append(outputs, val)
	})
	numBlocks := 0
	for i := 0; i < len(outputs); i += 3 {
		if outputs[i+2] == 2 {
			numBlocks++
		}
	}
	return numBlocks
}

func render(state []entity) {
	w, h := 37, 22
	grid := make([]rune, w*h)
	score := 0
	for _, e := range state {
		if e.x == -1 && e.y == 0 {
			score = e.kind
		} else {
			grid[e.x+e.y*w] = entityDisplay(e.kind)
		}
	}
	for j := 0; j < h; j++ {
		var sb strings.Builder
		for i := 0; i < w; i++ {
			sb.WriteRune(grid[i+j*w])
		}
		fmt.Println(sb.String())
	}
	fmt.Println("Score:", score)
}

func part2() int {
	program := intcode.ReadProgram("input.txt")
	program[0] = 2
	state := []entity{}
	outputCount := 0
	outputEntity := entity{}
	ballXPos := 0
	paddleXPos := 0
	score := 0
	intcode.ExecuteProgram(program, func() int {
		// render(state)
		if paddleXPos > ballXPos {
			return -1
		} else if paddleXPos < ballXPos {
			return 1
		}
		return 0
	}, func(val int) {
		switch outputCount % 3 {
		case 0:
			outputEntity.x = val
		case 1:
			outputEntity.y = val
		case 2:
			outputEntity.kind = val
		}
		outputCount++
		if outputCount == 3 {
			outputCount = 0
			state = append(state, outputEntity)
			if outputEntity.kind == 4 {
				ballXPos = outputEntity.x
			} else if outputEntity.kind == 3 {
				paddleXPos = outputEntity.x
			} else if outputEntity.x == -1 && outputEntity.y == 0 {
				score = outputEntity.kind
			}
		}
	})
	return score
}

func main() {
	fmt.Println("Part 1 Answer:", part1())
	fmt.Println("Part 2 Answer:", part2())
}
