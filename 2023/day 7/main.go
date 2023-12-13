package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Hand struct {
	cards []int64
	bid   int64
	power int64
	raw   string
	htype int64
}

var CARD_MAP = map[rune]int64{
	'A': 12,
	'K': 11,
	'Q': 10,
	'J': 9,
	'T': 8,
	'9': 7,
	'8': 6,
	'7': 5,
	'6': 4,
	'5': 3,
	'4': 2,
	'3': 1,
	'2': 0,
}

var MAGIC = int64(4826809 * 14)

func thingPower(hand Hand) int64 {

	freqs := make(map[int64]int64)
	for _, c := range hand.cards {
		freqs[c]++
	}

	fr_arr := make([]int64, 0)
	for _, v := range freqs {
		fr_arr = append(fr_arr, v)
	}

	sort.Slice(fr_arr, func(i, j int) bool {
		return fr_arr[i] >= fr_arr[j]
	})

	if fr_arr[0] >= 5 {
		return MAGIC * 10
	}
	if fr_arr[0] >= 4 {
		// fmt.Println("4!!!")
		// fmt.Println(fr_arr, hand)

		return MAGIC * 9
	}
	if fr_arr[0]+fr_arr[1] >= 5 {
		// fmt.Println("3 and 2!!!")
		// fmt.Println(fr_arr, hand)

		return MAGIC * 8
	}
	if fr_arr[0] >= 3 {
		// fmt.Println("3!!!")
		// fmt.Println(fr_arr, hand)

		return MAGIC * 7
	}
	if fr_arr[0]+fr_arr[1] >= 4 {
		// fmt.Println("2 and 2!!!")
		// fmt.Println(fr_arr, hand)
		return MAGIC * 6
	}
	if fr_arr[0] >= 2 {
		// fmt.Println("2!!!")
		// fmt.Println(fr_arr, hand)
		return MAGIC * 5
	}
	if fr_arr[0] >= 1 {
		return MAGIC * 4
	}

	log.Fatalf("Never happen")
	return 0
}

func orderPower(hand Hand) int64 {
	s := int64(0)
	for i, _ := range hand.cards {
		r := len(hand.cards) - 1 - i
		p := int64(math.Pow(14, float64(i)))
		s += p * hand.cards[r]
	}
	return s
}


func compare2(a Hand, b Hand) bool {
	return a.power > b.power
}

func seeSolutions(hands []Hand) {
	for i := range hands {
		hands[i].htype = thingPower(hands[i])
		hands[i].power = hands[i].htype + orderPower(hands[i])
	}

	for _, a := range hands {
		for _, b := range hands {
			if compare1(a, b) != compare2(a, b) && a.power != b.power {
				fmt.Println("DIFF", a, b, compare1(a, b), compare2(a, b))
			}
		}
	}
}

func solve(hands []Hand) int {
	for i := range hands {
		hands[i].htype = thingPower(hands[i])
		hands[i].power = hands[i].htype + orderPower(hands[i])
	}

	sort.Slice(hands, func(i int, j int) bool { return compare2(hands[i], hands[j]) })
	fmt.Println(hands)
	s := 0
	for i, h := range hands {
		fmt.Println((i + 1) * int(h.bid))
		s += (i + 1) * int(h.bid)
	}
	return s
}

func check(err error) {
	if err != nil {
		log.Fatalf("Fatal error %s", err)
	}
}

func read() []Hand {
	f, err := os.ReadFile("inp")
	data := string(f)
	check(err)
	lines := strings.Split(data, "\n")
	cards := make([]Hand, 0)
	for _, l := range lines {
		tmp := strings.Split(l, " ")
		card := tmp[0]
		values := make([]int64, 0)
		bid, err := strconv.ParseInt(strings.TrimSpace(tmp[1]), 10, 64)
		check(err)
		for _, c := range card {
			values = append(values, CARD_MAP[c])
		}
		cards = append(cards, Hand{values, bid, 0, tmp[0], 0})
	}

	return cards
}

func main() {
	cs := read()
	// fmt.Println(cs)
	seeSolutions(cs)
	res := solve(cs)
	fmt.Println(res)
}
