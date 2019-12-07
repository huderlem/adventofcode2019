package main

import (
	"fmt"
	"sync"

	"github.com/huderlem/adventofcode2019/intcode"
)

// Borrowed from https://stackoverflow.com/a/30226442
func permutations(arr []int) [][]int {
	var helper func([]int, int)
	res := [][]int{}

	helper = func(arr []int, n int) {
		if n == 1 {
			tmp := make([]int, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
}

func permutationCombine(arr1, arr2 [][]int) [][]int {
	results := [][]int{}
	for _, a := range arr1 {
		for _, b := range arr2 {
			tmp := make([]int, len(a)+len(b))
			copy(tmp, a)
			copy(tmp[len(a):], b)
			results = append(results, tmp)
		}
	}
	return results
}

func amplifierInput(phaseSetting, ampOutput int) func() int {
	inputCount := 0
	return func() int {
		switch inputCount {
		case 0:
			inputCount++
			return phaseSetting
		default:
			return ampOutput
		}
	}
}

func part1() int {
	program := intcode.ReadProgram("input.txt")
	phaseSettingsCombos := permutations([]int{0, 1, 2, 3, 4})
	maxSignal := 0
	for _, phaseSettings := range phaseSettingsCombos {
		programCopy := make([]int, len(program))
		copy(programCopy, program)
		ampOutput := 0
		for _, phaseSetting := range phaseSettings {
			intcode.ExecuteProgram(programCopy, amplifierInput(phaseSetting, ampOutput), func(output int) {
				ampOutput = output
			})
		}
		if ampOutput > maxSignal {
			maxSignal = ampOutput
		}
	}
	return maxSignal
}

func channeledAmpInput(ampID int, phaseSetting int, channels map[int]chan int) func() int {
	inputCount := 0
	return func() int {
		switch inputCount {
		case 0:
			inputCount++
			return phaseSetting
		default:
			if ampID == 0 && inputCount == 1 {
				inputCount++
				return 0
			}
			val := <-channels[ampID]
			inputCount++
			return val
		}
	}
}

func executeAmplifiers(settings []int) int {
	numAmps := len(settings)
	program := intcode.ReadProgram("input.txt")

	// Use buffered channels to pass outputs between amplifiers.
	// Buffered channels allows us to not end in a deadlock when the last amp
	// outputs its final value.
	inputChannels := make(map[int]chan int)
	for i := 0; i < numAmps; i++ {
		inputChannels[i] = make(chan int, 1)
	}

	var wg sync.WaitGroup
	output := 0
	for i := 0; i < numAmps; i++ {
		wg.Add(1)
		go func(ampID int) {
			programCopy := make([]int, len(program))
			copy(programCopy, program)
			lastOutput := 0
			intcode.ExecuteProgram(programCopy, channeledAmpInput(ampID, settings[ampID], inputChannels), func(output int) {
				lastOutput = output
				chID := (ampID + 1) % numAmps
				inputChannels[chID] <- output
			})
			if ampID == numAmps-1 {
				output = lastOutput
			}
			wg.Done()
		}(i)
	}

	wg.Wait()

	return output
}

func part2() int {
	phaseSettingCombos := permutations([]int{5, 6, 7, 8, 9})
	maxSignal := 0
	for _, settings := range phaseSettingCombos {
		signal := executeAmplifiers(settings)
		if signal > maxSignal {
			maxSignal = signal
		}
	}
	return maxSignal
}

func main() {
	fmt.Println("Part 1 Answer:", part1())
	fmt.Println("Part 2 Answer:", part2())
}
