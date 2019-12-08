package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	_ "image/png"
	"os"
	"strconv"

	"github.com/huderlem/adventofcode2019/util"
)

func readLayers(width, height int) [][]int {
	rawData := util.ReadFileString("input.txt")
	numbers := make([]int, len(rawData))
	for i, num := range rawData {
		numbers[i], _ = strconv.Atoi(string(num))
	}
	numLayers := len(numbers) / (width * height)
	layers := make([][]int, numLayers)
	for i := 0; i < numLayers; i++ {
		layer := make([]int, width*height)
		for j := 0; j < width*height; j++ {
			layer[j] = numbers[i*width*height+j]
		}
		layers[i] = layer
	}

	return layers
}

func countValue(val int, layer []int) int {
	count := 0
	for i := 0; i < len(layer); i++ {
		if layer[i] == val {
			count++
		}
	}
	return count
}

func part1() int {
	w, h := 25, 6
	layers := readLayers(w, h)
	min0Count := w * h
	min0Layer := layers[0]
	for _, layer := range layers {
		count := countValue(0, layer)
		if count < min0Count {
			min0Count = count
			min0Layer = layer
		}
	}
	digits1 := countValue(1, min0Layer)
	digits2 := countValue(2, min0Layer)
	return digits1 * digits2
}

const (
	pBlack       = 0
	pWhite       = 1
	pTransparent = 2
)

func renderImage(pixels [][]int, w, h int, filepath string) {
	img := image.NewRGBA(image.Rectangle{image.ZP, image.Point{w, h}})
	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			pixel := pixels[i][j]
			switch pixel {
			case pBlack:
				img.Set(i, j, color.RGBA{0, 0, 0, 255})
			case pWhite:
				img.Set(i, j, color.RGBA{255, 255, 255, 255})
			case pTransparent:
				img.Set(i, j, color.RGBA{0, 0, 0, 0})
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
	w, h := 25, 6
	layers := readLayers(w, h)
	pixels := make([][]int, w)
	for i := 0; i < w; i++ {
		column := make([]int, h)
		for j := 0; j < h; j++ {
			pixel := pTransparent
			for _, layer := range layers {
				p := layer[j*w+i]
				if p != pTransparent {
					pixel = p
					break
				}
			}
			column[j] = pixel
		}
		pixels[i] = column
	}
	filepath := "BIOS_password.png"
	renderImage(pixels, w, h, filepath)
	return filepath
}

func main() {
	fmt.Println("Part 1 Answer:", part1())
	fmt.Println("Part 2 Answer Filepath:", part2())
}
