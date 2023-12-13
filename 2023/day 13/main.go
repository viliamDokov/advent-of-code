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

func read(fname string) [][][]bool {
	f, err := os.ReadFile(fname)
	data := string(f)
	check(err)
	reflections := make([][][]bool, 0)
	segments := strings.Split(data, "\r\n\r\n")
	for _, s := range segments {
		reflection := make([][]bool, 0)
		lines := strings.Split(s, "\n")
		for _, l := range lines {
			row := make([]bool, 0)
			l = strings.TrimSpace(l)
			for _, c := range l {
				row = append(row, c == '#')
			}
			reflection = append(reflection, row)
		}
		reflections = append(reflections, reflection)
	}
	return reflections
}

func solve(reflections [][][]bool) int {
	acc := 0
	for _, ref := range reflections {

		orows := mirrorRows(ref)
		trans := transpose(ref)
		ocols := mirrorRows(trans)
		acc += 100*orows + ocols
	}
	return acc
}

func areSame(a []bool, b []bool) int {
	acc := 0
	for i := range a {
		if a[i] != b[i] {
			acc++
		}
	}
	return acc
}

func transpose(slice [][]bool) [][]bool {
	xl := len(slice[0])
	yl := len(slice)
	result := make([][]bool, xl)
	for i := range result {
		result[i] = make([]bool, yl)
	}
	for i := 0; i < xl; i++ {
		for j := 0; j < yl; j++ {
			result[i][j] = slice[j][i]
		}
	}
	return result
}

func mirrorRows(reflection [][]bool) int {
	for i := 0; i < len(reflection)-1; i++ {
		a := reflection[i]
		b := reflection[i+1]

		smudge := false
		diff := areSame(a, b)
		if diff <= 1 {
			smudge = diff == 1

			isMirror := true
			for sep := 1; i-sep >= 0 && i+1+sep < len(reflection); sep++ {
				a := reflection[i-sep]
				b := reflection[i+1+sep]
				diff = areSame(a, b)
				if diff > 1 || (diff == 1 && smudge) {
					isMirror = false
					break
				}
				if diff == 1 {
					smudge = true
				}
			}
			if isMirror && smudge {
				return i + 1 //???
			}
		}
	}

	return 0
}

func run(fname string) int {
	s := time.Now()
	data := read(fname)
	r := solve(data)
	fmt.Println("Took:", time.Since(s))
	return r
}

func main() {
	expected := 400
	rtest := run("inp.test")
	if rtest != expected {
		log.Fatalln("FAILED TEST CASE", rtest, expected)
	}
	rmain := run("inp.main")
	fmt.Println("Result", rmain)
}
