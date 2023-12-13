package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Formula struct {
	a  string
	op string
	b  string
}

type Monkey struct {
	value   int64
	formula Formula
	known   bool
}

var operations = []string{"+", "-", "*", "/"}

func findKnown(name string, monkeys map[string]*Monkey) {
	v := monkeys[name]
	// fmt.Println("at", name, *v)
	if name == "humn" {
		monkeys[name].known = false
		fmt.Println("human", monkeys[name])
		return
	}

	if v.known {
		// fmt.Println("returning", name, *v)
		return
	}

	findKnown(v.formula.a, monkeys)
	findKnown(v.formula.b, monkeys)

	a := monkeys[v.formula.a]
	b := monkeys[v.formula.b]

	if a.known && b.known {
		m1 := a.value
		m2 := b.value
		op := v.formula.op
		var res int64
		if op == "+" {
			res = m1 + m2
		} else if op == "-" {
			res = m1 - m2
		} else if op == "*" {
			res = m1 * m2
		} else if op == "/" {
			res = m1 / m2
		}

		monkeys[name].known = true
		monkeys[name].value = res
	}
	// fmt.Println("returniing", name, *monkeys[name])
}

func (m Monkey) String() string {
	if m.known {
		return fmt.Sprintf("%d", m.value)
	} else {
		return "x"
	}
}

func eval(name string, monkeys map[string]*Monkey, target int64) {
	m := monkeys[name]
	a := monkeys[m.formula.a]
	b := monkeys[m.formula.b]

	// fmt.Printf("At: %s %s %s = %d\n", a, m.formula.op, b, target)

	if name == "humn" {
		monkeys[name].known = true
		monkeys[name].value = target
		return
	}

	if !a.known && !b.known {
		panic("BOTH SIDES UNKNOWN")
	}

	if a.known {
		if name == "root" {
			eval(m.formula.b, monkeys, a.value)
		} else if m.formula.op == "+" {
			eval(m.formula.b, monkeys, target-a.value)
		} else if m.formula.op == "-" {
			eval(m.formula.b, monkeys, a.value-target)
		} else if m.formula.op == "*" {
			eval(m.formula.b, monkeys, target/a.value)
		} else if m.formula.op == "/" {
			eval(m.formula.b, monkeys, a.value/target)
		}
	} else {
		if name == "root" {
			eval(m.formula.a, monkeys, b.value)
		} else if m.formula.op == "+" {
			eval(m.formula.a, monkeys, target-b.value)
		} else if m.formula.op == "-" {
			eval(m.formula.a, monkeys, b.value+target)
		} else if m.formula.op == "*" {

			eval(m.formula.a, monkeys, target/b.value)
		} else if m.formula.op == "/" {

			eval(m.formula.a, monkeys, b.value*target)
		}
	}
	monkeys[name].known = true
	monkeys[name].value = target
}

func readInput() map[string]*Monkey {
	f, err := os.Open("inp")
	defer func() {
		err = f.Close()
		logErr(err)
	}()

	res := make(map[string]*Monkey)

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()

		line = strings.TrimSpace(line)
		spl := strings.Split(line, ": ")
		name := spl[0]
		valuestr := spl[1]
		num, err := strconv.Atoi(valuestr)
		var m Monkey
		if err != nil {
			for _, op := range operations {
				spl := strings.Split(valuestr, op)
				if len(spl) == 2 {
					f := Formula{strings.TrimSpace(spl[0]), op, strings.TrimSpace(spl[1])}
					m = Monkey{0, f, false}
					break
				}
			}
		} else {
			m = Monkey{int64(num), Formula{}, true}
		}
		res[name] = &m
	}

	return res
}

func printMonkeys(m map[string]*Monkey) {
	for k, v := range m {
		fmt.Println(k, *v)
	}
}

func main() {
	monkeys := readInput()
	printMonkeys(monkeys)
	findKnown("root", monkeys)
	eval("root", monkeys, 0)
	fmt.Println("Res:", monkeys["humn"].value)
}

func logErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
