package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type Number struct {
	val    int
	row    int
	col    int
	length int
}

func check(err error) {
	if err != nil {
		log.Fatalf("Fatal error %s", err)
	}
}

// var adjacents = []Coord{
// 	{-1, -1},
// 	{-1, 0},
// 	{-1, 1},
// 	{0, -1},
// 	{0, 1},
// 	{1, -1},
// 	{1, 0},
// 	{1, 1},
// }

// func getnum(lines []string, cs []Coord) int {
// 	s := 0
// 	for i, c := range cs {
// 		nc := lines[c.row][c.col : c.col+1]
// 		n, err := strconv.Atoi(nc)
// 		check(err)
// 		s += n * int(math.Pow(10, float64(len(cs)-i)))
// 	}
// 	return s
// }

type Coord struct {
	row int
	col int
}

var adjacents = []Coord{
	{-1, -1},
	{-1, 0},
	{-1, 1},
	{0, -1},
	{0, 1},
	{1, -1},
	{1, 0},
	{1, 1},
}

var invalid = []byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0', '.'}

func isPart(lines []string, num Number, gears map[Coord][]int) bool {
	for y := 0; y < num.length; y++ {
		for _, adj := range adjacents {
			nx, ny := num.row+adj.row, num.col+y+adj.col
			if nx < 0 || nx >= len(lines) {
				continue
			}
			if ny < 0 || ny >= len(lines[nx]) {
				continue
			}
			char := lines[nx][ny]
			bad := false
			for _, inv := range invalid {
				if char == inv {
					bad = true
					break
				}
			}
			if char == '*' {
				gears[Coord{nx, ny}] = append(gears[Coord{nx, ny}], num.val)
			}
			if !bad {
				return true
			}
		}
	}
	fmt.Println("Not a part!", num)
	return false
}

func solve(lines []string, nums []Number, gears map[Coord][]int) int {
	for _, num := range nums {
		isPart(lines, num, gears)
	}
	s := 0
	for _, nums := range gears {
		if len(nums) == 2 {
			s += nums[0] * nums[1]
		}
	}
	return s
}

func read() ([]string, []Number, map[Coord][]int) {
	f, err := os.Open("inp")
	check(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	// optionally, resize scanner's capacity for lines over 64K, see next example

	lines := make([]string, 0)
	nums := make([]Number, 0)
	gears := make(map[Coord][]int)
	j := 0
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
		re := regexp.MustCompile(`(\d+)`)
		regear := regexp.MustCompile(`\*`)

		gindexes := regear.FindAllStringIndex(line, -1)
		for _, gin := range gindexes {
			c := Coord{j, gin[0]}
			gears[c] = make([]int, 0)
		}

		indexes := re.FindAllStringIndex(line, -1)
		numbers := re.FindAllString(line, -1)

		for i, num := range numbers {
			val, err := strconv.Atoi(num)
			check(err)
			st, e := indexes[i][0], indexes[i][1]
			l := e - st
			nb := Number{val: val, row: j, col: st, length: l}
			nums = append(nums, nb)
		}
		j++
	}
	return lines, nums, gears
}

func main() {
	lines, nums, gears := read()
	fmt.Println(lines, nums)
	res := solve(lines, nums, gears)
	fmt.Println(res)
}
