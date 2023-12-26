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

type Rule struct {
	attribute  string
	comparison string
	value      int
	target     string
}

func read(fname string) (map[string][]Rule, []map[string]int) {
	f, err := os.ReadFile(fname)
	data := string(f)
	check(err)
	segments := strings.Split(data, "\r\n\r\n")
	ruless := segments[0]
	partss := segments[1]

	rules := make(map[string][]Rule)
	parts := make([]map[string]int, 0)

	rlines := strings.Split(ruless, "\n")
	for _, rline := range rlines {
		rline = strings.TrimSpace(rline)
		tmp := strings.Split(rline, "{")
		name := tmp[0]
		rule := make([]Rule, 0)
		conditions := strings.Split(tmp[1], ",")
		for i := 0; i < len(conditions)-1; i++ {
			cond := conditions[i]
			attr, sign, rest := cond[0:1], cond[1:2], cond[2:]
			tmp2 := strings.Split(rest, ":")
			value, err := strconv.Atoi(tmp2[0])
			check(err)
			target := tmp2[1]
			rule = append(rule, Rule{attr, sign, value, target})
		}
		lastcond := conditions[len(conditions)-1]
		target := lastcond[:len(lastcond)-1]
		rule = append(rule, Rule{"69", "69", 69, target})
		rules[name] = rule
	}

	plines := strings.Split(partss, "\n")
	for _, pline := range plines {
		part := make(map[string]int)
		pline = strings.TrimSpace(pline)
		pline = pline[1 : len(pline)-1]
		tmp := strings.Split(pline, ",")
		for _, prop := range tmp {
			tmp2 := strings.Split(prop, "=")
			val, err := strconv.Atoi(tmp2[1])
			check(err)
			part[tmp2[0]] = val
		}
		parts = append(parts, part)
	}
	return rules, parts
}

func solve(workflows map[string][]Rule, parts []map[string]int) int {
	acc := 0
	for _, part := range parts {
		fmt.Println(part)
		wfname := "in"
		for wfname != "R" && wfname != "A" {
			fmt.Println("At workflow", wfname)
			rules := workflows[wfname]
			for _, rule := range rules {
				fmt.Println("Doing rule", rule)
				br := false
				switch rule.comparison {
				case "69":

					wfname = rule.target
					fmt.Println("breaking")
					br = true
				case "<":
					if part[rule.attribute] < rule.value {
						fmt.Println("breaking")
						wfname = rule.target
						br = true

					}
				case ">":
					if part[rule.attribute] > rule.value {
						fmt.Println("breaking")
						wfname = rule.target
						br = true

					}
				default:
					log.Fatalln("Never happen!")
				}

				if br {
					break
				}
				fmt.Println("continuing")
			}
		}
		fmt.Println(wfname, part)
		if wfname == "A" {
			for _, v := range part {
				acc += v
			}
		}
	}
	return acc
}

type Range struct {
	start int
	end   int
}
type SplitStep struct {
	wfname string
	idx    int
	ranges map[string]Range
}

var attributes = []string{"x", "m", "a", "s"}

func GetRangesCount(ranges map[string]Range) int64 {
	acc := int64(1)
	for _, rang := range ranges {
		acc = acc * int64(rang.end-rang.start)
	}
	return acc
}

func copyMap(ranges map[string]Range) map[string]Range {
	r := make(map[string]Range)
	for k, v := range ranges {
		r[k] = v
	}
	return r
}

func solve2(workflows map[string][]Rule) int64 {
	startRanges := make(map[string]Range)
	for _, attr := range attributes {
		startRanges[attr] = Range{1, 4001}
	}
	q := make([]SplitStep, 0)
	q = append(q, SplitStep{"in", 0, startRanges})
	acc := int64(0)
	for len(q) > 0 {
		curr := q[0]
		q = q[1:]
		fmt.Println(curr)
		if curr.wfname == "R" {
			continue
		}
		if curr.wfname == "A" {
			acc += GetRangesCount(curr.ranges)
			continue
		}
		rule := workflows[curr.wfname][curr.idx]
		switch rule.comparison {
		case "69":
			q = append(q, SplitStep{rule.target, 0, curr.ranges})
		case "<":
			rng := curr.ranges[rule.attribute]
			if rule.value <= rng.start {
				reject := copyMap(curr.ranges)
				reject[rule.attribute] = rng
				q = append(q, SplitStep{curr.wfname, curr.idx + 1, reject})
			} else if rule.value < rng.end {
				reject := copyMap(curr.ranges)
				accept := copyMap(curr.ranges)
				reject[rule.attribute] = Range{rule.value, rng.end}
				accept[rule.attribute] = Range{rng.start, rule.value}
				q = append(q, SplitStep{curr.wfname, curr.idx + 1, reject})
				q = append(q, SplitStep{rule.target, 0, accept})
			} else {
				accept := copyMap(curr.ranges)
				accept[rule.attribute] = Range{rng.start, rng.end}
				q = append(q, SplitStep{rule.target, 0, accept})
			}
		case ">":
			rng := curr.ranges[rule.attribute]
			if rule.value < rng.start {
				accept := copyMap(curr.ranges)
				accept[rule.attribute] = Range{rng.start, rng.end}
				q = append(q, SplitStep{rule.target, 0, accept})
			} else if rule.value < rng.end-1 {
				reject := copyMap(curr.ranges)
				accept := copyMap(curr.ranges)
				reject[rule.attribute] = Range{rng.start, rule.value + 1}
				accept[rule.attribute] = Range{rule.value + 1, rng.end}
				q = append(q, SplitStep{curr.wfname, curr.idx + 1, reject})
				q = append(q, SplitStep{rule.target, 0, accept})
			} else {
				reject := copyMap(curr.ranges)
				reject[rule.attribute] = rng
				q = append(q, SplitStep{curr.wfname, curr.idx + 1, reject})
			}
		default:
			log.Fatalln("never happen!")
		}
	}
	return acc
}

func run(fname string) int64 {
	rules, parts := read(fname)
	fmt.Println(rules)
	fmt.Println(parts)
	r := solve2(rules)
	return r
}
func main() {
	expected := int64(167409079868000)
	rtest := run("inp.test")
	if rtest != expected {
		log.Fatalln("FAILED TEST CASE", rtest, expected)
	}
	rmain := run("inp.main")
	fmt.Println("Result", rmain)
}
