// Solution for Advent of Code 2019 -- Day 1
// https://adventofcode.com/2019/day/1

package main

import (
	"fmt"

	"github.com/huderlem/adventofcode2019/util"
)

func getRequiredFuelForMass(mass int) int {
	return mass/3 - 2
}

func part1() int {
	masses := util.ReadFileInts("input.txt")
	requiredFuel := 0
	for _, val := range masses {
		requiredFuel += getRequiredFuelForMass(val)
	}
	return requiredFuel
}

func getRequiredFuel(mass int) int {
	requiredFuel := getRequiredFuelForMass(mass)
	additionalFuel := getRequiredFuelForMass(requiredFuel)
	for additionalFuel > 0 {
		requiredFuel += additionalFuel
		additionalFuel = getRequiredFuelForMass(additionalFuel)
	}
	return requiredFuel
}

func part2() int {
	vals := util.ReadFileInts("input.txt")
	requiredFuel := 0
	for _, val := range vals {
		requiredFuel += getRequiredFuel(val)
	}
	return requiredFuel
}

func main() {
	fmt.Println("Part 1 Answer:", part1())
	fmt.Println("Part 2 Answer:", part2())
}
