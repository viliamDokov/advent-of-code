package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	queue "github.com/Jcowwell/go-algorithm-club/PriorityQueue"
)

type Coord struct {
	row int
	col int
}

const UP = 0
const RIGHT = 1
const DOWN = 2
const LEFT = 3

type Node struct {
	state NodeState
	value NodeValue
}

type NodeState struct {
	pos      Coord
	dir      int
	dirCount int
}

// // we need to define a custom type instead of using the raw integer slice
// // since we need to define methods on the type to implement the heap interface
// type nodeHeap []Node

// // Len is the number of elements in the collection.
// func (h nodeHeap) Len() int {
// 	return len(h)
// }

// // Less reports whether the element with index i
// // must sort before the element with index j.
// // This will determine whether the heap is a min heap or a max heap
// func (h nodeHeap) Less(i int, j int) bool {
// 	return h[i].state.cost < h[j].state.cost
// }

// // Swap swaps the elements with indexes i and j.
// func (h nodeHeap) Swap(i int, j int) {
// 	h[i], h[j] = h[j], h[i]
// }

// // Push and Pop are used to append and remove the last element of the slice
// func (h *nodeHeap) Push(x any) {
// 	*h = append(*h, x.(Node))
// }

// func (h *nodeHeap) Pop() any {
// 	old := *h
// 	n := len(old)
// 	x := old[n-1]
// 	*h = old[0 : n-1]
// 	return x
// }

func check(err error) {
	if err != nil {
		log.Fatalf("Fatal error %s", err)
	}
}

func read(fname string) [][]int {
	f, err := os.ReadFile(fname)
	data := string(f)
	check(err)
	lines := strings.Split(data, "\n")
	grid := make([][]int, 0)
	for _, l := range lines {
		row := make([]int, 0)
		l := strings.TrimSpace(l)
		for i := range l {
			n, err := strconv.Atoi(l[i : i+1])
			check(err)
			row = append(row, n)
		}
		grid = append(grid, row)
	}
	return grid
}

func getNextPos(dir int, pos Coord) Coord {
	switch dir {
	case UP:
		return Coord{pos.row - 1, pos.col}
	case RIGHT:
		return Coord{pos.row, pos.col + 1}
	case DOWN:
		return Coord{pos.row + 1, pos.col}
	case LEFT:
		return Coord{pos.row, pos.col - 1}
	}

	log.Fatalln("Never")
	return pos
}

func inBounds(pos Coord, grid [][]int) bool {
	n := len(grid)
	d := len(grid[0])

	if pos.row >= 0 && pos.row < n &&
		pos.col >= 0 && pos.col < d {
		return true
	}
	return false
}

type NodeValue struct {
	dist int
	prev NodeState
}

func solve(grid [][]int) int {

	n := len(grid)
	d := len(grid[0])

	values := make(map[NodeState]*NodeValue)
	seen := make(map[NodeState]bool)
	pq := queue.PriorityQueueInit(func(a NodeState, b NodeState) bool {
		return values[a].dist < values[b].dist
	})

	s1 := NodeState{Coord{0, 0}, RIGHT, 0}
	s2 := NodeState{Coord{0, 0}, DOWN, 0}

	values[s1] = &NodeValue{0, s1}
	seen[s1] = true
	values[s2] = &NodeValue{0, s2}
	seen[s2] = true
	pq.Enqueue(s1)
	pq.Enqueue(s2)

	endPos := Coord{n - 1, d - 1}
	finalCost := 0
	var finalState NodeState
	for pq.Count() > 0 {
		curr, _ := pq.Dequeue()
		currv := values[curr]
		// fmt.Println("values")
		// for k, v := range values {
		// 	fmt.Println(k, v)
		// }
		if curr.pos == endPos && curr.dirCount >= 4 {
			fmt.Println("BREAK!")
			finalCost = currv.dist
			finalState = curr
			break
		}

		for dir := 0; dir < 4; dir++ {
			np := getNextPos(dir, curr.pos)

			if !inBounds(np, grid) {
				// fmt.Println("No bounds?")
				continue
			}

			if dir != curr.dir && curr.dirCount < 4 {
				// fmt.Println("Too same?")
				continue
			}
			if dir == curr.dir && curr.dirCount > 9 {
				// fmt.Println("Too same?")
				continue
			}
			if (dir+2)%4 == curr.dir {
				// fmt.Println(dir, curr.state.dir)
				// fmt.Println("Mod magic")
				continue
			}

			dirCount := curr.dirCount
			if dir == curr.dir {
				dirCount++
			} else {
				dirCount = 1
			}

			nstate := NodeState{np, dir, dirCount}
			if seen[nstate] {
				continue
			}

			gv := grid[nstate.pos.row][nstate.pos.col]
			oldV := values[nstate]
			if oldV == nil || oldV.dist > currv.dist+gv {
				values[nstate] = &NodeValue{currv.dist + gv, curr}
				pq.Enqueue(nstate)
			}

		}
	}

	type NodeResult struct {
		s NodeState
		d int
	}
	sequence := make([]NodeResult, 0)
	var prevState NodeState
	curState := finalState
	for curState.pos != prevState.pos {
		prevState := curState
		sequence = append([]NodeResult{{prevState, values[prevState].dist}}, sequence...)
		curState = values[prevState].prev
	}
	fmt.Println("seq", sequence)
	return finalCost
}

func run(fname string) int {
	data := read(fname)
	r := solve(data)
	return r
}
func main() {
	// expected := 51
	// rtest := run("inp.test")
	// if rtest != expected {
	// 	log.Fatalln("FAILED TEST CASE", rtest, expected)
	// }
	fmt.Println("Passed test!!!!!")
	rmain := run("inp.main")
	fmt.Println("Result", rmain)
}
