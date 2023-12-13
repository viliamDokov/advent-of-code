package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

type Coord struct {
	x int
	y int
}

func check(err error) {
	if err != nil {
		log.Fatalf("Fatal error %s", err)
	}
}

func read(fname string) [][]byte {
	f, err := os.ReadFile(fname)
	data := string(f)
	check(err)
	lines := strings.Split(data, "\n")
	entries := make([][]byte, 0)
	for _, l := range lines {
		fmt.Println(l)
		ls := strings.TrimSpace(l)
		charRow := make([]byte, 0)
		for i := range ls {
			charRow = append(charRow, ls[i])
		}
		entries = append(entries, charRow)
	}
	return entries
}

func solve(data [][]byte) int {
	emptyRows := make([]int, 0)

	for i, row := range data {
		isEmpty := true
		for _, c := range row {
			if c == '#' {
				isEmpty = false
			}
		}
		if isEmpty {
			emptyRows = append(emptyRows, i)
		}
	}

	emptyCols := make([]int, 0)

	for j := range data[0] {
		isEmpty := true
		for i := range data {
			c := data[i][j]
			if c == '#' {
				isEmpty = false
			}
		}
		if isEmpty {
			emptyCols = append(emptyCols, j)
		}
	}

	// fmt.Println(emptyRows, emptyCols)

	// for i, er := range emptyRows {
	// 	emptyRow := make([]byte, 0)
	// 	for i := 0; i < len(data[0]); i++ {
	// 		emptyRow = append(emptyRow, '.')
	// 	}
	// 	tmp := append(data[:er+i], emptyRow)
	// 	data = append(tmp, data[er+i:]...)
	// }

	// for j, er := range emptyCols {
	// 	for i := 0; i < len(data); i++ {
	// 		// fmt.Println(i, er+j)
	// 		tmp := append(data[i][:er+j], '.')
	// 		data[i] = append(tmp, data[i][er+j:]...)
	// 	}
	// }

	galaxies := make([]Coord, 0)
	erp := 0
	expandFactor := 1000000 - 1
	for i := range data {
		if erp < len(emptyRows) {
			// fmt.Println("AAA", i, erp, emptyRows[erp])
			if emptyRows[erp] < i {
				erp++
			}
		}

		ecp := 0
		for j := range data[i] {
			if ecp < len(emptyCols) {
				if emptyCols[ecp] < j {
					ecp++
				}
			}

			c := data[i][j]
			if c == '#' {
				fmt.Println(erp, ecp)
				galaxies = append(galaxies, Coord{i + erp*expandFactor, j + ecp*expandFactor})
			}
		}
	}

	acc := 0
	for i := range galaxies {
		for j := i + 1; j < len(galaxies); j++ {
			c1 := galaxies[i]
			c2 := galaxies[j]
			dist := int(math.Abs(float64(c1.x-c2.x))) + int(math.Abs(float64(c1.y-c2.y)))
			acc += dist
			// fmt.Println(dist, acc)
		}
	}

	return acc
}

func run(fname string) int {
	data := read(fname)
	r := solve(data)
	return r
}
func main() {
	// expected := 8410
	// rtest := run("inp.test")
	// if rtest != expected {
	// 	log.Fatalln("FAILED TEST CASE", rtest, expected)
	// }
	rmain := run("inp.main")
	fmt.Println("Result", rmain)
}
