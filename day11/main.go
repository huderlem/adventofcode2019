package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"os"

	"github.com/huderlem/adventofcode2019/intcode"
)

const (
	cBlack = 0
	cWhite = 1

	tLeft  = 0
	tRight = 1

	fUp    = 0
	fRight = 1
	fDown  = 2
	fLeft  = 3
)

type point struct {
	x, y int
}

type robot struct {
	point
	facing int
}

func (r robot) turn(direction int) robot {
	switch direction {
	case tLeft:
		r.facing--
	case tRight:
		r.facing++
	}
	r.facing = (r.facing + 4) % 4
	return r
}

func (r robot) moveForward() robot {
	switch r.facing {
	case fUp:
		r.y--
	case fRight:
		r.x++
	case fDown:
		r.y++
	case fLeft:
		r.x--
	}
	return r
}

func paint(paintedCells map[point]int, program []int) map[point]int {
	robot := robot{point{0, 0}, 0}
	outputCount := 0
	intcode.ExecuteProgram(program, func() int {
		if color, ok := paintedCells[robot.point]; ok {
			return color
		}
		return 0
	}, func(val int) {
		switch outputCount % 2 {
		case 0:
			paintedCells[robot.point] = val
		case 1:
			robot = robot.turn(val).moveForward()
		}
		outputCount++
	})
	return paintedCells
}

func part1() int {
	program := intcode.ReadProgram("input.txt")
	paintedCells := make(map[point]int)
	paint(paintedCells, program)
	return len(paintedCells)
}

func getBoundaries(paintedCells map[point]int) (int, int, int, int) {
	left, right, top, bottom := math.MaxInt64, math.MinInt64, math.MinInt64, math.MaxInt64
	for p := range paintedCells {
		if p.x < left {
			left = p.x
		}
		if p.x+1 > right {
			right = p.x + 1
		}
		if p.y < bottom {
			bottom = p.y
		}
		if p.y+1 > top {
			top = p.y + 1
		}
	}
	return left, right, top, bottom
}

func renderImage(paintedCells map[point]int, left, right, top, bottom int, filepath string) {
	w := right - left + 1
	h := top - bottom + 1
	img := image.NewRGBA(image.Rectangle{image.ZP, image.Point{w, h}})
	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			p := point{i + left, j + bottom}
			if c, ok := paintedCells[p]; ok && c == cWhite {
				img.Set(i, j, color.RGBA{255, 255, 255, 255})
			} else {
				img.Set(i, j, color.RGBA{0, 0, 0, 255})
			}
		}
	}
	f, err := os.Create(filepath)
	if err != nil {
		panic(err)
	}
	png.Encode(f, img)
}

func part2() string {
	program := intcode.ReadProgram("input.txt")
	paintedCells := make(map[point]int)
	paintedCells[point{0, 0}] = cWhite
	paint(paintedCells, program)
	left, right, top, bottom := getBoundaries(paintedCells)
	filename := "registration_identifier.png"
	renderImage(paintedCells, left, right, top, bottom, filename)
	return filename
}

func main() {
	fmt.Println("Part 1 Answer:", part1())
	fmt.Println("Part 2 Answer Image Filepath:", part2())
}
