package main

import (
	"fmt"
	"hash/fnv"
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

const MOVEABLE = "O"
const STATIC = "#"
const AIR = "."

func gridToStr(grid [][]string) string {
	var sb strings.Builder
	for _, r := range grid {
		for _, c := range r {
			sb.WriteString(c)
		}
	}
	return sb.String()
}

func read(fname string) [][]string {
	f, err := os.ReadFile(fname)
	data := string(f)
	check(err)
	lines := strings.Split(data, "\n")

	grid := make([][]string, 0)
	for _, l := range lines {
		l = strings.TrimSpace(l)
		row := make([]string, 0)
		for i := range l {
			row = append(row, l[i:i+1])
		}
		grid = append(grid, row)
	}
	return grid
}

func reverseHor(grid [][]string) {
	s := grid
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func flipCol(grid [][]string) {
	n := len(grid[0])
	for i := range grid {
		for j := 0; j < len(grid)/2; j++ {
			grid[i][j], grid[i][n-1-j] = grid[i][n-1-j], grid[i][j]
		}
	}
}

func transpose(slice [][]string) [][]string {
	xl := len(slice[0])
	yl := len(slice)
	result := make([][]string, xl)
	for i := range result {
		result[i] = make([]string, yl)
	}
	for i := 0; i < xl; i++ {
		for j := 0; j < yl; j++ {
			result[i][j] = slice[j][i]
		}
	}
	return result
}

func tiltNorth(grid [][]string) {
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == MOVEABLE {
				k := 1
				for ; i-k >= 0; k++ {
					if grid[i-k][j] != AIR {
						break
					}
					grid[i-k+1][j] = AIR
					grid[i-k][j] = MOVEABLE
				}
			}
		}
	}
}

func tiltWest(grid [][]string) {
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == MOVEABLE {
				k := 1
				for ; j-k >= 0; k++ {
					if grid[i][j-k] != AIR {
						break
					}
					grid[i][j-k+1] = AIR
					grid[i][j-k] = MOVEABLE
				}
			}
		}
	}
}

func printG(grid [][]string) {
	fmt.Println("")
	for _, r := range grid {
		fmt.Println(r)
	}
}

func cycle(grid [][]string) {
	tiltNorth(grid) // N

	tiltWest(grid) // W

	reverseHor(grid)
	tiltNorth(grid) // S
	reverseHor(grid)

	flipCol(grid)
	tiltWest(grid)
	flipCol(grid)
}

func hash(text string) uint32 {
	algorithm := fnv.New32a()
	algorithm.Write([]byte(text))
	return algorithm.Sum32()
}

type SeenData struct {
	lastI int
	count int
}

func eval(grid [][]string) int {
	n := len(grid)
	acc := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == MOVEABLE {
				acc += (n - i)
			}
		}
	}
	return acc
}

func solve(grid [][]string) int {

	seen := make(map[string]*SeenData)
	cycles := 1000000000
	for i := 0; i < cycles; i++ {
		cycle(grid)
		str := gridToStr(grid)
		if seen[str] == nil {
			seen[str] = &SeenData{0, 0}
		}
		if seen[str].count > 0 {
			mul := (cycles - i) / (i - seen[str].lastI)
			i += mul * (i - seen[str].lastI)
		}
		seen[str].count++
		seen[str].lastI = i
	}
	r := eval(grid)

	return r
}

func run(fname string) int {
	grid := read(fname)
	r := solve(grid)
	return r
}
func main() {
	expected := 64
	rtest := run("inp.test")
	if rtest != expected {
		log.Fatalln("FAILED TEST CASE", rtest, expected)
	}
	start := time.Now()
	rmain := run("inp.main")
	fmt.Println("Result", rmain, "Took", time.Since(start))
}
