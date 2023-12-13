package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

type Card struct {
	winning map[int]bool
	current []int
}

func check(err error) {
	if err != nil {
		log.Fatalf("Fatal error %s", err)
	}
}

func solve1(cards []Card) int {
	s := 0
	for _, card := range cards {
		total := 0
		for _, n := range card.current {
			if card.winning[n] {
				total++
			}
		}
		s += int(math.Pow(2, float64(total-1)))
	}
	return s
}

func (card *Card) getWinningCount() int {
	total := 0
	for _, n := range card.current {
		if card.winning[n] {
			total++
		}
	}
	return total
}

func recurse(cards []Card, i int, cache map[int]int) int {
	if cache[i] > 0 {
		return cache[i]
	}

	t := cards[i].getWinningCount()
	s := 1

	for j := i + 1; j <= i+t; j++ {
		// fmt.Println(i, j)
		if j < len(cards) {
			s += recurse(cards, j, cache)
		}
	}

	cache[i] = s
	return s
}

func solve2(cards []Card) int {
	s := 0
	cache := make(map[int]int)
	for i := range cards {
		s += recurse(cards, i, cache)
	}
	return s
}

func read() []Card {
	f, err := os.Open("inp")
	check(err)
	defer f.Close()

	cards := make([]Card, 0)
	scanner := bufio.NewScanner(f)
	// optionally, resize scanner's capacity for lines over 64K, see next example

	for scanner.Scan() {
		line := scanner.Text()
		tmp := strings.Split(line, ": ")
		numbers := strings.Split(tmp[1], " | ")
		winning_str := strings.Split(numbers[0], " ")
		winning := make(map[int]bool)
		for _, wns := range winning_str {
			if wns == "" {
				continue
			}
			wn, err := strconv.Atoi(strings.TrimSpace(wns))
			check(err)
			winning[wn] = true
		}
		curr_str := strings.Split(numbers[1], " ")
		curr := make([]int, 0)

		for _, cns := range curr_str {
			if cns == "" {
				continue
			}
			cn, err := strconv.Atoi(strings.TrimSpace(cns))
			check(err)
			curr = append(curr, cn)
		}
		cards = append(cards, Card{winning: winning, current: curr})
	}
	return cards
}

func main() {
	start := time.Now()
	cards := read()
	fmt.Println(cards)
	// s := solve1(cards)
	// fmt.Println(s)
	s := solve2(cards)
	fmt.Println(s)
	elapsed := time.Since(start)
	log.Printf("Took %s\n", elapsed)
}
