package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	redblacktree "github.com/emirpasic/gods/trees/redblacktree"
)

type Edges map[string]map[string]int
type Nodes map[string]bool
type Graph struct {
	nodes Nodes
	edges Edges
}

func check(err error) {
	if err != nil {
		log.Fatalf("Fatal error %s", err)
	}
}

func read(fname string) Graph {
	f, err := os.ReadFile(fname)
	data := string(f)
	check(err)
	lines := strings.Split(data, "\n")
	edges := make(Edges)
	nodes := make(Nodes)
	for _, l := range lines {
		l = strings.TrimSpace(l)
		tmp := strings.Split(l, ": ")
		nodes[tmp[0]] = true
		_, in := edges[tmp[0]]
		if !in {
			edges[tmp[0]] = make(map[string]int)
		}

		conns := strings.Split(tmp[1], " ")
		for _, e := range conns {
			nodes[e] = true
			_, in := edges[e]
			if !in {
				edges[e] = make(map[string]int)
			}
			edges[e][tmp[0]] = 1
			edges[tmp[0]][e] = 1
		}
	}
	return Graph{nodes, edges}
}

func mergeNodes(graph Graph, n1 string, n2 string) {
	new := n1 + "+" + n2
	delete(graph.nodes, n1)
	delete(graph.nodes, n2)
	graph.nodes[new] = true
	graph.edges[new] = make(map[string]int)
	for other, w := range graph.edges[n1] {
		if other != n2 {
			graph.edges[new][other] += w
			graph.edges[other][new] += w
			delete(graph.edges[other], n1)
		}
	}
	for other, w := range graph.edges[n2] {
		if other != n1 {
			graph.edges[new][other] += w
			graph.edges[other][new] += w
			delete(graph.edges[other], n2)
		}
	}
	delete(graph.edges, n1)
	delete(graph.edges, n2)
}

type Weight struct {
	w    int
	node string
}

func printPq(pq redblacktree.Tree) {
	fmt.Print("[")
	it := pq.Iterator()
	for it.Begin(); it.Next(); {
		n := it.Node()
		fmt.Print(n.Key, " ")
	}
	fmt.Print("]\n")
	// for i,n := range  {

	// }
}

func contract(graph Graph, snode string) (int, int) {
	seen := make(map[string]bool)
	seen[snode] = true
	weights := make(map[string]int)
	pq := redblacktree.NewWith(func(a, b interface{}) int {
		n1 := a.(Weight)
		n2 := b.(Weight)
		if n1.w < n2.w {
			return -1
		}
		if n1.w > n2.w {
			return 1
		}
		if n1.node < n2.node {
			return -1
		}
		if n1.node > n2.node {
			return 1
		}
		return 0
	})

	for node := range graph.nodes {
		if !seen[node] {
			for ndoes := range seen {
				weights[node] += graph.edges[node][ndoes]
			}
		}
	}

	for node := range graph.nodes {
		if !seen[node] {
			pq.Put(Weight{weights[node], node}, nil)
		}
	}

	// printPq(*pq)

	n := len(graph.nodes)
	var last string
	for i := 0; i < n-2; i++ {
		// fmt.Println("contracting", i, pq.Size())
		maxNodeW := pq.Right().Key.(Weight)
		maxNode := maxNodeW.node
		pq.Remove(maxNodeW)
		// fmt.Println(maxNode, weights[maxNode])
		seen[maxNode] = true
		last = maxNode
		neighbours := graph.edges[maxNode]
		for node := range neighbours {
			if !seen[node] {
				pq.Remove(Weight{weights[node], node})
			}
		}
		for node, w := range neighbours {
			weights[node] += w
		}
		for node := range neighbours {
			if !seen[node] {
				pq.Put(Weight{weights[node], node}, nil)
			}
		}
	}
	var unseen string
	for node := range graph.nodes {
		if !seen[node] {
			unseen = node
		}
	}
	cutcost := 0
	for node := range seen {
		cutcost += graph.edges[unseen][node]
	}
	cutSize := len(strings.Split(unseen, "+"))

	mergeNodes(graph, last, unseen)
	return cutcost, cutSize
}

func solve1(graph Graph) int {
	mincut := 100000000000
	n := len(graph.nodes)
	for i := 0; i < n-1; i++ {
		fmt.Println(i, n)
		cutw, cutsize := contract(graph, "rsh")
		if cutw == 3 {
			return cutsize * (n - cutsize)
		}
		if mincut > cutw {
			mincut = cutw
		}
	}
	return mincut
}

func solve2(graph Graph) int {
	return 69
}

func run(fname string) int {
	g := read(fname)
	fmt.Println(g)
	r := solve1(g)
	return r
}
func main() {
	expected := 54
	rtest := run("inp.test")
	if rtest != expected {
		log.Fatalln("FAILED TEST CASE", rtest, expected)
	}
	rmain := run("inp.main")
	fmt.Println("Result", rmain)
}
