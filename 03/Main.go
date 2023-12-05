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
		inds := findNumberIndices(line, i)
		for _, ind := range inds {
			//fmt.Printf("%d %s, ", readNumber(line, ind), checkSymbols(grid, ind, i))
			if checkSymbols(grid, ind) {
				res += ind.value
				// fmt.Printf("%d ", readNumber(line, ind))
			}
		}
		// println()
	}
	return res
}

func runPartTwo(rawLines []string) int {
	grid := parseLines(rawLines)
	var numbers []Indices
	for i, line := range rawLines {
		numbers = append(numbers, findNumberIndices(line, i)...)
	}

	res := 0
	for i, gridline := range grid {
		for j, r := range gridline {
			if dot, _ := utf8.DecodeRuneInString("*"); r == dot {
				res += findAdjecentNumbers(numbers, i, j)
			}
		}
	}
	return res
}

func findAdjecentNumbers(numbers []Indices, i int, j int) int {
	vals := make([]Indices, 0)
	for _, num := range numbers {
		if (num.line-1 <= i && i <= num.line+1) && (num.start-1 <= j && j <= num.end) {
			vals = append(vals, num)
		}
	}

	if len(vals) == 2 {
		res := vals[0].value * vals[1].value
		return res
	}
	fmt.Println()
	return 0
}

func parseLines(rawLines []string) [][]rune {
	grid := make([][]rune, len(rawLines)+1)
	for i, line := range rawLines {
		gridLine := make([]rune, len(line)+1)
		for j, char := range line {
			gridLine[j] = char
		}
		grid[i] = gridLine
	}
	grid[len(rawLines)] = make([]rune, len(rawLines[0])+1)
	return grid
}

func readNumber(line string, start int, end int) int {
	v, _ := strconv.Atoi(line[start:end])
	return v
}

func validSymbol(r rune) bool {
	dot, _ := utf8.DecodeRuneInString(".")
	//return !unicode.IsNumber(r) && r != dot
	var defaultRune rune
	return r != dot && !unicode.IsNumber(r) && r != defaultRune
}

func checkSymbols(grid [][]rune, ind Indices) bool {
	// Check line above and below
	var symbols []rune
	startIndex := max(0, ind.start-1)
	if ind.line > 0 {
		symbols = append(symbols, grid[ind.line-1][startIndex:ind.end+1]...)
	}
	symbols = append(symbols, grid[ind.line+1][startIndex:ind.end+1]...)

	symbols = append(symbols, grid[ind.line][startIndex])
	symbols = append(symbols, grid[ind.line][ind.end])

	// fmt.Printf("For (%d, %d:%d): ", ind.line, ind.start, ind.end)
	// for _, r := range symbols {
	// 	fmt.Print(string(r))
	// }

	for _, r := range symbols {
		if validSymbol(r) {
			// fmt.Println(" Has Symbol!")
			return true
		}

	}
	// fmt.Println()
	return false
}

type Indices struct {
	start, end, line, value int
}

func findNumberIndices(line string, lineIndex int) []Indices {
	curStart := -1
	numbers := make([]Indices, 0)
	line = line + "."
	for i, r := range line {
		if unicode.IsDigit(r) {
			if curStart < 0 {
				curStart = i
			}

		} else {
			if curStart >= 0 {
				numbers = append(numbers, Indices{curStart, i, lineIndex, readNumber(line, curStart, i)})
				curStart = -1
			}
		}
	}

	return numbers
}
