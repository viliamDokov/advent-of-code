package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type Vavle struct {
	name string
	flow int
	next []string
}

type Step struct {
	startMe string
	startEl string
	timeMe  int
	timeEl  int
	valves  []string
	flow    int
	time    int
}

func getDistMatrix(valves map[string]Vavle) map[string]map[string]int {
	dists := make(map[string]map[string]int)

	for name, valve := range valves {
		dists[name] = make(map[string]int)
		dists[name][name] = 0
		q := make([]Vavle, len(valves))
		seen := make(map[string]bool, len(valves))
		seen[name] = true
		q = append(q, valve)
		for len(q) > 0 {
			var curr Vavle
			curr, q = q[0], q[1:]
			for _, nb := range curr.next {
				if !seen[nb] {
					seen[nb] = true
					q = append(q, valves[nb])
					dists[valve.name][nb] = dists[valve.name][curr.name] + 1
				}
			}
		}
	}

	return dists
}

func getFlow(valves map[string]Vavle, dists map[string]map[string]int, start Vavle, time int) int {

	start_valves := make([]string, len(valves))
	i := 0
	for k := range valves {
		start_valves[i] = k
		i++
	}

	// fmt.Println(start_valves)
	stack := make([]Step, 100)
	first := Step{startMe: start.name, startEl: start.name, timeMe: time, timeEl: time, valves: start_valves, flow: 0, time: time}
	stack = append(stack, first)
	total_flow := 0
	for _, valve := range valves {
		total_flow += valve.flow
	}

	best := 0

	for len(stack) > 0 {
		var curr Step
		curr, stack = stack[len(stack)-1], stack[:len(stack)-1]
		me := curr.timeMe == curr.time
		for _, name := range curr.valves {

			var start string
			if !me {
				start = curr.startEl
			} else {
				start = curr.startMe
			}

			valve := valves[name]
			new_time := curr.time - 1 - dists[start][name]

			if new_time < 0 && (!me || curr.timeEl < 0) && (me || curr.timeMe < 0) {
				continue
			}
			new_flow := curr.flow + valve.flow*new_time
			if new_flow > best {
				best = new_flow
			}

			new_valves := make([]string, len(curr.valves)-1)
			i := 0
			for _, v := range curr.valves {
				if v != name {
					new_valves[i] = v
					i++
				}
			}

			var new_step Step
			if !me {
				max_time := MaxInts(curr.timeMe, new_time)
				new_step = Step{startMe: curr.startMe, startEl: name, timeMe: curr.timeMe, timeEl: new_time, valves: new_valves, flow: new_flow, time: max_time}
			} else {
				max_time := MaxInts(curr.timeEl, new_time)
				new_step = Step{startMe: name, startEl: curr.startEl, timeMe: new_time, timeEl: curr.timeEl, valves: new_valves, flow: new_flow, time: max_time}
			}

			stack = append(stack, new_step)
		}
	}

	return best
}

func MaxInts(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func main() {
	f, err := os.Open("inp")
	if err != nil {
		log.Fatal(err)
	}

	var valves map[string]Vavle
	var nonEmpty map[string]Vavle

	valves = make(map[string]Vavle)
	nonEmpty = make(map[string]Vavle)

	defer f.Close()
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		tmp := strings.Split(strings.TrimSpace(line), "; ")
		var valve Vavle
		valve.name = strings.Split(tmp[0][6:], " has flow rate=")[0]
		valve.flow, err = strconv.Atoi(strings.Split(tmp[0][6:], " has flow rate=")[1])
		if err != nil {
			log.Fatal(err)
		}
		valve.next = make([]string, 0)
		tmp2 := strings.Split(tmp[1][22:], ", ")
		valve.next = append(valve.next, tmp2...)

		valves[valve.name] = valve
		if valve.flow > 0 {
			nonEmpty[valve.name] = valve
		}

	}

	dists := getDistMatrix(valves)
	fmt.Println(dists)
	t1 := time.Now()
	r := getFlow(nonEmpty, dists, valves["AA"], 26)
	fmt.Printf("flow took %v\n", time.Since(t1))
	fmt.Printf("Result %d", r)
}
