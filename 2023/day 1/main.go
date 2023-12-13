package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func check(err error) {
	if err != nil {
		log.Fatalf("error %s", err)
	}
}

func read() []string {
	f, err := os.Open("inp")
	check(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	lines := make([]string, 0, 0)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	return lines
}

var digits = map[string]int{"one": 1, "two": 2, "three": 3, "four": 4, "five": 5, "six": 6, "seven": 7, "eight": 8, "nine": 9}

func getDigit(line string, i int) int {
	d, derr := strconv.Atoi(line[i : i+1])
	if derr == nil {
		return d
	}

	for k, v := range digits {
		if strings.HasPrefix(line[i:], k) {
			return v
		}
	}

	return -1
}

func solve(lines []string) int {
	s := 0
	for _, line := range lines {
		n := len(line)
		first := -1
		last := -1
		for j := range line {
			d1 := getDigit(line, j)
			d2 := getDigit(line, n-j-1)
			if d1 > 0 && first < 0 {
				first = d1
			}
			if d2 > 0 && last < 0 {
				last = d2
			}

		}

		fmt.Println(first, last)
		s += first*10 + last
	}
	return s
}

func main() {
	lines := read()
	res := solve(lines)
	fmt.Println(res)
}
