package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal("File reading error", err)
	}
	rawLines := strings.Split(strings.TrimSpace(string(data)), "\n")

	// Part 1
	part1 := 0
	part2 := 0

	for _, line := range rawLines {
		v1, v2 := getFirstLastNumber(line)
		part1 += 10*v1 + v2

		part2 += calcPartTwo(line)
		fmt.Println(calcPartTwo(line))

	}

	fmt.Println("======== PART 1 ========")
	fmt.Println(part1)

	fmt.Println("======== PART 2 ========")
	fmt.Println(part2)
}

func getFirstLastNumber(line string) (int, int) {
	numStr := strings.Map(isDigit, line)

	v1, _ := strconv.Atoi(numStr[0:1])
	v2, _ := strconv.Atoi(numStr[len(numStr)-1:])
	return v1, v2
}
func isDigit(r rune) rune {
	if unicode.IsDigit(r) {
		return r
	}
	return -1
}

func calcPartTwo(input string) int {
	numMap := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	var v1, v2 int

	for idx := 0; idx < len(input) && v1 == 0; idx++ {
		r, _ := utf8.DecodeRuneInString(input[idx : idx+1])
		if unicode.IsDigit(r) {
			v1, _ = strconv.Atoi(input[idx : idx+1])

		} else {
			for writtenNumber, numberValue := range numMap {
				if strings.HasPrefix(input[idx:], writtenNumber) {
					v1 = numberValue
				}
			}
		}
	}

	for idx := len(input) - 1; idx >= 0 && v2 == 0; idx-- {
		r, _ := utf8.DecodeRuneInString(input[idx : idx+1])
		if unicode.IsDigit(r) {
			v2, _ = strconv.Atoi(input[idx : idx+1])

		} else {
			for writtenNumber, numberValue := range numMap {
				if strings.HasPrefix(input[idx:], writtenNumber) {
					v2 = numberValue
				}
			}
		}
	}
	return 10*v1 + v2
}
