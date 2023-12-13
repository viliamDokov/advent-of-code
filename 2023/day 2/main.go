package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Game struct {
	id     int
	reds   []int
	blues  []int
	greens []int
}

func check(err error) {
	if err != nil {
		log.Fatalf("Fatal error %s", err)
	}
}

func parseColor(slc string) (int, string) {
	tmp := strings.Split(slc, " ")
	num, err := strconv.Atoi(tmp[0])
	check(err)

	return num, tmp[1]
}

var R_MAX = 12
var GR_MAX = 13
var BL_MAX = 14

func solve(games []Game) int {
	res := 0
	for _, game := range games {
		mg := slices.Max(game.greens)
		mr := slices.Max(game.reds)
		mb := slices.Max(game.blues)

		score := mr * mg * mb

		res += score
	}
	return res
}

func read() []string {
	f, err := os.Open("inp")
	check(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	lines := make([]string, 0)
	// optionally, resize scanner's capacity for lines over 64K, see next example

	games := make([]Game, 0)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	return lines
}

func main() {
	lines := read()
	fmt.Println(lines)
}
