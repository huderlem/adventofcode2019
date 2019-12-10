package main

import (
	"fmt"
	"math"
	"sort"

	"github.com/huderlem/adventofcode2019/util"
)

type point struct {
	x, y int
}

type asteroidPoint struct {
	point
	distance float64
}

func readAsteroids() []point {
	lines := util.ReadFileLines("input.txt")
	points := []point{}
	y := 0
	for _, line := range lines {
		for x := 0; x < len(line); x++ {
			c := line[x]
			if c == '#' {
				points = append(points, point{x, y})
			}
		}
		y++
	}
	return points
}

type byDistance []asteroidPoint

func (a byDistance) Len() int           { return len(a) }
func (a byDistance) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byDistance) Less(i, j int) bool { return a[i].distance < a[j].distance }

func getAsteroidVisibilities(base point, asteroidField []point) map[float64][]asteroidPoint {
	visibilities := make(map[float64][]asteroidPoint)
	for _, asteroid := range asteroidField {
		if base == asteroid {
			continue
		}

		angle := math.Atan2(float64(asteroid.y-base.y), float64(asteroid.x-base.x))
		// Adjust angle so that a directly vertical angle is deemed the smallest angle.
		// This will allow us to conveniently rotate the laser for question part 2.
		if angle < -math.Pi/2 {
			angle += math.Pi * 2
		}
		if _, ok := visibilities[angle]; !ok {
			visibilities[angle] = []asteroidPoint{}
		}
		xf := float64(asteroid.x - base.x)
		yf := float64(asteroid.y - base.y)
		distance := math.Sqrt(xf*xf + yf*yf)
		visibilities[angle] = append(visibilities[angle], asteroidPoint{
			point:    asteroid,
			distance: distance,
		})
	}
	// Sort the asteroids in each line-of-sight from closest to furthest.
	for angle := range visibilities {
		sort.Sort(byDistance(visibilities[angle]))
	}
	return visibilities
}

func evaluateAsteroids(asteroids []point) (point, map[float64][]asteroidPoint) {
	bestAsteroid := asteroids[0]
	bestVisibilites := make(map[float64][]asteroidPoint)
	for _, asteroid := range asteroids {
		visibilites := getAsteroidVisibilities(asteroid, asteroids)
		numVisible := len(visibilites)
		if numVisible > len(bestVisibilites) {
			bestAsteroid = asteroid
			bestVisibilites = visibilites
		}
	}
	return bestAsteroid, bestVisibilites
}

func part1() int {
	asteroids := readAsteroids()
	_, visibilities := evaluateAsteroids(asteroids)
	return len(visibilities)
}

func part2() int {
	asteroids := readAsteroids()
	_, visibilities := evaluateAsteroids(asteroids)
	angles := []float64{}
	for angle := range visibilities {
		angles = append(angles, angle)
	}
	sort.Float64s(angles)
	numDestroyed := 0
	var targetAsteroid asteroidPoint
	for _, angle := range angles {
		if len(visibilities[angle]) > 0 {
			numDestroyed++
			if numDestroyed != 200 {
				visibilities[angle] = visibilities[angle][1:]
			} else {
				targetAsteroid = visibilities[angle][0]
				break
			}
		}
	}
	return targetAsteroid.x*100 + targetAsteroid.y
}

func main() {
	fmt.Println("Part 1 Answer:", part1())
	fmt.Println("Part 2 Answer:", part2())
}
