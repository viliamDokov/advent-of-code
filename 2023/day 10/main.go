package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type Node struct {
	loc   Coord
	conns []bool
	d     int
}

type Coord struct {
	r int
	c int
}

func check(err error) {
	if err != nil {
		log.Fatalf("Fatal error %s", err)
	}
}

func read(fname string) ([]string, int, int) {
	f, err := os.ReadFile(fname)
	data := string(f)
	check(err)
	lines := strings.Split(data, "\n")
	var sr int
	var sc int

	for row, l := range lines {
		lines[row] = strings.TrimSpace(lines[row])
		for col := range l {
			c := l[col : col+1]
			if c == "S" {
				sr = row
				sc = col
			}
		}
	}

	return lines, sr, sc
}

func convertC(chr string) []bool {
	// fmt.Println(chr)
	switch chr {
	case "|":
		return []bool{true, false, true, false}
	case "-":
		return []bool{false, true, false, true}
	case "L":
		return []bool{true, true, false, false}

	case "J":
		return []bool{true, false, false, true}

	case "7":
		return []bool{false, false, true, true}

	case "F":
		return []bool{false, true, true, false}
	case ".":
		return []bool{false, false, false, false}

	case "S":
		return []bool{true, true, true, true}
	default:
		log.Fatalf("Invalid %s \n", chr)
	}
	log.Fatalln("Invalid AAAAAAAAAA")
	return []bool{true, true, true, true}
}

func getNS(lines []string, r int, c int) []Node {
	// fmt.Println("DAddy", r, c, lines[r][c:c+1])
	ns := make([]Node, 0)

	bg := convertC(lines[r][c : c+1])

	if r-1 >= 0 {
		arr := convertC(lines[r-1][c : c+1])
		if arr[2] && bg[0] {
			c := Coord{r - 1, c}
			n := Node{c, arr, 0}
			// fmt.Println("add", n)
			ns = append(ns, n)
		}
	}
	if r+1 < len(lines) {
		arr := convertC(lines[r+1][c : c+1])
		if arr[0] && bg[2] {
			c := Coord{r + 1, c}
			n := Node{c, arr, 0}
			ns = append(ns, n)
			// fmt.Println("add", n)
		}
	}
	if c-1 >= 0 {
		arr := convertC(lines[r][c-1 : c])
		if arr[1] && bg[3] {
			c := Coord{r, c - 1}
			n := Node{c, arr, 0}
			ns = append(ns, n)
			// fmt.Println("add", n)

		}
	}
	if c+1 < len(lines[0]) {
		arr := convertC(lines[r][c+1 : c+2])
		if arr[3] && bg[1] {
			c := Coord{r, c + 1}
			n := Node{c, arr, 0}
			ns = append(ns, n)
			// fmt.Println("add", n)

		}
	}
	return ns
}

func solve(lines []string, sr int, sc int) int {

	i, j := sr, sc
	seen := make(map[Coord]bool)

	q := make([]Node, 0)
	c := Coord{i, j}
	q = append(q, Node{c, []bool{true, true, true, true}, 0})
	seen[c] = true
	maxD := 0

	for len(q) > 0 {
		guy := q[0]
		// fmt.Println("POP", guy, lines[guy.loc.r][guy.loc.c:guy.loc.c+1])
		q = q[1:]
		ns := getNS(lines, guy.loc.r, guy.loc.c)
		for _, n := range ns {
			if !seen[n.loc] {
				n.d = guy.d + 1
				if n.d > maxD {
					maxD = n.d
				}
				seen[n.loc] = true
				q = append(q, n)
			}
		}
	}

	// fmt.Println(seen)
	return maxD
}

func convert(lines []string) [][]int {
	dots2 := 0
	newPipes := make([][]int, 0)
	for _, row := range lines {
		r1 := make([]int, 0)
		r2 := make([]int, 0)
		r3 := make([]int, 0)
		for _, c := range row {
			switch c {
			case '|':
				r1 = append(r1, []int{0, 1, 0}...)
				r2 = append(r2, []int{0, 1, 0}...)
				r3 = append(r3, []int{0, 1, 0}...)
			case '-':
				r1 = append(r1, []int{0, 0, 0}...)
				r2 = append(r2, []int{1, 1, 1}...)
				r3 = append(r3, []int{0, 0, 0}...)
			case 'L':
				r1 = append(r1, []int{0, 1, 0}...)
				r2 = append(r2, []int{0, 1, 1}...)
				r3 = append(r3, []int{0, 0, 0}...)
			case 'J':
				r1 = append(r1, []int{0, 1, 0}...)
				r2 = append(r2, []int{1, 1, 0}...)
				r3 = append(r3, []int{0, 0, 0}...)
			case '7':
				r1 = append(r1, []int{0, 0, 0}...)
				r2 = append(r2, []int{1, 1, 0}...)
				r3 = append(r3, []int{0, 1, 0}...)
			case 'F':
				r1 = append(r1, []int{0, 0, 0}...)
				r2 = append(r2, []int{0, 1, 1}...)
				r3 = append(r3, []int{0, 1, 0}...)
			case '.':
				r1 = append(r1, []int{0, 0, 0}...)
				r2 = append(r2, []int{0, 2, 0}...)
				r3 = append(r3, []int{0, 0, 0}...)
				dots2++
			default:
				log.Fatalln("Invalid", c)

			}
		}
		newPipes = append(newPipes, r1)
		newPipes = append(newPipes, r2)
		newPipes = append(newPipes, r3)
	}
	fmt.Println("DOTS2", dots2)
	return newPipes
}

func solve2(lines []string, sr int, sc int) int {
	i, j := sr, sc
	seen := make(map[Coord]bool)

	q := make([]Node, 0)
	c := Coord{i, j}
	q = append(q, Node{c, []bool{true, true, true, true}, 0})
	seen[c] = true
	maxD := 0

	for len(q) > 0 {
		guy := q[0]
		// fmt.Println("POP", guy, lines[guy.loc.r][guy.loc.c:guy.loc.c+1])
		q = q[1:]
		ns := getNS(lines, guy.loc.r, guy.loc.c)
		for _, n := range ns {
			if !seen[n.loc] {
				n.d = guy.d + 1
				if n.d > maxD {
					maxD = n.d
				}
				seen[n.loc] = true
				q = append(q, n)
			}
		}
	}

	rs := len(lines)
	cs := len(lines[0])

	// infer start
	sarr := []bool{false, false, false, false}
	if sr-1 >= 0 {
		bm := convertC(lines[sr-1][sc : sc+1])
		sarr[0] = bm[2]
	}
	if sc+1 < cs {
		bm := convertC(lines[sr][sc+1 : sc+2])
		sarr[1] = bm[3]
	}
	if sr+1 < rs {
		bm := convertC(lines[sr+1][sc : sc+1])
		sarr[2] = bm[0]
	}
	if sc-1 >= 0 {
		bm := convertC(lines[sr][sc-1 : sc])
		sarr[3] = bm[1]
	}

	sval := "89"
	if sarr[0] {
		if sarr[1] {
			sval = "L"
		}
		if sarr[2] {
			sval = "|"
		}
		if sarr[3] {
			sval = "J"
		}
	}

	if sarr[1] {
		if sarr[2] {
			sval = "F"
		}
		if sarr[3] {
			sval = "-"
		}
	}

	if sarr[2] {
		if sarr[3] {
			sval = "7"
		}
	}

	fmt.Println("SVAL", sval, sarr, sr, sc)
	lines[sr] = lines[sr][:sc] + sval + lines[sr][sc+1:]

	for i := 0; i < rs; i++ {
		for j := 0; j < cs; j++ {
			c := Coord{i, j}
			if !seen[c] {
				lines[i] = lines[i][:j] + "." + lines[i][j+1:]
			}
		}
	}

	dots := 0
	for i := 0; i < rs; i++ {
		for j := 0; j < cs; j++ {
			if lines[i][j:j+1] == "." {
				dots++
			}
		}
	}
	fmt.Println("DOTS", dots)

	bigDaddy := convert(lines)
	rd := len(bigDaddy)
	cd := len(bigDaddy[0])

	dc := 0
	for i := 0; i < rd; i++ {
		fmt.Println(bigDaddy[i])
		for j := 0; j < cd; j++ {
			if bigDaddy[i][j] == 2 {
				dc++
			}
		}
	}

	fmt.Println("DC", dc)

	q2 := make([]Coord, 0)

	newSeen := make(map[Coord]bool)
	for i := 0; i < rd; i++ {
		c1 := Coord{i, 0}
		c2 := Coord{i, cd - 1}
		q2 = append(q2, c1)
		q2 = append(q2, c2)
		newSeen[c1] = true
		newSeen[c2] = true
	}

	for i := 0; i < cd; i++ {
		c1 := Coord{0, i}
		c2 := Coord{rd - 1, 0}
		q2 = append(q2, c1)
		q2 = append(q2, c2)
		newSeen[c1] = true
		newSeen[c2] = true
	}

	vD := 0
	for len(q2) > 0 {
		curr := q2[0]
		q2 = q2[1:]

		if bigDaddy[curr.r][curr.c] == 2 {
			vD++
		}

		ns := make([]Coord, 0)
		if curr.r-1 >= 0 {
			c := Coord{curr.r - 1, curr.c}
			if bigDaddy[c.r][c.c] != 1 {
				ns = append(ns, c)
			}
		}
		if curr.c+1 < cd {
			c := Coord{curr.r, curr.c + 1}
			if bigDaddy[c.r][c.c] != 1 {
				ns = append(ns, c)
			}
		}
		if curr.r+1 < rd {
			c := Coord{curr.r + 1, curr.c}
			if bigDaddy[c.r][c.c] != 1 {
				ns = append(ns, c)
			}
		}
		if curr.c-1 >= 0 {
			c := Coord{curr.r, curr.c - 1}
			if bigDaddy[c.r][c.c] != 1 {
				ns = append(ns, c)
			}
		}

		for _, n := range ns {
			if !newSeen[n] {
				newSeen[n] = true
				q2 = append(q2, n)
			}
		}
	}

	fmt.Println("VD", vD)
	return dc - vD
}

func run(fname string) int {
	liens, i, j := read(fname)
	r := solve2(liens, i, j)
	fmt.Println("res", fname, r)
	return r
}

func main() {
	expected := 8
	rtest := run("inp.test")
	if rtest != expected {
		log.Fatalln("FAILED TEST CASE", rtest, expected)
	}
	rmain := run("inp.main")
	fmt.Println("Result", rmain)
}
