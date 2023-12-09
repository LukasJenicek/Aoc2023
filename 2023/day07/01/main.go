package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed example.txt
var input string

type hand struct {
	cards string
	bid   int
}

func main() {
	hands := []hand{}

	fiveOfKind := []hand{}
	fourOfKind := []hand{}
	fullhouse := []hand{}
	threeOfKind := []hand{}
	twoPair := []hand{}
	onePair := []hand{}
	highCard := []hand{}

	for _, line := range strings.Split(input, "\n") {
		splitBySpace := strings.Split(line, " ")

		cards := splitBySpace[0]
		bid := parseLineNumbers(splitBySpace[1])[0]

		values := map[rune]int{
			'A': 0,
			'K': 0,
			'Q': 0,
			'J': 0,
			'T': 0,
			'9': 0,
			'8': 0,
			'7': 0,
			'6': 0,
			'5': 0,
			'4': 0,
			'3': 0,
			'2': 0,
		}

		for _, card := range cards {
			values[card]++
		}

		found := false
	Loop:
		for card, value := range values {
			switch value {
			case 5:
				fiveOfKind = append(fiveOfKind, hand{cards: cards, bid: bid})
				found = true
				break Loop
			case 4:
				fourOfKind = append(fourOfKind, hand{cards: cards, bid: bid})
				found = true
				break Loop
			case 3:
				found = true
				for c, v := range values {
					if v == 2 && c != card {
						fullhouse = append(fullhouse, hand{cards: cards, bid: bid})
						break Loop
					}
				}
				threeOfKind = append(threeOfKind, hand{cards: cards, bid: bid})
			case 2:
				found = true

				for c, v := range values {
					if v == 3 && c != card {
						fullhouse = append(fullhouse, hand{cards: cards, bid: bid})
						break Loop
					}

					if v == 2 && c != card {
						twoPair = append(twoPair, hand{cards: cards, bid: bid})
						found = true
						break Loop
					}
				}

				onePair = append(onePair, hand{cards: cards, bid: bid})
				break Loop
			}
		}

		if !found {
			highCard = append(highCard, hand{cards: cards, bid: bid})
		}
	}

	fmt.Printf("%v", hands)
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
