package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

func main() {
	flag.Parse()
	fileName := flag.Arg(0)
	switch fileName {
	case "test":
		fileName = "sample.txt"
	case "":
		fileName = "input.txt"
	}

	data, _ := os.ReadFile(fileName)
	rawLines := strings.Split(strings.TrimSpace(string(data)), "\n")

	fmt.Println("======== PART 1 ========")
	fmt.Println(runPartOne(rawLines))

	fmt.Println("======== PART 2 ========")
	fmt.Println(runPartTwo(rawLines))
}

func runPartOne(rawLines []string) int {
	grid := parseLines(rawLines)
	res := 0
	for i, line := range rawLines {
		inds := findNumberIndices(line)
		for _, ind := range inds {
			//fmt.Printf("%d %s, ", readNumber(line, ind), checkSymbols(grid, ind, i))
			if checkSymbols(grid, ind, i) {
				res += readNumber(line, ind)
				fmt.Printf("%d ", readNumber(line, ind))
			}
		}
		println()
	}
	return res
}

func runPartTwo(rawLines []string) int {
	return 0
}

func parseLines(rawLines []string) [][]rune {
	grid := make([][]rune, len(rawLines))
	for i, line := range rawLines {
		gridLine := make([]rune, len(line))
		for j, char := range line {
			gridLine[j] = char
		}
		grid[i] = gridLine
	}
	return grid
}

func readNumber(line string, ind Indices) int {
	v, _ := strconv.Atoi(line[ind.start:ind.end])
	return v
}

func validSymbol(r rune) bool {
	dot, _ := utf8.DecodeRuneInString(".")
	//return !unicode.IsNumber(r) && r != dot
	return r != dot
}

func checkSymbols(grid [][]rune, ind Indices, lineIndex int) bool {
	maxIndex := len(grid[lineIndex])
	// Check line above and below
	for i := max(0, ind.start-1); i < min(maxIndex, ind.end+1); i++ {
		if lineIndex-1 >= 0 {
			if validSymbol(grid[lineIndex-1][i]) {
				return true
			}
		}
		if lineIndex+1 < maxIndex {
			if validSymbol(grid[lineIndex+1][i]) {
				return true
			}

		}
	}

	if ind.start-1 >= 0 {
		if validSymbol(grid[lineIndex][ind.start-1]) {
			return true
		}
	}
	if ind.end < maxIndex {
		if validSymbol(grid[lineIndex][ind.end]) {
			return true
		}

	}

	return false
}

type Indices struct {
	start, end int
}

func findNumberIndices(line string) []Indices {
	curStart := -1
	numbers := make([]Indices, 0)
	for i, r := range line {
		if unicode.IsDigit(r) {
			if curStart < 0 {
				curStart = i
			}

		} else {
			if curStart >= 0 {
				numbers = append(numbers, Indices{curStart, i})
				curStart = -1
			}
		}
	}
	return numbers
}
