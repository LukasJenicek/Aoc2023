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
		lastIndex := len(scheme[j]) - 1
		value += scheme[j][lastIndex]
	}

	return value
}

func buildScheme(scheme [][]int, row []int) [][]int {
	newRow := []int{}
	maxIndex := len(row) - 1
	rowValue := 0

	for i, val := range row {
		if i == maxIndex && rowValue == 0 {
			scheme = append(scheme, newRow)
			break
		}

		if i+1 > maxIndex && rowValue > 0 {
			scheme = append(scheme, newRow)
			scheme = buildScheme(scheme, newRow)
			break
		}

		nextVal := row[i+1] - val
		newRow = append(newRow, nextVal)

		rowValue += nextVal
	}

	return scheme
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
