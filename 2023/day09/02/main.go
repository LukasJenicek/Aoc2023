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
	history := [][]int{}

	for _, line := range strings.Split(input, "\n") {
		row := parseLineNumbers(line)
		history = append(history, row)
	}

	historyValue := 0
	for _, row := range history {
		scheme := [][]int{}
		scheme = append(scheme, row)

		historyValue += calculateHistoryValue(scheme)
	}

	fmt.Printf("%v", historyValue)
}

func calculateHistoryValue(scheme [][]int) int {
	scheme = buildScheme(scheme, scheme[0])
	rowNumbers := len(scheme) - 1

	value := 0
	for j := rowNumbers; j >= 0; j-- {
		// that's the end
		if j == 0 {
			break
		}

		// first iteration
		if j == rowNumbers {
			value = scheme[j-1][0] - scheme[j][0]
		} else {
			value = scheme[j-1][0] - value
		}
	}

	return value
}

func buildScheme(scheme [][]int, row []int) [][]int {
	newRow := []int{}
	maxIndex := len(row) - 1

	for i, val := range row {
		if i == maxIndex && onlyZeros(newRow) {
			scheme = append(scheme, newRow)
			break
		}

		if i+1 > maxIndex {
			scheme = append(scheme, newRow)
			scheme = buildScheme(scheme, newRow)
			break
		}

		nextVal := row[i+1] - val
		newRow = append(newRow, nextVal)
	}

	return scheme
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

func onlyZeros(row []int) bool {
	for _, num := range row {
		if num != 0 {
			return false
		}
	}

	return true
}
