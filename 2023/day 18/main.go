package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type DigStep struct {
	dir    int64
	length int64
}

func check(err error) {
	if err != nil {
		log.Fatalf("Fatal error %s", err)
	}
}

var dirMap = map[string]int64{"U": 0, "R": 1, "D": 2, "L": 3}
var dirCoord = []Coord{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

func read(fname string) []DigStep {
	f, err := os.ReadFile(fname)
	data := string(f)
	check(err)
	lines := strings.Split(data, "\n")
	digs := make([]DigStep, 0)
	for _, l := range lines {
		l := strings.TrimSpace(l)
		tmp := strings.Split(l, " ")
		dirs, dists := tmp[0], tmp[1]

		dir := int64(dirMap[dirs])
		check(err)
		length, err := strconv.ParseInt(dists, 10, 64)

		check(err)
		digs = append(digs, DigStep{dir, length})
	}
	return digs
}

func read2(fname string) []DigStep {
	f, err := os.ReadFile(fname)
	data := string(f)
	check(err)
	lines := strings.Split(data, "\n")
	digs := make([]DigStep, 0)
	for _, l := range lines {
		l := strings.TrimSpace(l)
		tmp := strings.Split(l, " ")
		hex := tmp[2][2 : len(tmp[2])-1]
		dirs, dists := hex[len(hex)-1:], hex[:len(hex)-1]

		dir, err := strconv.Atoi(dirs)
		dir64 := int64((dir + 1) % 4)
		check(err)
		length, err := strconv.ParseInt(dists, 16, 64)

		check(err)
		digs = append(digs, DigStep{dir64, length})
	}
	return digs
}

type Coord struct {
	row int64
	col int64
}

type HorLine struct {
	start Coord
	end   Coord
}

func solve(digs []DigStep) int64 {
	gridMap := make(map[Coord]bool)
	curr := Coord{0, 0}
	gridMap[curr] = true
	rows := int64(0)
	cols := int64(0)

	minrows := int64(0)
	mincols := int64(0)
	for _, d := range digs {

		for i := int64(0); i < d.length; i++ {
			diff := dirCoord[d.dir]
			nc := Coord{curr.row + diff.row, curr.col + diff.col}

			if nc.row < minrows {
				minrows = nc.row
			}
			if nc.col < mincols {
				mincols = nc.col
			}
			gridMap[nc] = true
			curr = nc
		}
	}

	offsetMap := make(map[Coord]bool)
	for k := range gridMap {
		offsetMap[Coord{k.row - minrows, k.col - mincols}] = true
	}

	for nc := range offsetMap {
		if nc.row > rows {
			rows = nc.row
		}
		if nc.col > cols {
			cols = nc.col
		}
	}

	rows = rows + 1
	cols = cols + 1

	grid := make([][]bool, rows)
	for i := range grid {
		grid[i] = make([]bool, cols)
	}

	for k := range offsetMap {
		grid[k.row][k.col] = true
	}

	q := make([]Coord, 0)
	seen := make(map[Coord]bool)
	for i := int64(0); i < rows; i++ {
		tc := Coord{i, 0}
		bc := Coord{i, cols - 1}
		if !grid[tc.row][tc.col] {
			seen[tc] = true
			q = append(q, tc)
		}
		if !grid[bc.row][bc.col] {
			seen[bc] = true
			q = append(q, bc)
		}
	}
	for i := int64(0); i < cols; i++ {
		tc := Coord{0, i}
		bc := Coord{rows - 1, i}
		if !grid[tc.row][tc.col] {
			seen[tc] = true
			q = append(q, tc)
		}
		if !grid[bc.row][bc.col] {
			seen[bc] = true
			q = append(q, bc)
		}
	}

	for len(q) > 0 {
		curr := q[0]
		q = q[1:]
		for i := 0; i < 4; i++ {
			diff := dirCoord[i]
			nd := Coord{curr.row + diff.row, curr.col + diff.col}
			if seen[nd] {
				continue
			}
			if !(0 <= nd.row && nd.row < rows && 0 <= nd.col && nd.col < cols) {
				continue
			}
			if grid[nd.row][nd.col] {
				continue
			}
			seen[nd] = true
			q = append(q, nd)
		}
	}
	return rows*cols - int64(len(seen))
}

type Range struct {
	start int64
	end   int64
}

func mergeRanges(ranges []Range, newRanges []Range) []Range {

	sort.Slice(newRanges, func(i, j int) bool { return newRanges[i].start < newRanges[j].start })
	sort.Slice(ranges, func(i, j int) bool { return ranges[i].start < ranges[j].start })

	op := 0
	np := 0

	resultR := make([]Range, 0)
	fmt.Println(ranges, newRanges)
	if op >= len(ranges) {
		resultR = append(resultR, newRanges[np:]...)
		return resultR
	}
	if np >= len(newRanges) {
		resultR = append(resultR, ranges[op:]...)
		return resultR
	}

	or := ranges[op]
	nr := newRanges[np]

	moveOR := false
	moveNR := false
	for {

		if nr.start < or.start && nr.end <= or.start {
			resultR = append(resultR, nr)
			moveNR = true
		} else if nr.start < or.start && nr.end <= or.end {
			resultR = append(resultR, Range{nr.start, or.start})
			moveNR = true
			or = Range{nr.end, or.end}
		} else if nr.start < or.start && nr.end > or.end {
			resultR = append(resultR, Range{nr.start, or.start})
			nr = Range{or.end, nr.end}
			moveOR = true
		} else if nr.start < or.end && nr.end <= or.end {
			resultR = append(resultR, Range{or.start, nr.start})
			moveNR = true
			or = Range{nr.end, or.end}
		} else if nr.start < or.end && nr.end > or.end {
			resultR = append(resultR, Range{or.start, nr.start})
			nr = Range{or.end, nr.end}
			moveOR = true
		} else if nr.start >= or.end {
			resultR = append(resultR, or)
			moveOR = true
		} else {
			log.Fatalln("never happen", nr, or)
		}

		if moveNR {
			np++
			if np >= len(newRanges) {
				break
			}
			moveNR = false
			nr = newRanges[np]
		}
		if moveOR {
			op++
			if op >= len(ranges) {
				break
			}
			or = ranges[op]
			moveOR = false
		}

	}

	if op >= len(ranges) {
		resultR = append(resultR, nr)
		resultR = append(resultR, newRanges[np+1:]...)
	}
	if np >= len(newRanges) {
		resultR = append(resultR, or)
		resultR = append(resultR, ranges[op+1:]...)
	}

	for i := 0; i < len(resultR); {
		if resultR[i].start == resultR[i].end {
			resultR = append(resultR[:i], resultR[i+1:]...)
		} else {
			i++
		}
	}

	for i := 0; i < len(resultR)-1; {
		if resultR[i].end == resultR[i+1].start {
			resultR[i+1].start = resultR[i].start
			resultR = append(resultR[:i], resultR[i+1:]...)
		} else {
			i++
		}
	}

	// sort.Slice(resultR, func(i, j int) bool { return resultR[i].start < resultR[j].start })
	return resultR
}

func getRangeThiccccc(ranges []Range) int64 {
	acc := int64(0)
	for _, r := range ranges {
		acc += r.end - r.start
	}
	return acc
}

func solve2(digs []DigStep, bonus int64) int64 {
	horLines := make([]HorLine, 0)
	curr := Coord{0, 0}

	for i := 0; i < len(digs); i++ {
		prevD := digs[(i+len(digs)-1)%len(digs)].dir
		nextD := digs[(i+1)%len(digs)].dir

		d := digs[i]
		var offset int64
		if prevD == nextD {
			bonus = bonus * -1
			offset = 0
		} else {
			offset = bonus
		}

		var nc Coord
		switch d.dir {
		case 0:
			nc = Coord{curr.row - (d.length + offset), curr.col}
		case 1:
			nc = Coord{curr.row, curr.col + (d.length + offset)}
		case 2:
			nc = Coord{curr.row + (d.length + offset), curr.col}
		case 3:
			nc = Coord{curr.row, curr.col - (d.length + offset)}
		default:
			log.Fatalln("Never")
		}

		if d.dir == 1 {
			horLines = append(horLines, HorLine{curr, nc})
		}
		if d.dir == 3 {
			horLines = append(horLines, HorLine{nc, curr})
		}
		fmt.Println(nc, bonus, offset)
		curr = nc
	}

	sort.Slice(horLines, func(i int, j int) bool { return horLines[i].start.row < horLines[j].start.row })
	fmt.Println(horLines)

	currRow := horLines[0].start.row
	oldRow := currRow
	acc := int64(0)
	ranges := make([]Range, 0)
	for len(horLines) > 0 {
		currRow = horLines[0].start.row
		thicc := getRangeThiccccc(ranges)
		acc += thicc * (currRow - oldRow)
		fmt.Println(thicc, (currRow - oldRow), acc, ranges)
		newRanges := make([]Range, 0)
		for len(horLines) > 0 && horLines[0].start.row == currRow {
			newRanges = append(newRanges, Range{horLines[0].start.col, horLines[0].end.col})
			horLines = horLines[1:]
		}
		ranges = mergeRanges(ranges, newRanges)

		oldRow = currRow
	}

	return acc
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func run(fname string) int64 {
	data := read2(fname)
	// for i := range data {
	// 	fmt.Println(data[i])
	// }
	r := max(solve2(data, -1), solve2(data, 1))
	return r
}
func main() {
	// 952408144115
	expected := int64(952408144115)
	rtest := run("inp.test")
	if rtest != expected {
		log.Fatalln("FAILED TEST CASE", rtest, expected)
	}
	fmt.Println("Passed Test!")
	rmain := run("inp.main")
	fmt.Println("Result", rmain)
}
