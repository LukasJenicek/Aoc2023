package main

import (
	_ "embed"
	"strings"
	"strconv"
	"fmt"
)

//go:embed example.txt
var input string

func main() {
	seeds := []int{}
	var mappings []map[int]int

	for lineNumber, line := range strings.Split(input, "\n") {
		// load seeds
		if lineNumber == 0 {
			parts := strings.Split(line, ":")

			seeds = parseLineNumbers(parts[1])
			continue
		}

		if line == "" {
			continue
		}

		// just name of mapping
		if line[0] >= 'a' && line[0] <= 'z' {
			continue
		}

		lineNumbers := parseLineNumbers(line)
		for i := 0; i < lineNumbers[2]; i++ {
			mappings = append(mappings, map[int]int{
				lineNumbers[1] + i: lineNumbers[0] + i,
			})
		}
	}

	fmt.Printf("%v\n", seeds)
	fmt.Printf("%v\n", mappings)
}

func parseLineNumbers(line string) []int {
	var numbers []int

	number := ""
	lineChars := len(line)
	for i, ch := range line {
		if ch >= '0' && ch <= '9' || ch == '-' {
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
