package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func check(err error) {
	if err != nil {
		log.Fatalf("Fatal error %s", err)
	}
}

type Pos struct {
	row int64
	col int64
}

func read(fname string) ([][]string, Pos) {
	f, err := os.ReadFile(fname)
	data := string(f)
	check(err)
	lines := strings.Split(data, "\n")

	gardens := make([][]string, 0)
	var sPos Pos
	for i, l := range lines {
		row := make([]string, 0)
		l = strings.TrimSpace(l)
		for j := range l {
			if l[j:j+1] == "S" {
				sPos = Pos{int64(i), int64(j)}
			}
			row = append(row, l[j:j+1])
		}
		gardens = append(gardens, row)
	}
	return gardens, sPos
}

type Step struct {
	pos   Pos
	steps int
}

var NEIGHBOURS = []Pos{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

func inBounds(gardens [][]string, pos Pos) bool {
	return 0 <= pos.row && pos.row < int64(len(gardens)) && 0 <= pos.col && pos.col < int64(len(gardens[0]))
}

func Abs(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}

func mod(a int, b int) int {
	return (b + (a % b)) % b
}

func oneNorm(p Pos) int64 {
	return int64(Abs(p.row) + Abs(p.col))
}

func solve2(gardens [][]string, maxSteps int64) int64 {
	n := int64(len(gardens))
	nSq := (maxSteps - (n / 2)) / n

	fmt.Println("nSq", nSq)
	acc := int64(0)

	// UP
	corners := []Pos{{n - 1, 0}, {n - 1, n - 1}, {0, n - 1}, {0, 0}}
	mids := []Pos{{n - 1, (n / 2)}, {(n / 2), n - 1}, {0, (n / 2)}, {(n / 2), 0}}

	for i := range mids {
		c1 := int64(poop(gardens, corners[i], n+(n/2)-1))
		fmt.Println("c1", corners[i], n+(n/2)-1, c1)
		c2 := int64(poop(gardens, corners[i], (n/2)-1))
		fmt.Println("c2", corners[i], (n/2)-1, c2)
		m := int64(poop(gardens, mids[i], n-1))
		fmt.Println("m", mids[i], n-1, m)
		acc += m + (nSq)*c2 + (nSq-1)*c1
	}

	// for i := int64(0); i < nSq; i++ {
	// 	posUp1 := Pos{nSq - i, i}
	// 	posUp2 := Pos{nSq - (i), (i + 1)}
	// 	posRight1 := Pos{-i, nSq - i}
	// 	posRight2 := Pos{-i - 1, nSq - i}
	// 	posDown1 := Pos{-nSq + i, -i}
	// 	posDown2 := Pos{-nSq + i, -i - 1}
	// 	posLeft1 := Pos{i, -nSq + i}
	// 	posLeft2 := Pos{i + 1, -nSq + i}

	// 	if posUp1.col == 0 {
	// 		startP := Pos{int64(n/2+1) + n*(posUp1.row-1), 0}
	// 		ms := maxSteps - oneNorm(startP)
	// 		r := poop(gardens, Pos{n - 1, (n / 2) - 1}, ms)
	// 		acc += int64(r)
	// 		// fmt.Println("p1 s", posUp1, ms, r)
	// 	} else {
	// 		startP := Pos{int64(n/2+1) + n*(posUp1.row-1), int64(n/2+1) + n*(posUp1.col-1)}
	// 		ms := maxSteps - oneNorm(startP)
	// 		r := poop(gardens, Pos{n - 1, 0}, ms)
	// 		fmt.Println("p1 s", posUp1, ms, r)
	// 		acc += int64(r)
	// 	}
	// 	startP := Pos{int64(n/2+1) + n*(posUp2.row-1), int64(n/2+1) + n*(posUp2.col-1)}
	// 	ms := maxSteps - oneNorm(startP)
	// 	r := poop(gardens, Pos{n - 1, 0}, ms)
	// 	fmt.Println("p2 s", posUp2, ms, r)
	// 	acc += int64(r)
	// }

	nSq = nSq - 1
	even := (nSq + 1 - nSq%2) * (nSq + 1 - nSq%2)
	odd := (nSq + nSq%2) * (nSq + nSq%2)

	fmt.Println("REM", nSq%2)
	var evenMulti int
	var oddMulti int
	if nSq%2 == 1 { // for some reason we reverse the parity depending on nSq
		evenMulti = poop(gardens, Pos{0, 0}, 4*n+1)
		oddMulti = poop(gardens, Pos{0, 0}, 4*n)
	} else {
		evenMulti = poop(gardens, Pos{0, 0}, 4*n)
		oddMulti = poop(gardens, Pos{0, 0}, 4*n+1)
	}
	acc += even*int64(evenMulti) + odd*int64(oddMulti)
	fmt.Println("acc", acc, evenMulti, oddMulti, even, odd)
	return acc
}

func solve1(gardens [][]string, sPos Pos, maxSteps int64) int {
	fmt.Println(maxSteps)
	q := make([]Step, 0)
	q = append(q, Step{sPos, 0})
	seen := make(map[Pos]bool)
	acc := 0
	i := 0

	// reached := make([][]bool, len(gardens))
	// for i := range reached {
	// 	reached[i] = make([]bool, len(gardens[0]))
	// }

	for len(q) > 0 {
		i++
		curr := q[0]
		q = q[1:]
		if i%10_000_000 == 0 {
		}

		if seen[curr.pos] {
			continue
		}
		seen[curr.pos] = true

		if int64(curr.steps%2) == maxSteps%2 {
			// fmt.Println(curr)
			// reached[curr.pos.row][curr.pos.col] = true
			acc++
		}

		if int64(curr.steps) == maxSteps {
			// fmt.Println(curr)
			continue
		}

		for _, d := range NEIGHBOURS {
			nn := Pos{curr.pos.row + d.row, curr.pos.col + d.col}
			ns := Step{nn, curr.steps + 1}
			if !seen[nn] && gardens[mod(int(nn.row), len(gardens))][mod(int(nn.col), len(gardens[0]))] != "#" {
				q = append(q, ns)
			}
		}
	}

	return acc
}

func poop(gardens [][]string, sPos Pos, maxSteps int64) int {
	fmt.Println(maxSteps, sPos)
	q := make([]Step, 0)
	q = append(q, Step{sPos, 0})
	seen := make(map[Pos]bool)
	acc := 0
	i := 0

	// reached := make([][]bool, len(gardens))
	// for i := range reached {
	// 	reached[i] = make([]bool, len(gardens[0]))
	// }

	for len(q) > 0 {
		i++
		curr := q[0]
		q = q[1:]
		if i%10_000_000 == 0 {
		}

		if seen[curr.pos] {
			continue
		}
		seen[curr.pos] = true

		if int64(curr.steps%2) == maxSteps%2 {
			// fmt.Println(curr)
			// reached[curr.pos.row][curr.pos.col] = true
			acc++
		}

		if int64(curr.steps) == maxSteps {
			// fmt.Println(curr)
			continue
		}

		for _, d := range NEIGHBOURS {
			nn := Pos{curr.pos.row + d.row, curr.pos.col + d.col}
			ns := Step{nn, curr.steps + 1}
			if inBounds(gardens, nn) && !seen[nn] && gardens[nn.row][nn.col] != "#" {
				q = append(q, ns)
			}
		}
	}

	return acc
}

func run(fname string) int64 {
	start := time.Now()
	grid, s := read(fname)
	fmt.Println(s)
	r := solve2(grid, 26501365)
	fmt.Println("Took", time.Since(start))
	// r2 := solve1(grid, s, 327)
	// fmt.Println("r2", r2)
	// p := poop(grid, s, 64)
	// fmt.Println("p", p)
	// fmt.Println(reached)
	// for i := range reached {
	// 	for j := range reached[i] {
	// 		if reached[i][j] {
	// 			fmt.Printf("+")
	// 		} else {
	// 			fmt.Printf(grid[i][j])
	// 		}
	// 	}
	// 	fmt.Println()
	// }
	return r
}

func main() {
	// expected := 16733044
	// rtest := run("inp.test")
	// if rtest != expected {
	// 	log.Fatalln("FAILED TEST CASE", rtest, expected)
	// }
	rmain := run("inp.asen")
	fmt.Println("Result", rmain)
}
