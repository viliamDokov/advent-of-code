package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type Coord struct {
	row int
	col int
}

type BeamState struct {
	p   Coord
	dir int
}

const UP = 0
const RIGHT = 1
const DOWN = 2
const LEFT = 3

func check(err error) {
	if err != nil {
		log.Fatalf("Fatal error %s", err)
	}
}

func read(fname string) [][]string {
	f, err := os.ReadFile(fname)
	data := string(f)
	check(err)
	lines := strings.Split(data, "\n")
	grid := make([][]string, 0)
	for _, l := range lines {
		row := make([]string, 0)
		l = strings.TrimSpace(l)
		for i := range l {
			row = append(row, l[i:i+1])
		}
		grid = append(grid, row)
	}
	return grid
}

func getNextPos(s BeamState) BeamState {
	switch s.dir {
	case UP:
		s.p.row--
		return s
	case RIGHT:
		s.p.col++
		return s
	case DOWN:
		s.p.row++
		return s
	case LEFT:
		s.p.col--
		return s
	default:
		log.Fatalln("NEVER", s.dir)
	}
	log.Fatalln("NEVER", s.dir)
	return s
}

func inBounds(s BeamState, grid [][]string) bool {
	if 0 <= s.p.row && s.p.row < len(grid) && 0 <= s.p.col && s.p.col < len(grid[0]) {
		return true
	}
	return false
}

var nexDirsCW = []int{LEFT, DOWN, RIGHT, UP}
var nexDirsCC = []int{RIGHT, UP, LEFT, DOWN}

func part2(grid [][]string) int {
	best := 0
	rows, cols := len(grid), len(grid[0])
	for i := range grid {
		s1 := BeamState{Coord{-1, i}, DOWN}
		s2 := BeamState{Coord{rows, i}, UP}
		v1 := part1(s1, grid)
		v2 := part1(s2, grid)
		if v1 > best {
			best = v1
		}
		if v2 > best {
			best = v2
		}
	}

	for i := range grid[0] {
		s1 := BeamState{Coord{i, -1}, DOWN}
		s2 := BeamState{Coord{i, cols}, UP}
		v1 := part1(s1, grid)
		v2 := part1(s2, grid)
		if v1 > best {
			best = v1
		}
		if v2 > best {
			best = v2
		}
	}

	return best
}

func part1(start BeamState, grid [][]string) int {
	q := make([]BeamState, 0)
	q = append(q, start)
	seen := make(map[BeamState]bool)
	coords := make(map[Coord]bool)
	for len(q) > 0 {
		s := q[0]
		q = q[1:]
		if seen[s] {
			continue
		}
		coords[s.p] = inBounds(s, grid)
		seen[s] = inBounds(s, grid)
		ns := getNextPos(s)
		if inBounds(ns, grid) {
			tile := grid[ns.p.row][ns.p.col]
			switch tile {
			case ".":
				q = append(q, ns)
			case "/":
				nd := nexDirsCC[ns.dir]
				ns.dir = nd
				q = append(q, ns)
			case "\\":
				nd := nexDirsCW[ns.dir]
				ns.dir = nd
				q = append(q, ns)
			case "|":
				if ns.dir == LEFT || ns.dir == RIGHT {
					nd1, nd2 := nexDirsCW[ns.dir], nexDirsCC[ns.dir]
					ns1, ns2 := BeamState{ns.p, ns.dir}, BeamState{ns.p, ns.dir}
					ns1.dir, ns2.dir = nd1, nd2
					q = append(q, ns1)
					q = append(q, ns2)
				} else {
					q = append(q, ns)
				}
			case "-":
				if ns.dir == UP || ns.dir == DOWN {
					nd1, nd2 := nexDirsCW[ns.dir], nexDirsCC[ns.dir]
					ns1, ns2 := BeamState{ns.p, ns.dir}, BeamState{ns.p, ns.dir}
					ns1.dir, ns2.dir = nd1, nd2
					q = append(q, ns1)
					q = append(q, ns2)
				} else {
					q = append(q, ns)
				}
			}
		}
	}
	return len(coords) - 1
}

func run(fname string) int {
	grid := read(fname)
	r := part2(grid)
	return r
}

func main() {
	expected := 51
	rtest := run("inp.test")
	if rtest != expected {
		log.Fatalln("FAILED TEST CASE", rtest, expected)
	}
	fmt.Println("=========== Pass main! ======================")
	rmain := run("inp.main")
	fmt.Println("Result", rmain)
}
