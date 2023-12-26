package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

const (
	BROADCAST int = iota
	FLIP_FLOP
	CONJ
)

type Node struct {
	name        string
	t           int
	ffstate     bool
	connections []string
	prevs       map[string]bool
}

func check(err error) {
	if err != nil {
		log.Fatalf("Fatal error %s", err)
	}
}

func read(fname string) map[string]*Node {
	f, err := os.ReadFile(fname)
	data := string(f)
	check(err)
	lines := strings.Split(data, "\n")
	nodes := make(map[string]*Node, 0)
	for _, l := range lines {
		l = strings.TrimSpace(l)
		spl := strings.Split(l, " -> ")
		node := Node{}
		if spl[0] == "broadcaster" {
			node.name = spl[0]
			node.t = BROADCAST
		} else if spl[0][0] == '%' {
			node.name = spl[0][1:]
			node.t = FLIP_FLOP
		} else if spl[0][0] == '&' {
			node.name = spl[0][1:]
			node.t = CONJ
			node.prevs = make(map[string]bool)
		} else {
			log.Fatalln("Never happen")
		}

		node.connections = make([]string, 0)
		nbs := strings.Split(spl[1], ",")
		for _, nb := range nbs {
			nb := strings.TrimSpace(nb)
			node.connections = append(node.connections, nb)
		}
		nodes[node.name] = &node
	}

	for _, node := range nodes {
		for _, nname := range node.connections {
			nn := nodes[nname]
			if nn != nil && nn.t == CONJ {
				nn.prevs[node.name] = false
			}
		}
	}

	// fmt.Println(*nodes["inv"])
	return nodes
}

const LOW = false
const HIHG = true

type Signal struct {
	source string
	dest   string
	pulse  bool
}

func solve(nodes map[string]*Node) int {

	// ITERATIONS := 1000

	lows := 0
	highs := 0
	i := 0
	s := time.Now()
	// prev := 0
	for {
		rx := false
		q := make([]Signal, 0)
		q = append(q, Signal{"button", "broadcaster", LOW})

		i++
		if i%1000000 == 0 {
			fmt.Println(i, time.Since(s))
			s = time.Now()
		}
		for len(q) > 0 {
			signal := q[0]
			q = q[1:]
			node := nodes[signal.dest]
			// fmt.Printf("%s - %t -> %s\n", signal.source, signal.pulse, signal.dest)
			if signal.pulse {
				highs++
			} else {
				lows++
				if signal.dest == "zc" {
					fmt.Println(i)
				}
			}
			if node == nil {
				continue
			}
			switch node.t {
			case BROADCAST:
				for _, n := range node.connections {
					q = append(q, Signal{node.name, n, signal.pulse})
				}
			case FLIP_FLOP:
				if signal.pulse == LOW {
					node.ffstate = !node.ffstate
					for _, n := range node.connections {
						q = append(q, Signal{node.name, n, node.ffstate})
					}
				}
			case CONJ:
				node.prevs[signal.source] = signal.pulse
				s := 0
				for _, v := range node.prevs {
					if v {
						s++
					}
				}
				for _, n := range node.connections {
					q = append(q, Signal{node.name, n, s != len(node.prevs)})
				}
			default:
				log.Fatalln("Never happen!")
			}
		}
		if rx {
			break
		}
	}
	return i
}

func run(fname string) int {
	nodes := read(fname)
	fmt.Println(nodes)
	r := solve(nodes)
	return r
}
func main() {
	// expected := 32000000
	// rtest := run("inp.test")
	// if rtest != expected {
	// 	log.Fatalln("FAILED TEST CASE", rtest, expected)
	// }
	rmain := run("inp.main")
	fmt.Println("Result", rmain)
}
