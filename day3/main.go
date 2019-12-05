package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/huderlem/adventofcode2019/util"
)

type point struct {
	x, y int
}

// Read the input paths as sets of points visited by tracing the paths.
// Keep track of how many steps the wire took along the way to reach each
// individual point.
func readPaths() []map[point]int {
	lines := util.ReadFileLines("input.txt")
	paths := make([]map[point]int, len(lines))
	for i, line := range lines {
		instructions := strings.Split(line, ",")
		curPoint := point{0, 0}
		steps := 0
		path := map[point]int{}
		for _, instruction := range instructions {
			num, _ := strconv.Atoi(instruction[1:])
			for i := 0; i < num; i++ {
				switch instruction[0] {
				case 'U':
					curPoint.y++
				case 'R':
					curPoint.x++
				case 'D':
					curPoint.y--
				case 'L':
					curPoint.x--
				}
				steps++
				if _, ok := path[curPoint]; !ok {
					path[curPoint] = steps
				}
			}
		}
		paths[i] = path
	}
	return paths
}

func manhattanDistance(p1, p2 point) int {
	xDist := p1.x - p2.x
	if xDist < 0 {
		xDist = -xDist
	}
	yDist := p1.y - p2.y
	if yDist < 0 {
		yDist = -yDist
	}
	return xDist + yDist
}

func part1() int {
	paths := readPaths()
	intersections := []point{}
	for point := range paths[1] {
		if _, ok := paths[0][point]; ok {
			intersections = append(intersections, point)
		}
	}
	zero := point{0, 0}
	minDistance := math.MaxInt64
	for _, intersection := range intersections {
		dist := manhattanDistance(zero, intersection)
		if dist < minDistance {
			minDistance = dist
		}
	}
	return minDistance
}

func part2() int {
	paths := readPaths()
	intersections := make(map[point]int)
	for point := range paths[1] {
		if _, ok := paths[0][point]; ok {
			intersections[point] = paths[0][point] + paths[1][point]
		}
	}
	minSteps := math.MaxInt64
	for _, steps := range intersections {
		if steps < minSteps {
			minSteps = steps
		}
	}
	return minSteps
}

func main() {
	fmt.Println("Part 1 Answer:", part1())
	fmt.Println("Part 2 Answer:", part2())
}
