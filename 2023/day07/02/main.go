package main

import (
	_ "embed"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

//go:embed example.txt
var input string

type hand struct {
	cards string
	bid   int
}

var cardToPriority = map[rune]int{
	'A': 20,
	'K': 19,
	'Q': 18,
	'T': 16,
	'9': 15,
	'8': 14,
	'7': 13,
	'6': 12,
	'5': 11,
	'4': 10,
	'3': 9,
	'2': 8,
	'J': 7,
}

type hands []hand

func (h hands) Len() int      { return len(h) }
func (h hands) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h hands) Less(i, j int) bool {
	iCards := h[i].cards
	jCards := h[j].cards

	for i, c := range iCards {
		// reverse order because then I am appending all those elements to one big slice
		if cardToPriority[c] < cardToPriority[rune(jCards[i])] {
			return true
		}

		if cardToPriority[c] > cardToPriority[rune(jCards[i])] {
			return false
		}
	}

	return cardToPriority[rune(iCards[0])] > cardToPriority[rune(jCards[0])]
}

func main() {
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
			'T': 0,
			'9': 0,
			'8': 0,
			'7': 0,
			'6': 0,
			'5': 0,
			'4': 0,
			'3': 0,
			'2': 0,
			'J': 0,
		}

		for _, card := range cards {
			values[card]++
		}

		if values['J'] > 0 && values['J'] != 5 {
			count := 0
			c := 'J'
			for _, card := range cards {
				if card == 'J' {
					continue
				}

				if count == values[card] && cardToPriority[card] > cardToPriority[c] {
					c = card
					continue
				}

				if values[card] > count {
					c = card
					count = values[card]
				}
			}

			values[c] = values[c] + values['J']
			values['J'] = 0
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

	sort.Sort(hands(fiveOfKind))
	sort.Sort(hands(fourOfKind))
	sort.Sort(hands(fullhouse))
	sort.Sort(hands(threeOfKind))
	sort.Sort(hands(twoPair))
	sort.Sort(hands(onePair))
	sort.Sort(hands(highCard))

	allHands := []hand{}
	allHands = append(allHands, highCard...)
	allHands = append(allHands, onePair...)
	allHands = append(allHands, twoPair...)
	allHands = append(allHands, threeOfKind...)
	allHands = append(allHands, fullhouse...)
	allHands = append(allHands, fourOfKind...)
	allHands = append(allHands, fiveOfKind...)

	result := 0

	for i, h := range allHands {
		result += (i + 1) * h.bid
	}

	fmt.Printf("%v", result)
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
