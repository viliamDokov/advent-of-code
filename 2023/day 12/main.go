package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func check(err error) {
	if err != nil {
		log.Fatalf("Fatal error %s", err)
	}
}

const WORK = "."
const UNKNOWN = "?"
const BROKEN = "#"

func read(fname string) ([][]string, [][]int) {
	f, err := os.ReadFile(fname)
	data := string(f)
	check(err)

	lines := strings.Split(data, "\n")

	schematics := make([][]string, 0)
	orders := make([][]int, 0)

	for _, l := range lines {
		fmt.Println(l)
		l = strings.TrimSpace(l)
		tmp := strings.Split(l, " ")
		schematic := make([]string, 0)
		for i := range tmp[0] {
			schematic = append(schematic, tmp[0][i:i+1])
		}

		ns := strings.Split(tmp[1], ",")
		order := make([]int, 0)
		for _, n := range ns {
			n, err := strconv.Atoi(n)
			check(err)
			order = append(order, n)
		}

		schematics = append(schematics, schematic)
		orders = append(orders, order)
	}
	return schematics, orders
}

func solve(schematics [][]string, orders [][]int) int {
	acc := 0
	for i := range schematics {
		schematics[i], orders[i] = multiplySchematic(schematics[i], orders[i])
		cache := make(map[State]int)
		r := recurse(schematics[i], orders[i], State{0, 0}, cache)
		fmt.Println("result", r)
		acc += r
	}
	return acc
}

func multiplySchematic(schematic []string, order []int) ([]string, []int) {
	sr := make([]string, 0)
	or := make([]int, 0)
	for i := 0; i < 5; i++ {
		sr = append(sr, schematic...)
		if i != 4 {
			sr = append(sr, UNKNOWN)
		}
		or = append(or, order...)
	}
	return sr, or
}

type State struct {
	si int
	oi int
}

func recurse(schematic2 []string, order2s []int, s State, cache map[State]int) int {
	v, in := cache[s]
	if in {
		// fmt.Println("HIT", cache, s)
		return v
	}

	order := order2s[s.oi:]
	schematic := schematic2[s.si:]

	if len(order) == 0 {
		// fmt.Println("left broken", schematic, schematic2, s)

		for i := range schematic {

			if schematic[i] == BROKEN {
				cache[s] = 0
				// fmt.Println("left broken")
				return 0
			}
		}
		// fmt.Println("return", 1)
		cache[s] = 1
		return 1
	}

	if len(schematic) == 0 && len(order) > 0 {
		cache[s] = 0
		// fmt.Println("MORE")
		return 0
	}

	n := order[0]
	ways := 0
	for i := 0; i <= len(schematic)-n; i++ {
		cont := false
		for j := 0; j < n; j++ {
			if schematic[i+j] == WORK {
				cont = true
				break
			}
		}

		if cont {
			continue
		}

		br := false
		for _, t := range schematic[:i] {
			if t == BROKEN {
				br = true
				break
			}
		}
		if br {
			break
		}

		if i+n == len(schematic) {
			// fmt.Println("MATCH", schematic, n, i, schematic[:i])
			// r := recurse(schematic[i+n:], order[1:])
			r := recurse(schematic2, order2s, State{s.si + i + n, s.oi + 1}, cache)
			ways += r
			continue
		}

		if schematic[i+n] == WORK || schematic[i+n] == UNKNOWN {
			// r := recurse(schematic[i+n+1:], order[1:])
			// fmt.Println("MATCH", schematic, n, i, schematic[:i])
			r := recurse(schematic2, order2s, State{s.si + i + n + 1, s.oi + 1}, cache)
			ways += r
		}
	}

	cache[s] = ways
	return ways
}

func run(fname string) int {
	schematics, orders := read(fname)
	fmt.Println(schematics, orders)
	r := solve(schematics, orders)
	fmt.Println(r)
	return r
}
func main() {
	expected := 525152
	rtest := run("inp.test")
	if rtest != expected {
		log.Fatalln("FAILED TEST CASE", rtest, expected)
	}
	rmain := run("inp.main")
	fmt.Println("Result", rmain)
}
