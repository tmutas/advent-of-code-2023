package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
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
	res := 0
	for i, line := range rawLines {
		tries := parseLine(line)
		valid := true
		for _, try := range tries {
			if try["red"] > 12 || try["green"] > 13 || try["blue"] > 14 {
				valid = false
			}
		}
		if valid {
			res += (i + 1)
		}
	}
	return res
}

func runPartTwo(rawLines []string) int {
	res := 0
	for _, line := range rawLines {
		tries := parseLine(line)
		var r, g, b int
		for _, try := range tries {
			r = max(r, try["red"])
			g = max(g, try["green"])
			b = max(b, try["blue"])
		}
		res += r * g * b
	}
	return res
}

func parseLine(s string) []map[string]int {
	_, games, _ := strings.Cut(s, ": ")

	tries := strings.Split(games, ";")

	var result []map[string]int

	for _, try := range tries {
		mapHere := make(map[string]int)
		colorCount := strings.Split(strings.TrimSpace(try), ",")
		for _, colorString := range colorCount {
			count, color, _ := strings.Cut(strings.TrimSpace(colorString), " ")
			mapHere[color], _ = strconv.Atoi(count)
		}
		result = append(result, mapHere)
	}
	return result
}
