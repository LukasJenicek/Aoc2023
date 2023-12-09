package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed example02.txt
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

	for _, line := range lines[startOfMap:] {
		parts := strings.Split(line, "=")

		mapIndex := ""
		for _, ch := range parts[0] {
			if ch == ' ' {
				break
			}
			mapIndex += string(ch)
		}

		rightPart := parts[1]

		bracketStartIndex := strings.Index(rightPart, "(")
		bracketEndIndex := strings.Index(rightPart, ")")
		left := rightPart[bracketStartIndex+1 : bracketStartIndex+4]
		right := rightPart[bracketEndIndex-3 : bracketEndIndex]

		instructionsMap[mapIndex] = []string{left, right}
	}

	steps := 0
	jump := "AAA"
FOREVER:
	for {
		for _, instruction := range instructions {
			var index int

			if instruction == 'L' {
				index = 0
			} else {
				index = 1
			}

			jump = instructionsMap[jump][index]
			steps++

			if jump == "ZZZ" {
				break FOREVER
			}
		}
	}

	fmt.Printf("%v", steps)
}
