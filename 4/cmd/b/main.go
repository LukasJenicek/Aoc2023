package main

import (
	_ "embed"
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type Game struct {
	Cards []*Card
}

type Card struct {
	CardNumber int
	First      []int
	Second     []int
}

func main() {
	lines := strings.Split(input, "\n")
	colonIndex := strings.Index(lines[0], ":")

	game := &Game{make([]*Card, len(lines)-1)}

	for i, line := range lines {
		// skip empty lines
		if line == "" {
			continue
		}

		game.Cards[i] = &Card{
			CardNumber: i + 1,
			First:      []int{},
			Second:     []int{},
		}

		var numbers []int
		number := ""
		for _, ch := range line[colonIndex+2:] {
			if isDigit(ch) {
				number += fmt.Sprintf("%c", ch)
			}

			if ch == ' ' && number != "" {
				n, err := strconv.Atoi(number)
				if err != nil {
					log.Fatal(err)
				}

				numbers = append(numbers, n)
				number = ""
			}

			if ch == '|' {
				slices.Sort(numbers)

				game.Cards[i].First = numbers
				numbers = []int{}
			}
		}

		n, err := strconv.Atoi(number)
		if err != nil {
			log.Fatal(err)
		}

		numbers = append(numbers, n)

		slices.Sort(numbers)
		game.Cards[i].Second = numbers
	}

	matchesPerGame := map[int]int{}
	result := 0

	for _, card := range game.Cards {
		matches := 0
		for _, first := range card.First {
			for _, second := range card.Second {
				if second > first {
					break
				}

				if first == second {
					matches++
					break
				}
			}
		}

		matchesPerGame[card.CardNumber]++
		result += matchesPerGame[card.CardNumber]

		if matches > 0 {
			for i := 0; i < matchesPerGame[card.CardNumber]; i++ {
				start := card.CardNumber + 1
				for j := 0; j < matches; j++ {
					matchesPerGame[start]++

					start++
				}
			}

		}
	}

	fmt.Printf("%d", result)
}

func isDigit(ch int32) bool {
	return ch >= 48 && ch <= 57
}
