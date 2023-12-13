package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func check(err error) {
	if err != nil {
		log.Fatalf("Fatal error %s", err)
	}
}

func calc(t int, d int) int {
	acc := 0
	for i := 1; i < t; i++ {
		rem := t - i
		md := i * rem
		if md > d {
			acc++
		}
	}

	return acc
}

func solve(times []int, dists []int) int {
	res := 1
	for i := range times {
		r := calc(times[i], dists[i])
		fmt.Println(r)
		res = res * r
	}
	return res
}

func read() ([]int, []int) {
	f, err := os.ReadFile("inp")
	data := string(f)
	check(err)
	lines := strings.Split(data, "\n")

	tmp := strings.Split(lines[0], ":")
	times := make([]int, 0)
	for _, s := range strings.Split(tmp[1], " ") {
		s = strings.TrimSpace(s)
		if s != "" {
			n, err := strconv.Atoi(s)
			check(err)
			times = append(times, n)
		}
	}

	tmp2 := strings.Split(lines[1], ":")
	dists := make([]int, 0)
	for _, s := range strings.Split(tmp2[1], " ") {
		s = strings.TrimSpace(s)
		if s != "" {
			n, err := strconv.Atoi(s)
			check(err)
			dists = append(dists, n)
		}
	}
	return times, dists

}

func read2() (int, int) {
	f, err := os.ReadFile("inp")
	data := string(f)
	check(err)
	lines := strings.Split(data, "\n")

	tmp := strings.Split(lines[0], ":")
	s := strings.TrimSpace(strings.ReplaceAll(tmp[1], " ", ""))
	t, err := strconv.Atoi(s)
	check(err)

	tmp2 := strings.Split(lines[1], ":")
	s = strings.TrimSpace(strings.ReplaceAll(tmp2[1], " ", ""))
	d, err := strconv.Atoi(s)
	check(err)

	return t, d

}

func main() {
	start := time.Now()
	ts, ds := read()
	fmt.Println(ts, ds)
	r := solve(ts, ds)
	fmt.Println(r)
	fmt.Println("Part1", time.Since(start))

	start = time.Now()
	t, d := read2()
	fmt.Println(t, d)
	r = calc(t, d)
	fmt.Println(r)
	fmt.Println("Part2", time.Since(start))

}
