package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func check(err error) {
	if err != nil {
		log.Fatalf("Fatal error %s", err)
	}
}

type Grid [][]string
type Pos struct {
	row int
	col int
}

func read(fname string) Grid {
	f, err := os.ReadFile(fname)
	data := string(f)
	check(err)
	lines := strings.Split(data, "\n")
	grid := make(Grid, 0)
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

func read2(fname string) Grid {
	f, err := os.ReadFile(fname)
	data := string(f)
	check(err)
	lines := strings.Split(data, "\n")
	grid := make(Grid, 0)
	for _, l := range lines {
		row := make([]string, 0)
		l = strings.TrimSpace(l)
		l = strings.ReplaceAll(l, "^", ".")
		l = strings.ReplaceAll(l, ">", ".")
		l = strings.ReplaceAll(l, "v", ".")
		l = strings.ReplaceAll(l, "<", ".")
		for i := range l {
			row = append(row, l[i:i+1])
		}
		grid = append(grid, row)
	}
	return grid
}

func inBounds(grid Grid, pos Pos) bool {
	return 0 <= pos.row && pos.row < len(grid) && 0 <= pos.col && pos.col < len(grid[0])
}

var NEGHBOURS = []Pos{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
var ARROWS = map[string]int{"^": 0, ">": 1, "v": 2, "<": 3}

func explore(grid Grid, pos Pos, seen map[Pos]bool, acc int) []int {
	END_POS := Pos{len(grid) - 1, len(grid[0]) - 2}
	if pos == END_POS {
		return []int{acc}
	}

	if !inBounds(grid, pos) {
		return []int{}
	}

	if grid[pos.row][pos.col] == "#" {
		return []int{}
	}

	paths := make([]int, 0)
	if grid[pos.row][pos.col] == "." {
		for i := range NEGHBOURS {
			np := Pos{pos.row + NEGHBOURS[i].row, pos.col + NEGHBOURS[i].col}
			if !seen[np] {
				seen[np] = true
				newPaths := explore(grid, np, seen, acc+1)
				paths = append(paths, newPaths...)
				seen[np] = false
			}

		}
	} else {
		i := ARROWS[grid[pos.row][pos.col]]
		np := Pos{pos.row + NEGHBOURS[i].row, pos.col + NEGHBOURS[i].col}
		if !seen[np] {
			seen[np] = true
			newPaths := explore(grid, np, seen, acc+1)
			paths = append(paths, newPaths...)
			seen[np] = false
		}
	}

	return paths
}

type Node struct {
	pos        Pos
	neighbours []Pos
	weights    []uint
}

func getNeighbours(grid Grid, pos Pos) []Pos {
	r := make([]Pos, 0)
	for i := range NEGHBOURS {
		np := Pos{pos.row + NEGHBOURS[i].row, pos.col + NEGHBOURS[i].col}
		if inBounds(grid, np) && grid[np.row][np.col] == "." {
			r = append(r, np)
		}
	}
	return r
}

func findIntersection(grid Grid, pos Pos, seen map[Pos]bool) (Pos, uint) {
	acc := 1
	END_POS := Pos{len(grid) - 1, len(grid[0]) - 2}
	START_POS := Pos{0, 1}
	ns := getNeighbours(grid, pos)
	for len(ns) <= 2 && pos != END_POS && pos != START_POS {
		// fmt.Println(ns, pos)
		seen[pos] = true
		for _, n := range ns {
			if !seen[n] {
				pos = n
				acc++
			}
		}
		ns = getNeighbours(grid, pos)
	}
	return pos, uint(acc)
}

func makeGraph(grid Grid, pos Pos) map[Pos]Node {
	nodes := make(map[Pos]Node)
	q := make([]Pos, 0)
	q = append(q, pos)

	for len(q) > 0 {
		curr := q[0]
		q = q[1:]
		nns := make([]Pos, 0)
		ws := make([]uint, 0)
		node := Node{curr, nns, ws}
		nodes[curr] = node
		for _, n := range getNeighbours(grid, curr) {
			seen := make(map[Pos]bool)
			seen[curr] = true
			ints, w := findIntersection(grid, n, seen)
			_, exists := nodes[ints]
			nns = append(nns, ints)
			ws = append(ws, w)

			if exists {
				continue
			}
			q = append(q, ints)
		}
		node = Node{curr, nns, ws}
		nodes[curr] = node
	}

	return nodes
}

func solve(grid Grid) int {
	seen := make(map[Pos]bool)
	s := Pos{0, 1}
	seen[s] = true
	paths := explore(grid, s, seen, 0)
	maxX := 0
	for _, p := range paths {
		if p > maxX {
			maxX = p
		}
	}
	return maxX
}

func solve2(grid Grid) int {
	nodes := makeGraph(grid, Pos{0, 1})
	seen := make(map[Pos]bool)
	s := Pos{0, 1}
	seen[s] = true
	end := Pos{len(grid) - 1, len(grid[0]) - 2}
	paths := solveGraph(nodes, s, seen, 0, end)
	maxX := uint(0)
	for _, p := range paths {
		if p > maxX {
			maxX = p
		}
	}
	return int(maxX)
}

func solveGraph(nodes map[Pos]Node, pos Pos, seen map[Pos]bool, acc uint, end Pos) []uint {
	if pos == end {
		return []uint{acc}
	}

	node := nodes[pos]
	res := make([]uint, 0)
	for i, n := range node.neighbours {
		if !seen[n] {
			seen[n] = true
			r := solveGraph(nodes, n, seen, acc+node.weights[i], end)
			seen[n] = false
			res = append(res, r...)
		}
	}
	return res
}

func run(fname string) int {
	data := read(fname)
	// fmt.Println(data)
	r := solve(data)
	return r
}

func run2(fname string) int {
	data := read2(fname)
	// fmt.Println(data)
	r := solve2(data)

	return r
}

func main() {
	expected := 154
	rtest := run2("inp.test")
	if rtest != expected {
		log.Fatalln("FAILED TEST CASE", rtest, expected)
	}
	rmain := run2("inp.main")
	fmt.Println("Result", rmain)
}
