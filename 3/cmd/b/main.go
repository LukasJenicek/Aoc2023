package main

import (
	_ "embed"
	"fmt"
	"log"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	engineSchema := [][]int32{}
	rowsNumber := 0

	for _, line := range strings.Split(input, "\n") {
		m := []int32{}
		for _, char := range line {
			m = append(m, char)
		}
		engineSchema = append(engineSchema, m)
		rowsNumber++
	}

	result := 0

	for rowIndex, row := range engineSchema {
		for colIndex, _ := range row {
			cellValue := engineSchema[rowIndex][colIndex]

			// special character is star
			if cellValue == 42 {
				numbers := loadAdjacentNumbers(engineSchema, rowsNumber, rowIndex, colIndex)

				// exactly two part numbers
				if len(numbers) == 2 {
					result += numbers[0] * numbers[1]
				}
			}
		}
	}

	fmt.Printf("%d", result)
}

func loadAdjacentNumbers(engineSchema [][]int32, maxRows, row, col int) []int {
	numbers := []int{}

	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			curRow := row + i
			curCol := col + j
			val := engineSchema[curRow][curCol]
			number := ""

			if isNumber(val) {
				// scan left
				var start int
				if isNumber(engineSchema[curRow][curCol-1]) {
					c := curCol
					for isNumber(engineSchema[curRow][c]) {
						start = c
						c--

						if c == -1 {
							break
						}
					}
				} else {
					start = curCol
				}

				for isNumber(engineSchema[curRow][start]) {
					number += fmt.Sprintf("%c", engineSchema[curRow][start])
					engineSchema[curRow][start] = 46
					start++

					if start == maxRows {
						break
					}
				}

				n, err := strconv.Atoi(number)
				if err != nil {
					log.Fatal(err)
				}

				numbers = append(numbers, n)
			}
		}
	}

	return numbers
}

func isNumber(value int32) bool {
	return value >= 48 && value <= 57
}
