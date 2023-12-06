package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"slices"
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
	cards := parseLines(rawLines)
	// for _, v := range cards {
	// 	fmt.Printf("Card: %v\n", v)
	// }
	res := 0
	for _, c := range cards {
		c.Evaluate()
		res += c.score
	}
	return res
}

func runPartTwo(rawLines []string) int {
	return 0
}

type Card struct {
	targets []int
	mynums  []int
	wins    int
	score   int
}

func (c *Card) Evaluate() {
	// fmt.Printf("%v", c.targets)
	wins := 0
	for _, mynum := range c.mynums {
		if slices.Contains(c.targets, mynum) {
			wins += 1
		}
	}
	c.wins = wins
	if c.wins >= 1 {
		c.score = int(math.Pow(2, float64(c.wins-1)))
	} else {
		c.score = 0
	}

}

func parseLines(rawLines []string) []Card {
	res := make([]Card, len(rawLines))
	for i, line := range rawLines {
		rawTargets, rawMyNums, _ := strings.Cut(line[9:], "|")
		res[i] = Card{parseNumberString(rawTargets), parseNumberString(rawMyNums), 0, 0}
	}
	return res
}

func parseNumberString(line string) []int {
	line = strings.TrimRight(line, " ")
	l := len(line)
	res := make([]int, 0)
	for i := 1; i+2 <= l; i += 3 {
		number, _ := strconv.Atoi(strings.TrimSpace(line[i : i+2]))
		res = append(res, number)
	}
	return res
}
