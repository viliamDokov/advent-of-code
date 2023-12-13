package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strings"
	"time"
)

type Node struct {
	name  string
	left  string
	right string
}

func check(err error) {
	if err != nil {
		log.Fatalf("Fatal error %s", err)
	}
}

func read(fname string) (string, map[string]Node) {
	f, err := os.ReadFile(fname)
	data := string(f)
	check(err)
	segments := strings.Split(data, "\r\n\r\n")

	directions := strings.TrimSpace(segments[0])
	lines := strings.Split(segments[1], "\n")

	nodes := make(map[string]Node)

	for _, l := range lines {
		tmp := strings.Split(l, " = ")
		name := strings.TrimSpace(tmp[0])
		tmp2 := strings.Split(tmp[1][1:len(tmp[1])-2], ",")
		left := strings.TrimSpace(tmp2[0])
		right := strings.TrimSpace(tmp2[1])
		nodes[name] = Node{name, left, right}
	}

	return directions, nodes
}

func solve(dirs string, nodes map[string]Node) int {

	i := 0
	node := nodes["AAA"]
	s := 0
	for node.name != "ZZZ" {
		d := dirs[i]
		switch d {
		case 'R':
			fmt.Println("node", node, "right")
			node = nodes[node.right]
		case 'L':
			fmt.Println("node", node, "left")
			node = nodes[node.left]
		}

		i += 1
		if i >= len(dirs) {
			i = 0
		}
		s++
	}
	return s
}

type Step struct {
	node string
	step int64
}

type StepData struct {
	bias  int64
	cycle int64
}

type StepThing struct {
	c   int64
	add int64
}

func done2(seen map[Step]*StepData) bool {
	for _, v := range seen {
		if v.cycle <= 0 {
			return false
		}
	}
	return len(seen) > 0
}

func getNext(things []StepThing) int64 {
	min := int64(math.MaxInt64)
	t := 0
	for i, e := range things {
		if min > e.c {
			t = i
			min = e.c
		}
	}
	res := things[t].c
	things[t].c += things[t].add
	return res
}

func allHold(i int64, states [][]StepData) bool {
	// fmt.Println("start", i, states)
	for _, s := range states {
		holds := false
		for _, d := range s {
			// fmt.Printf("i:%d, b:%d, c:%d r:%d \n", i, d.bias,d.cycle)
			// fmt.Println("vibe: ", i-d.bias, d.cycle, (i-d.bias)%d.cycle)
			if (i-d.bias)%d.cycle == 0 {
				holds = true
				break
			}
		}
		if !holds {
			return false
		}
	}
	return true
}

func solve3(dirs string, nodes map[string]Node) int64 {

	cnodes := make([]Node, 0)
	for _, n := range nodes {
		if strings.HasSuffix(n.name, "A") {
			cnodes = append(cnodes, n)
		}
	}
	fmt.Println(cnodes)

	datas := make([][]StepData, 0)
	things := make([]StepThing, 0)

	for _, n := range cnodes {
		seen := make(map[Step]*StepData)
		i := int64(0)
		s := int64(0)
		curr := n

		for !done2(seen) {
			d := dirs[i]
			n := Node{}
			switch d {
			case 'R':
				if nodes[curr.right] == n {
					log.Fatalln(curr, n)
				}
				curr = nodes[curr.right]
			case 'L':
				if nodes[curr.right] == n {
					log.Fatalln(curr, n)
				}
				curr = nodes[curr.left]
			}

			if strings.HasSuffix(curr.name, "Z") {
				if seen[Step{curr.name, i}] == nil {
					seen[Step{curr.name, i}] = &StepData{s, 0}
				} else {
					seen[Step{curr.name, i}].cycle = s - seen[Step{curr.name, i}].bias
				}
			}

			i += 1
			if i >= int64(len(dirs)) {
				i = 0
			}
			s++
		}

		fmt.Println("end")
		sdatas := make([]StepData, 0)
		for _, v := range seen {
			sdatas = append(sdatas, *v)
			things = append(things, StepThing{v.bias, v.cycle})
		}
		datas = append(datas, sdatas)
	}

	fmt.Println("heap", things)
	fmt.Println("data", datas)
	i := int64(0)
	step := 1_000_000_000_000
	c := 0
	s := time.Now()
	it := 0
	for !allHold(i, datas) {
		if i > int64(c*step) {
			c++
			fmt.Println("thing", c, time.Since(s), it)
			s = time.Now()
			it = 0
		}
		it++
		i = getNext(things)
	}

	return int64(i + 1)
}

func run(fname string) int64 {
	dirs, ns := read(fname)
	fmt.Println("read", dirs, ns)
	r := solve3(dirs, ns)
	return r
}

func main() {
	expected := int64(6)
	rtest := run("inp.test")
	if rtest != expected {
		log.Fatalln("FAILED TEST CASE", rtest, expected)
	}

	rmain := run("inp.main")
	fmt.Println("Result", rmain)
}
