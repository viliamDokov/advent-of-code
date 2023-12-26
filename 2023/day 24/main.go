package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Point [3]float64
type Hail struct {
	p Point
	v Point
}

func check(err error) {
	if err != nil {
		log.Fatalf("Fatal error %s", err)
	}
}

func read(fname string) []Hail {
	f, err := os.ReadFile(fname)
	data := string(f)
	check(err)
	lines := strings.Split(data, "\n")
	hails := make([]Hail, 0)
	for _, l := range lines {
		l := strings.TrimSpace(l)
		tmp := strings.Split(l, " @ ")
		var p Point
		for i, nms := range strings.Split(tmp[0], ", ") {
			nms = strings.TrimSpace(nms)
			n, err := strconv.ParseFloat(nms, 64)
			check(err)
			p[i] = n
		}
		var v Point
		for i, nms := range strings.Split(tmp[1], ", ") {
			nms = strings.TrimSpace(nms)
			n, err := strconv.ParseFloat(nms, 64)
			check(err)
			v[i] = n
		}
		hails = append(hails, Hail{p, v})
	}
	return hails
}

func solve(hails []Hail, lb float64, ub float64) int {
	acc := 0
	for i := range hails {
		for j := range hails {
			if i == j {
				continue
			}

			h1 := hails[i]
			h2 := hails[j]
			// fmt.Println(h1, h2)

			if h1.v[1]*h2.v[0]-h1.v[0]*h2.v[1] == 0 {
				// fmt.Println("Parallel!")
				continue
			}
			t1 := ((h2.p[1]-h1.p[1])*h2.v[0] + (h1.p[0]-h2.p[0])*h2.v[1]) / (h1.v[1]*h2.v[0] - h1.v[0]*h2.v[1])
			t2 := ((h1.p[1]-h2.p[1])*h1.v[0] + (h2.p[0]-h1.p[0])*h1.v[1]) / (h2.v[1]*h1.v[0] - h2.v[0]*h1.v[1])

			if t1 < 0 || t2 < 0 {
				// fmt.Println("Past")
				continue
			}

			x1 := h1.p[0] + t1*h1.v[0]
			y1 := h1.p[1] + t1*h1.v[1]

			if x1 < lb || x1 > ub || y1 < lb || y1 > ub {
				// fmt.Println("Out of bounds")
				continue
			}

			// fmt.Println("sucesss")
			acc++
		}
	}
	return acc / 2
}

func solve2(hails []Hail, lb float64, ub float64) int {
	acc := 0
	for i := range hails {
		for j := range hails {
			if i == j {
				continue
			}

			h1 := hails[i]
			h2 := hails[j]
			// fmt.Println(h1, h2)

			if h1.v[1]*h2.v[0]-h1.v[0]*h2.v[1] == 0 {
				// fmt.Println("Parallel!")
				continue
			}
			t1 := ((h2.p[1]-h1.p[1])*h2.v[0] + (h1.p[0]-h2.p[0])*h2.v[1]) / (h1.v[1]*h2.v[0] - h1.v[0]*h2.v[1])
			t2 := ((h1.p[1]-h2.p[1])*h1.v[0] + (h2.p[0]-h1.p[0])*h1.v[1]) / (h2.v[1]*h1.v[0] - h2.v[0]*h1.v[1])

			if t1 < 0 || t2 < 0 {
				// fmt.Println("Past")
				continue
			}

			x1 := h1.p[0] + t1*h1.v[0]
			y1 := h1.p[1] + t1*h1.v[1]

			if x1 < lb || x1 > ub || y1 < lb || y1 > ub {
				// fmt.Println("Out of bounds")
				continue
			}

			// fmt.Println("sucesss")
			acc++
		}
	}
	return acc / 2
}

func run(fname string) int {
	hails := read(fname)
	fmt.Println(hails)
	var lb float64
	var ub float64
	if strings.HasSuffix(fname, ".test") {
		lb = 7
		ub = 27
	} else {
		lb = 200000000000000
		ub = 400000000000000
	}
	r := solve(hails, lb, ub)
	return r
}

func main() {
	expected := 2
	rtest := run("inp.test")
	if rtest != expected {
		log.Fatalln("FAILED TEST CASE", rtest, expected)
	}
	rmain := run("inp.main")
	fmt.Println("Result", rmain)
}
