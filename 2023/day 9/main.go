package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

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
	nums := make([][]int, 0)

	for _, l := range lines {
		if strings.TrimSpace(l) == "" {
			continue
		}
		row := make([]int, 0)
		for _, ln := range strings.Split(strings.TrimSpace(l), " ") {
			n, err := strconv.Atoi(ln)
			check(err)
			row = append(row, n)
		}
		nums = append(nums, row)
	}
	return nums
}

func done(subseqs [][]int) bool {
	sq := subseqs[len(subseqs)-1]
	for _, v := range sq {
		if v != 0 {
			return false
		}
	}
	return true
}

func solve(seqs [][]int) int {
	acc := 0
	for _, seq := range seqs {
		subseqs := make([][]int, 0)
		subseqs = append(subseqs, seq)
		for !done(subseqs) {
			n := len(subseqs)
			subseq := make([]int, 0)
			for i := 1; i < len(subseqs[n-1]); i++ {
				subseq = append(subseq, subseqs[n-1][i]-subseqs[n-1][i-1])
			}
			subseqs = append(subseqs, subseq)
		}

		curr := 0
		for i := (len(subseqs) - 2); i >= 0; i-- {
			sq := subseqs[i]
			first := sq[0]
			curr = first - curr
		}
		fmt.Println(seq, curr)
		acc += curr
	}

	return acc
}

func run(fname string) int {
	data := read(fname)
	fmt.Println(data)
	r := solve(data)
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
