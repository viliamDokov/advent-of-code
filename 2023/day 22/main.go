package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Pos [3]int

type Brick struct {
	st  Pos
	end Pos
	idx int
}

func check(err error) {
	if err != nil {
		log.Fatalf("Fatal error %s", err)
	}
}

func read(fname string) []Brick {
	f, err := os.ReadFile(fname)
	data := string(f)
	check(err)
	bricks := make([]Brick, 0)
	lines := strings.Split(data, "\n")
	for i, l := range lines {
		l = strings.TrimSpace(l)
		tmp := strings.Split(l, "~")
		c1spl := strings.Split(tmp[0], ",")
		c2spl := strings.Split(tmp[1], ",")
		var p1 Pos
		var p2 Pos
		for i := range c1spl {
			c, err := strconv.Atoi(c1spl[i])
			check(err)
			p1[i] = c
			c2, err := strconv.Atoi(c2spl[i])
			check(err)
			p2[i] = c2
		}

		bricks = append(bricks, Brick{p1, p2, i})
		// fmt.Println(l)
	}
	return bricks
}

func makeGrid(bricks []Brick) [][][]int {
	var maxPos Pos

	for _, b := range bricks {
		for i := range b.st {
			if maxPos[i] < b.st[i] {
				maxPos[i] = b.st[i]
			}
			if maxPos[i] < b.end[i] {
				maxPos[i] = b.end[i]
			}
		}

	}

	grid := make([][][]int, maxPos[0]+1)
	for i := range grid {
		grid[i] = make([][]int, maxPos[1]+1)
		for j := range grid[i] {
			grid[i][j] = make([]int, maxPos[2]+1)
			for k := range grid[i][j] {
				grid[i][j][k] = -1
			}
		}
	}
	return grid
}

func collapse(bricks []Brick, grid [][][]int) int {
	sort.Slice(bricks, func(i, j int) bool {
		return bricks[i].st[2] < bricks[j].st[2]
	})
	// fmt.Println(bricks)

	collapsed := 0
	for i := range bricks {
		// fmt.Println(i)
		falls := 0
		for {
			if falls == 1 {
				collapsed++
			}
			if bricks[i].st[2] == 1 {
				break
			}
			// fmt.Println(bricks[i])
			diff := Pos{bricks[i].end[0] - bricks[i].st[0], bricks[i].end[1] - bricks[i].st[1], bricks[i].end[2] - bricks[i].st[2]}
			nz := 0
			for i := range diff {
				if diff[i] != 0 {
					nz = i
					break
				}
			}

			count := diff[nz]
			if diff[nz] < 0 {
				count = -diff[nz]
			}

			curr := bricks[i].st
			stop := false
			for j := 0; j <= count; j++ {
				below := grid[curr[0]][curr[1]][curr[2]-1]
				// fmt.Println(curr, below, bricks[i], diff, nz, count)
				if below != bricks[i].idx && below != -1 {
					stop = true
					break
				}

				if diff[nz] > 0 {
					curr[nz]++
				} else {
					curr[nz]--
				}
			}

			if stop {
				break
			}

			curr = bricks[i].st
			for j := 0; j <= count; j++ {
				grid[curr[0]][curr[1]][curr[2]-1] = bricks[i].idx
				grid[curr[0]][curr[1]][curr[2]] = -1

				if diff[nz] > 0 {
					curr[nz]++
				} else {
					curr[nz]--
				}
			}
			bricks[i].st[2]--
			bricks[i].end[2]--
			falls++
		}
	}

	return collapsed
}

func solve2(bricks []Brick) int {
	grid := makeGrid(bricks)

	for _, br := range bricks {
		diff := Pos{br.end[0] - br.st[0], br.end[1] - br.st[1], br.end[2] - br.st[2]}
		nz := 0
		for i := range diff {
			if diff[i] != 0 {
				nz = i
				break
			}
		}

		nonNeg := diff[nz] > 0
		if !nonNeg {
			diff[nz] = -diff[nz]
		}

		curr := br.st
		for i := 0; i <= diff[nz]; i++ {

			grid[curr[0]][curr[1]][curr[2]] = br.idx
			if nonNeg {
				curr[nz]++
			} else {
				curr[nz]--
			}
		}
	}

	collapse(bricks, grid)
	// fmt.Println(bricks, grid)

	acc := 0
	for q := range bricks {
		newBricks := make([]Brick, 0)
		for j := range bricks {
			if q != j {
				newBricks = append(newBricks, bricks[j])
			}
		}

		newGrid := make([][][]int, len(grid))
		for i := range newGrid {
			newGrid[i] = make([][]int, len(grid[i]))
			for j := range grid[i] {
				newGrid[i][j] = make([]int, len(grid[i][j]))
				for k := range newGrid[i][j] {
					newGrid[i][j][k] = grid[i][j][k]
				}
			}
		}
		// fmt.Println()
		// fmt.Println(newGrid)

		br := bricks[q]
		diff := Pos{br.end[0] - br.st[0], br.end[1] - br.st[1], br.end[2] - br.st[2]}
		nz := 0
		for i := range diff {
			if diff[i] != 0 {
				nz = i
				break
			}
		}

		nonNeg := diff[nz] > 0
		if !nonNeg {
			diff[nz] = -diff[nz]
		}

		curr := br.st
		for i := 0; i <= diff[nz]; i++ {
			newGrid[curr[0]][curr[1]][curr[2]] = -1
			if nonNeg {
				curr[nz]++
			} else {
				curr[nz]--
			}
		}

		c := collapse(newBricks, newGrid)
		// fmt.Println("collapsing", br, c)
		acc += c

	}

	return acc
}

func solve(bricks []Brick) int {
	grid := makeGrid(bricks)
	for _, br := range bricks {
		diff := Pos{br.end[0] - br.st[0], br.end[1] - br.st[1], br.end[2] - br.st[2]}
		nz := 0
		for i := range diff {
			if diff[i] != 0 {
				nz = i
				break
			}
		}

		nonNeg := diff[nz] > 0
		if !nonNeg {
			diff[nz] = -diff[nz]
		}

		curr := br.st
		for i := 0; i <= diff[nz]; i++ {
			// fmt.Println("Poop", br, curr)

			grid[curr[0]][curr[1]][curr[2]] = br.idx
			if nonNeg {
				curr[nz]++
			} else {
				curr[nz]--
			}
		}
	}

	sort.Slice(bricks, func(i, j int) bool {
		return bricks[i].st[2] < bricks[j].st[2]
	})
	// fmt.Println(bricks)

	supported := make([]map[int]bool, len(bricks))
	for i := range supported {
		supported[i] = make(map[int]bool, 0)
	}

	for i := range bricks {
		// fmt.Println(i)
		for {
			if bricks[i].st[2] == 1 {
				break
			}

			// fmt.Println(bricks[i])
			diff := Pos{bricks[i].end[0] - bricks[i].st[0], bricks[i].end[1] - bricks[i].st[1], bricks[i].end[2] - bricks[i].st[2]}
			nz := 0
			for i := range diff {
				if diff[i] != 0 {
					nz = i
					break
				}
			}

			count := diff[nz]
			if diff[nz] < 0 {
				count = -diff[nz]
			}

			curr := bricks[i].st
			for j := 0; j <= count; j++ {
				below := grid[curr[0]][curr[1]][curr[2]-1]
				// fmt.Println(curr, below, bricks[i], diff, nz, count)
				if below != bricks[i].idx && below != -1 {
					if !supported[bricks[i].idx][below] { // Windows flags this as a virus if we remove the if statement
						supported[bricks[i].idx][below] = true
					}
				}

				if diff[nz] > 0 {
					curr[nz]++
				} else {
					curr[nz]--
				}
			}

			if len(supported[bricks[i].idx]) > 0 {
				break
			}

			curr = bricks[i].st
			for j := 0; j <= count; j++ {
				grid[curr[0]][curr[1]][curr[2]-1] = bricks[i].idx
				grid[curr[0]][curr[1]][curr[2]] = -1

				if diff[nz] > 0 {
					curr[nz]++
				} else {
					curr[nz]--
				}
			}
			bricks[i].st[2]--
			bricks[i].end[2]--
		}
	}

	// fmt.Println("supp", supported)
	vital := make(map[int]bool)
	for _, mp := range supported {
		if len(mp) == 1 {
			for k := range mp {
				vital[k] = true
			}
		}
	}
	return len(bricks) - len(vital)
}

func run(fname string) int {
	bricks := read(fname)
	// fmt.Println(bricks)
	r := solve2(bricks)
	return r
}

func main() {
	expected := 7
	rtest := run("inp.test")
	if rtest != expected {
		log.Fatalln("FAILED TEST CASE", rtest, expected)
	}
	rmain := run("inp.main")
	fmt.Println("Result", rmain)
}
