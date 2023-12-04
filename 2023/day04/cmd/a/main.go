package main

import (
	_ "embed"
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"
)

//go:embed ../../input.txt
var input string

type ScratchCardGame struct {
	Cards []*Card
}

type Card struct {
	CardNumber     int
	WinningNumbers []int
	OwnNumbers     []int
}

func main() {
	lines := strings.Split(input, "\n")
	colonIndex := strings.Index(lines[0], ":")

	result := 0

	game := &ScratchCardGame{make([]*Card, len(lines)-1)}

	for i, line := range lines {
		game.Cards[i] = &Card{
			CardNumber:     i + 1,
			WinningNumbers: []int{},
			OwnNumbers:     []int{},
		}

		var cards []int
		number := ""
		for _, ch := range line[colonIndex+2:] {
			if isDigit(ch) {
				number += fmt.Sprintf("%c", ch)
			}

			// space
			if ch == 32 && number != "" {
				n, err := strconv.Atoi(number)
				if err != nil {
					log.Fatal(err)
				}

				cards = append(cards, n)
				number = ""
			}

			// |
			if ch == 124 {
				slices.Sort(cards)

				game.Cards[i].WinningNumbers = cards
				cards = []int{}
			}
		}

		n, err := strconv.Atoi(number)
		if err != nil {
			log.Fatal(err)
		}

		cards = append(cards, n)
		number = ""

		slices.Sort(cards)
		// second part of the cards
		game.Cards[i].OwnNumbers = cards
	}

	for _, card := range game.Cards {
		perGame := 0
		for _, first := range card.WinningNumbers {
			for _, second := range card.OwnNumbers {
				if second > first {
					break
				}

				if first == second {
					if perGame == 0 {
						perGame = 1
					} else {
						perGame *= 2
					}
					break
				}
			}
		}

		result += perGame
	}

	fmt.Printf("%d", result)
}

func isDigit(ch int32) bool {
	return ch >= 48 && ch <= 57
}
