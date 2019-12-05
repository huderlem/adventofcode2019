package main

import (
	"fmt"
	"math"
)

func digit(val, n int) int {
	div := int(math.Pow10(n))
	return (val / div) % 10
}

func hasAdjacentDigits(digits []int) bool {
	for i := 1; i < 6; i++ {
		if digits[i] == digits[i-1] {
			return true
		}
	}
	return false
}

func hasAdjacentDigitsWithLimit(digits []int, limit int) bool {
	i := 0
	for i < 6 {
		streak := 0
		curVal := digits[i]
		for i < 6 && digits[i] == curVal {
			streak++
			i++
		}
		if streak == limit {
			return true
		}
	}
	return false
}

func isMonotonicIncrease(digits []int) bool {
	for i := 1; i < 6; i++ {
		if digits[i] < digits[i-1] {
			return false
		}
	}
	return true
}

func part1() int {
	count := 0
	digits := make([]int, 6)
	for val := 206938; val <= 679128; val++ {
		for i := 0; i < 6; i++ {
			digits[i] = digit(val, 5-i)
		}
		if hasAdjacentDigits(digits) && isMonotonicIncrease(digits) {
			count++
		}
	}
	return count
}

func part2() int {
	count := 0
	digits := make([]int, 6)
	for val := 206938; val <= 679128; val++ {
		for i := 0; i < 6; i++ {
			digits[i] = digit(val, 5-i)
		}
		if hasAdjacentDigitsWithLimit(digits, 2) && isMonotonicIncrease(digits) {
			count++
		}
	}
	return count
}

func main() {
	fmt.Println("Part 1 Answer:", part1())
	fmt.Println("Part 2 Answer:", part2())
}
