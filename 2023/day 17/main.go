package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func check(err error) {
	if err != nil {
		log.Fatalf("Fatal error %s", err)
	}
}

func read(fname string) []string {
	f, err := os.ReadFile(fname)
	data := string(f)
	check(err)
	lines := strings.Split(data, "\n")
	for _, l := range lines {
		fmt.Println(l)
	}
	return lines
}

func solve() int {
	return 69
}

func run(fname string) int {
	data := read(fname)
	return 69
}
func main() {
	expected := 0
	rtest := run("inp.test")
	if rtest != expected {
		log.Fatalln("FAILED TEST CASE", rtest, expected)
	}
	rmain := run("inp.main")
	fmt.Println("Result", rmain)
}
