package main

import (
	"fmt"
	"strings"

	"github.com/huderlem/adventofcode2019/util"
)

func readOrbits() map[string]string {
	lines := util.ReadFileLines("input.txt")
	orbits := make(map[string]string)
	for _, line := range lines {
		parts := strings.Split(line, ")")
		orbits[parts[1]] = parts[0]
	}
	return orbits
}

func countOrbitsForObject(object string, orbits map[string]string) int {
	count := 0
	curObject := object
	for {
		if _, ok := orbits[curObject]; !ok {
			return count
		}
		curObject = orbits[curObject]
		count++
	}
}

func part1() int {
	orbits := readOrbits()
	total := 0
	for object := range orbits {
		total += countOrbitsForObject(object, orbits)
	}
	return total
}

func getOrbitStepsForObject(object string, orbits map[string]string) map[string]int {
	steps := make(map[string]int)
	count := 0
	curObject := object
	for {
		steps[curObject] = count
		if _, ok := orbits[curObject]; !ok {
			return steps
		}
		curObject = orbits[curObject]
		count++
	}
}

func part2() int {
	orbits := readOrbits()
	youOrbitSteps := getOrbitStepsForObject(orbits["YOU"], orbits)
	santaOrbitSteps := getOrbitStepsForObject(orbits["SAN"], orbits)
	curObject := "YOU"
	for {
		var ok bool
		if curObject, ok = orbits[curObject]; !ok {
			return -1
		}
		if santaSteps, ok := santaOrbitSteps[curObject]; ok {
			youSteps := youOrbitSteps[curObject]
			return youSteps + santaSteps
		}
	}
}

func main() {
	fmt.Println("Part 1 Answer:", part1())
	fmt.Println("Part 2 Answer:", part2())
}
