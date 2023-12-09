package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed example.txt
var input string

func main() {
	lines := strings.Split(input, "\n")

	times := parseLineNumbers(lines[0])
	distances := parseLineNumbers(lines[1])

	finalCombinations := []int{}

	for timeIndex, time := range times {
		speed := 0
		combinations := []int{}
		distance := distances[timeIndex]
		deltaTime := time - speed

		for deltaTime > 0 {
			if distance < speed*deltaTime {
				combinations = append(combinations, 1)
			}

			speed++
			deltaTime = time - speed
		}

		finalCombinations = append(finalCombinations, len(combinations))
	}

	result := 0

	for _, combination := range finalCombinations {
		if result == 0 {
			result = combination
			continue
		}

		result *= combination
	}

	fmt.Printf("%v\n", result)
}

func parseLineNumbers(line string) []int {
	var numbers []int

	number := ""
	lineChars := len(line)
	for i, ch := range line {
		if ch >= '0' && ch <= '9' {
			number += fmt.Sprintf("%c", ch)

			if i != lineChars-1 {
				continue
			}
		}

		// either space or last char tells me that I need to parse number
		if (ch == ' ' && number != "") || (i == lineChars-1 && number != "") {
			n, _ := strconv.Atoi(number)
			numbers = append(numbers, n)
			number = ""
		}
	}

	return numbers
}
