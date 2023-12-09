package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	instructions := ""
	lines := strings.Split(input, "\n")

	// read instructions
	startOfMap := 0
	for i, line := range lines {
		if line == "" {
			startOfMap = i + 1
			break
		}

		instructions += line
	}

	instructionsMap := map[string][]string{}

	jumpIndexes := []string{}

	for _, line := range lines[startOfMap:] {
		parts := strings.Split(line, "=")

		mapIndex := ""
		for _, ch := range parts[0] {
			if ch == ' ' {
				break
			}
			mapIndex += string(ch)
		}
		if mapIndex[len(mapIndex)-1] == 'A' {
			jumpIndexes = append(jumpIndexes, mapIndex)
		}

		rightPart := parts[1]

		bracketStartIndex := strings.Index(rightPart, "(")
		bracketEndIndex := strings.Index(rightPart, ")")
		left := rightPart[bracketStartIndex+1 : bracketStartIndex+4]
		right := rightPart[bracketEndIndex-3 : bracketEndIndex]

		instructionsMap[mapIndex] = []string{left, right}
	}

	n := []int{}

	steps := 0
FOREVER:
	for {
		for _, instruction := range instructions {
			var direction int

			if instruction == 'L' {
				direction = 0
			} else {
				direction = 1
			}

			steps++

			del := -1
			// brute force is not gonna work on the real input :(
			for i, jumpIndex := range jumpIndexes {
				val := instructionsMap[jumpIndex][direction]

				if val[len(val)-1] == 'Z' {
					n = append(n, steps)
					del = i
				}

				jumpIndexes[i] = val
			}

			if del != -1 {
				jumpIndexes = append(jumpIndexes[:del], jumpIndexes[del+1:]...)
			}

			if len(jumpIndexes) == 0 {
				break FOREVER
			}
		}
	}

	fmt.Printf("%v", LCM(n[0], n[1], n[2:]...))
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
