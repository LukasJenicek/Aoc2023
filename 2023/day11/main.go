package main

import (
	_ "embed"
	"fmt"
	"math"
	"slices"
	"strings"
)

//go:embed input.txt
var input string

type Point struct {
	x int
	y int
}

type Grid struct {
	Stars     []Point
	TakenRows map[int]bool
	TakenCols map[int]bool
}

func main() {
	grid := loadGrid()

	result := sumDistances(grid, 999999)

	fmt.Printf("%v", result)
}

func sumDistances(grid Grid, expansion int) int {
	result := 0
	for i, star := range grid.Stars {
		for _, other := range grid.Stars[i+1:] {
			distance := calculateDistance(star, other)

			minX := slices.Min([]int{star.x, other.x})
			maxX := slices.Max([]int{star.x, other.x})
			minY := slices.Min([]int{star.y, other.y})
			maxY := slices.Max([]int{star.y, other.y})

			cols := 0
			rows := 0
			// cols
			for x := minX; x < maxX; x++ {
				_, ok := grid.TakenCols[x]

				if !ok {
					cols++
				}
			}

			// rows
			for x := minY; x < maxY; x++ {
				_, ok := grid.TakenRows[x]
				if !ok {
					rows++
				}
			}

			cols = cols * expansion
			rows = rows * expansion

			result += distance + rows + cols
		}
	}
	return result
}

func calculateDistance(star Point, other Point) int {
	return int(math.Abs(float64(star.x-other.x)) + math.Abs(float64(star.y-other.y)))
}

func loadGrid() Grid {
	grid := Grid{}
	grid.TakenCols = make(map[int]bool)
	grid.TakenRows = make(map[int]bool)

	for y, line := range strings.Split(input, "\n") {

		for x, ch := range line {
			if ch == '#' {
				grid.Stars = append(grid.Stars, Point{x: x, y: y})

				grid.TakenRows[y] = true
				grid.TakenCols[x] = true
			}
		}
	}

	return grid
}
