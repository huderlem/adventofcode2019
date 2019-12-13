package main

import (
	"fmt"
	"math"
	"regexp"
	"strconv"

	"github.com/huderlem/adventofcode2019/util"
)

type vec3 struct {
	x, y, z int
}

type moon struct {
	id       int
	pos      vec3
	velocity vec3
}

const (
	axisX   byte = 0b001
	axisY   byte = 0b010
	axisZ   byte = 0b100
	axisAll byte = 0b111
)

func (m moon) getEnergy() int {
	potential := math.Abs(float64(m.pos.x)) + math.Abs(float64(m.pos.y)) + math.Abs(float64(m.pos.z))
	kinetic := math.Abs(float64(m.velocity.x)) + math.Abs(float64(m.velocity.y)) + math.Abs(float64(m.velocity.z))
	return int(potential * kinetic)
}

func readMoons() []moon {
	lines := util.ReadFileLines("input.txt")
	moons := make([]moon, len(lines))
	r := regexp.MustCompile(`x=(-?\d+), y=(-?\d+), z=(-?\d+)`)
	for i, line := range lines {
		matches := r.FindStringSubmatch(line)
		x, _ := strconv.Atoi(matches[1])
		y, _ := strconv.Atoi(matches[2])
		z, _ := strconv.Atoi(matches[3])
		moons[i] = moon{
			id:       i,
			pos:      vec3{x, y, z},
			velocity: vec3{0, 0, 0},
		}
	}
	return moons
}

func getVelocityDelta(a, b int) int {
	if a < b {
		return 1
	}
	if a > b {
		return -1
	}
	return 0
}

func updateVelocity(moons []moon, axes byte) {
	for i := range moons {
		for _, moon := range moons {
			if moons[i].id == moon.id {
				continue
			}
			if axes&axisX != 0 {
				moons[i].velocity.x += getVelocityDelta(moons[i].pos.x, moon.pos.x)
			}
			if axes&axisY != 0 {
				moons[i].velocity.y += getVelocityDelta(moons[i].pos.y, moon.pos.y)
			}
			if axes&axisZ != 0 {
				moons[i].velocity.z += getVelocityDelta(moons[i].pos.z, moon.pos.z)
			}
		}
	}
}

func updatePositions(moons []moon) {
	for i := range moons {
		moons[i].pos.x += moons[i].velocity.x
		moons[i].pos.y += moons[i].velocity.y
		moons[i].pos.z += moons[i].velocity.z
	}
}

func calcTotalEnergy(moons []moon) int {
	energy := 0
	for _, moon := range moons {
		energy += moon.getEnergy()
	}
	return energy
}

func part1() int {
	moons := readMoons()
	for i := 0; i < 1000; i++ {
		updateVelocity(moons, axisAll)
		updatePositions(moons)
	}
	return calcTotalEnergy(moons)
}

func findRepetitionPoint(moons []moon, axis byte) int {
	states := make(map[string]int)
	for i := 0; ; i++ {
		state := fmt.Sprintf("%v", moons)
		if _, ok := states[state]; ok {
			return i
		}
		states[state] = i
		updateVelocity(moons, axis)
		updatePositions(moons)
	}
}

// Greatest common divisor (GCD) via Euclidean algorithm
// Borrowed from https://siongui.github.io/2017/06/03/go-find-lcm-by-gcd/
func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// Find Least Common Multiple (LCM) via greatest common divisor
// Borrowed from https://siongui.github.io/2017/06/03/go-find-lcm-by-gcd/
func lcm(a, b int, integers ...int) int {
	result := a * b / gcd(a, b)
	for i := 0; i < len(integers); i++ {
		result = lcm(result, integers[i])
	}
	return result
}

func part2() int {
	moons := readMoons()
	// Find state repetition by simulating a single axis at a time.
	// The first repetition for the entire system using all axes will
	// be the least common multiple of each axis's repetition.
	xRepetitionPoint := findRepetitionPoint(moons, axisX)
	yRepetitionPoint := findRepetitionPoint(moons, axisY)
	zRepetitionPoint := findRepetitionPoint(moons, axisZ)
	return lcm(xRepetitionPoint, yRepetitionPoint, zRepetitionPoint)
}

func main() {
	fmt.Println("Part 1 Answer:", part1())
	fmt.Println("Part 2 Answer:", part2())
}
