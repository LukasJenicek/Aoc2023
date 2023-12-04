package main

import (
	_ "embed"
	"fmt"
	"log"
	"strconv"
	"strings"
)

//go:embed first.txt
var input string

var bag = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func main() {
	result := 0

	for _, line := range strings.Split(input, "\n") {
		splitByColon := strings.Split(line, ":")
		sets := strings.Split(splitByColon[1], ";")

		minimumNumberOfStones := map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}
		for _, set := range sets {
			for _, game := range strings.Split(set, ",") {
				number := ""
				color := ""
				for _, ch := range game {
					if ch >= 48 && ch <= 57 {
						number += fmt.Sprintf("%c", ch)
						continue
					}

					if ch >= 65 && ch <= 122 {
						color += fmt.Sprintf("%c", ch)
					}
				}
				count, err := strconv.Atoi(number)
				if err != nil {
					log.Fatal(err)
				}

				if count > minimumNumberOfStones[color] {
					minimumNumberOfStones[color] = count
				}
			}
		}

		result += minimumNumberOfStones["red"] * minimumNumberOfStones["green"] * minimumNumberOfStones["blue"]
	}
	fmt.Printf("%d\n", result)
}
