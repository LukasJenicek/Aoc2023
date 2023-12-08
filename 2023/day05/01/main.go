package main

import (
	_ "embed"
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type Island struct {
	Mapping []*Mapping
}

type Mapping struct {
	// seed-to-soil||soil-to-fertilizer ....
	Type string
	Maps []Map
}

type Map struct {
	Destination Interval
	Source      Interval
}

type Interval struct {
	Min int
	Max int
}

func main() {
	seeds := []int{}
	island := &Island{}

	var mapping *Mapping
	for i, line := range strings.Split(input, "\n") {
		// seeds
		if i == 0 {
			number := ""

			for _, ch := range line {
				if isDigit(ch) {
					number += fmt.Sprintf("%c", ch)
					continue
				}

				if ch == ' ' && number != "" {
					seed, err := strconv.Atoi(number)
					if err != nil {
						log.Fatal(err)
					}
					seeds = append(seeds, seed)
					number = ""
				}
			}

			seed, _ := strconv.Atoi(number)
			seeds = append(seeds, seed)

			continue
		}

		if line == "" {
			if mapping != nil {
				island.Mapping = append(island.Mapping, mapping)
			}

			mapping = &Mapping{}
			continue
		}

		if line[0] >= 'a' && line[0] <= 'z' {
			split := strings.Split(line, " ")

			mapping.Type = split[0]
			continue
		}

		if line[0] >= '0' && line[0] <= '9' {
			number := ""
			destInterval := -1
			sourceInterval := -1
			lastChar := len(line) - 1

			for i, ch := range line {
				if isDigit(ch) {
					number += fmt.Sprintf("%c", ch)

					if i == lastChar {
						add, _ := strconv.Atoi(number)
						add -= 1

						mapping.Maps = append(mapping.Maps, Map{
							Destination: Interval{Min: destInterval, Max: destInterval + add},
							Source:      Interval{Min: sourceInterval, Max: sourceInterval + add},
						})
					}

					continue
				}

				if ch == ' ' && number != "" {
					if destInterval == -1 {
						destInterval, _ = strconv.Atoi(number)
						number = ""
						continue
					}

					if sourceInterval == -1 {
						sourceInterval, _ = strconv.Atoi(number)
					}

					number = ""
				}
			}
		}
	}

	island.Mapping = append(island.Mapping, mapping)

	finalSeedLocations := []int{}

	for _, seed := range seeds {
		locations := []int{}
		destination := seed
		for _, mapping := range island.Mapping {
			for _, m := range mapping.Maps {
				if destination >= m.Source.Min && destination <= m.Source.Max {
					destination = destination + (m.Destination.Min - m.Source.Min)

					break
				}
			}

			locations = append(locations, destination)
		}

		finalSeedLocations = append(finalSeedLocations, locations[len(locations)-1])
	}

	sort.Ints(finalSeedLocations)

	fmt.Printf("%v\n", finalSeedLocations)
	fmt.Printf("result: %v", finalSeedLocations[0])
}

func isDigit(ch int32) bool {
	return ch >= 48 && ch <= 57
}
