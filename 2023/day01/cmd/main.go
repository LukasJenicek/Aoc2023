package main

import (
	_ "embed"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

//go:embed first.txt
var input string

var validNumbers = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

func main() {
	re := regexp.MustCompile("([0-9]+|one|two|three|four|five|six|seven|eight|nine)")
	result := 0
	for _, line := range strings.Split(input, "\n") {
		allNumbers := []string{}
		buffer := ""
		for _, n := range line {
			if n >= 48 && n <= 57 {
				buffer = ""
				allNumbers = append(allNumbers, fmt.Sprintf("%c", n))
				continue
			}

			if n >= 65 && n <= 122 {
				buffer += fmt.Sprintf("%c", n)
			}

			// number has at least 3 characters
			if len(buffer) >= 3 {
				find := re.FindString(buffer)

				if find != "" {
					val, _ := validNumbers[find]
					allNumbers = append(allNumbers, val)
					buffer = buffer[len(buffer)-1:]
				}
			}
		}

		n, err := strconv.Atoi(fmt.Sprintf("%s%s", allNumbers[0], allNumbers[len(allNumbers)-1]))
		if err != nil {
			log.Fatal(err)
		}

		result += n
	}

	log.Printf("result %v", result)
}
