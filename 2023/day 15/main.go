package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Instruction struct {
	label string
	op    string
	lens  int
}

func check(err error) {
	if err != nil {
		log.Fatalf("Fatal error %s", err)
	}
}

func read(fname string) []Instruction {
	f, err := os.ReadFile(fname)
	check(err)

	data := string(f)
	hashes := strings.Split(data, ",")
	instructions := make([]Instruction, 0)
	for _, h := range hashes {
		tmp := strings.Split(h, "-")
		if len(tmp) > 1 {
			label := tmp[0]
			op := "-"
			instructions = append(instructions, Instruction{label, op, 0})
			continue
		}
		tmp = strings.Split(h, "=")
		if len(tmp) > 1 {
			label := tmp[0]
			op := "="
			lens, err := strconv.Atoi(tmp[1])
			check(err)
			instructions = append(instructions, Instruction{label, op, lens})
			continue
		}
		log.Fatalln("NEver", tmp, len(tmp), h)
	}

	return instructions
}

func hash(s string) int {
	v := 0
	for _, c := range s {
		ascii := int(c)
		v += ascii
		v = v * 17
		v = v % 256
	}
	return v
}

func solve(instructions []Instruction) int {
	boxes := make([][]Instruction, 256, 256)
	for i := 0; i < len(boxes); i++ {
		boxes[i] = make([]Instruction, 0)
	}
	for i := range instructions {
		inst := instructions[i]
		fmt.Println(inst)
		switch inst.op {
		case "-":
			boxi := hash(inst.label)
			for j := range boxes[boxi] {
				if boxes[boxi][j].label == inst.label {
					boxes[boxi] = append(boxes[boxi][:j], boxes[boxi][j+1:]...)
					break
				}
			}
		case "=":
			boxi := hash(inst.label)
			added := false
			for j := range boxes[boxi] {
				if boxes[boxi][j].label == inst.label {
					boxes[boxi][j].lens = inst.lens
					added = true
					break
				}
			}
			if !added {
				fmt.Println("Adding!")
				boxes[boxi] = append(boxes[boxi], inst)
			}

		}
	}
	fmt.Println(boxes)
	acc := 0
	for i := range boxes {
		box := boxes[i]
		for j := range box {
			lens := box[j]
			power := (i + 1) * (j + 1) * lens.lens
			acc += power
		}
	}

	return acc
}

func run(fname string) int {
	data := read(fname)
	fmt.Println("res")
	r := solve(data)
	return r
}
func main() {
	expected := 145
	rtest := run("inp.test")
	if rtest != expected {
		log.Fatalln("FAILED TEST CASE", rtest, expected)
	}
	fmt.Println("Passed test!")
	rmain := run("inp.main")
	fmt.Println("Result", rmain)
}
