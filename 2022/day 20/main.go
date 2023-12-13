package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	prev    *Node
	next    *Node
	val     int64
	isStart bool
}

func moveBack(a *Node) {
	if a.isStart {
		a.isStart = false
		a.prev.isStart = true
	} else if a.prev.isStart {
		a.isStart = true
		a.prev.isStart = false
	}

	a.prev.prev.next = a
	a.next.prev = a.prev

	a.prev.next = a.next
	a.next = a.prev

	a.prev = a.next.prev
	a.next.prev = a
}

func moveForward(a *Node) {
	if a.isStart {
		a.next.isStart = true
		a.isStart = false
	} else if a.next.isStart {
		a.next.isStart = false
		a.isStart = true
	}

	a.next.next.prev = a
	a.prev.next = a.next

	a.next.prev = a.prev
	a.prev = a.next

	a.next = a.prev.next
	a.prev.next = a
}

func printNodes(n *Node) {
	fmt.Println("---------")

	curr := n
	for !curr.isStart {
		curr = curr.next
	}

	fmt.Print(curr.val, " ")
	curr = curr.next
	for !curr.isStart {
		fmt.Print(curr.val, " ")
		curr = curr.next
	}
	fmt.Println("\n---------")

}

func readInput() []Node {
	f, err := os.Open("inp")
	defer func() {
		err = f.Close()
		logErr(err)
	}()

	nodes := make([]Node, 0, 5000)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		v, err := strconv.Atoi(strings.TrimSpace(line))
		v64 := int64(v) * 811589153
		logErr(err)
		node := Node{nil, nil, v64, false}
		nodes = append(nodes, node)
		if len(nodes) > 1 {
			nodes[len(nodes)-1].prev = &nodes[len(nodes)-2]
			nodes[len(nodes)-2].next = &nodes[len(nodes)-1]
		}
	}

	nodes[0].prev = &nodes[len(nodes)-1]
	nodes[len(nodes)-1].next = &nodes[0]
	nodes[0].isStart = true
	return nodes
}

func iAbs(a int64) int64 {
	if a < 0 {
		return -a
	}
	return a
}

func scramble(nodes []Node) {
	for i := range nodes {
		curr := &nodes[i]
		// fmt.Println("Moving", curr, iAbs(curr.val))
		next := curr.next
		for j := 0; j < int(iAbs(curr.val)%int64((len(nodes)-1))); j++ {
			// fmt.Println("thing", iAbs(curr.val)%len(nodes))
			// fmt.Println(curr.val)
			if curr.val > 0 {
				moveForward(curr)
			} else {

				// fmt.Println("Moving back")
				moveBack(curr)
			}
			// printNodes(curr)
		}
		curr = next
	}
}

func getResult(n *Node, l int) int64 {
	curr := n
	for curr.val != 0 {
		curr = curr.next
	}
	i := 0
	curr = curr.next
	acc := int64(0)
	for true {
		i++
		if i == 1000 {
			fmt.Println(curr.val)
			acc += curr.val
		}

		if i == 2000 {
			fmt.Println(curr.val)
			acc += curr.val
		}

		if i == 3000 {
			fmt.Println(curr.val)
			acc += curr.val
			break
		}
		curr = curr.next
	}
	return acc
}

func main() {
	nodes := readInput()
	start := &nodes[0]
	printNodes(start)
	for i := 0; i < 10; i++ {
		scramble(nodes)
	}

	r := getResult(start, len(nodes))
	fmt.Println("REs", r)

}

func logErr(err error) {
	if err != nil {
		log.Fatal(err)
	}

}
