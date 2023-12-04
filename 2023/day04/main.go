package main

import (
	_ "embed"
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"
)

//go:embed example.txt
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
	game := loadData(strings.Split(input, "\n"))

	result := 0
	for _, card := range game.Cards {
		perGame := 0
		for _, winNumber := range card.WinningNumbers {
			for _, ownNumber := range card.OwnNumbers {
				// since we have the numbers sorted we know that once the right part has bigger numbers we don't have to continue
				if ownNumber > winNumber {
					break
				}

				if winNumber == ownNumber {
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

func loadData(lines []string) *ScratchCardGame {
	game := &ScratchCardGame{make([]*Card, len(lines))}

	for i, line := range lines {
		parts := strings.Split(line, "|")
		winingPart := strings.Split(parts[0], ":")

		winingNumbers := []int{}
		number := ""
		for _, ch := range winingPart[1] {
			if isDigit(ch) {
				number += fmt.Sprintf("%c", ch)
				continue
			}

			if ch == ' ' && number != "" {
				winNumber, err := strconv.Atoi(number)
				if err != nil {
					log.Fatal(err)
				}
				winingNumbers = append(winingNumbers, winNumber)
				number = ""
			}
		}

		ownNumbers := []int{}
		number = ""
		for _, ch := range parts[1] {
			if isDigit(ch) {
				number += fmt.Sprintf("%c", ch)
				continue
			}

			if ch == ' ' && number != "" {
				ownNumber, err := strconv.Atoi(number)
				if err != nil {
					log.Fatal(err)
				}
				ownNumbers = append(ownNumbers, ownNumber)
				number = ""
			}
		}

		// do not forget last number
		ownNumber, err := strconv.Atoi(number)
		if err != nil {
			log.Fatal(err)
		}
		ownNumbers = append(ownNumbers, ownNumber)

		slices.Sort(winingNumbers)
		slices.Sort(ownNumbers)

		game.Cards[i] = &Card{
			CardNumber:     i + 1,
			WinningNumbers: winingNumbers,
			OwnNumbers:     ownNumbers,
		}
	}

	return game
}

func isDigit(ch int32) bool {
	return ch >= 48 && ch <= 57
}
