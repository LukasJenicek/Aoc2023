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

	fmt.Printf("%v\n", times)
	fmt.Printf("%v\n", distances)
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
