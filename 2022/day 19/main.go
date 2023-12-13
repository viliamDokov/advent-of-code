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

type Resources struct {
	ore      uint16
	clay     uint16
	obsidian uint16
	geode    uint16
}

type Costs struct {
	ore      Resources
	clay     Resources
	obsidian Resources
	geode    Resources
}

func addResource(r1 Resources, r2 Resources) Resources {
	return Resources{
		r1.ore + r2.ore,
		r1.clay + r2.clay,
		r1.obsidian + r2.obsidian,
		r1.geode + r2.geode,
	}
}

func removeResource(r1 Resources, r2 Resources) Resources {
	return Resources{
		r1.ore - r2.ore,
		r1.clay - r2.clay,
		r1.obsidian - r2.obsidian,
		r1.geode - r2.geode,
	}
}

func allGE(r1 Resources, r2 Resources) bool {
	return r1.ore >= r2.ore && r1.clay >= r2.clay && r1.obsidian >= r2.obsidian && r1.geode >= r2.geode
}

func getMaxProd(costs Costs) Resources {
	prod := Resources{costs.ore.ore, costs.obsidian.clay, costs.geode.obsidian, 0}
	if costs.clay.ore > prod.ore {
		prod.ore = costs.clay.ore
	}
	if costs.obsidian.ore > prod.ore {
		prod.ore = costs.obsidian.ore
	}
	if costs.geode.ore > prod.ore {
		prod.ore = costs.geode.ore
	}

	return prod
}

var visited = 0

func bestGeode(time uint16, resources Resources, production Resources, costs Costs, best uint16, maxProd Resources) uint16 {
	// fmt.Println(time, resources, costs)
	visited++
	if time == 0 {
		return resources.geode
	}

	ub := resources.geode + time*production.geode + (((time + 1) * time) / 2)
	if ub <= best {
		return resources.geode
	}

	if allGE(resources, costs.geode) {
		extra := Resources{0, 0, 0, 1}
		newResources := addResource(resources, production)
		newResources = removeResource(newResources, costs.geode)
		r := bestGeode(time-1, newResources, addResource(production, extra), costs, best, maxProd)
		if r > best {
			best = r
		}
	}

	if allGE(resources, costs.obsidian) && production.obsidian < maxProd.obsidian {
		extra := Resources{0, 0, 1, 0}
		newResources := addResource(resources, production)
		newResources = removeResource(newResources, costs.obsidian)
		r := bestGeode(time-1, newResources, addResource(production, extra), costs, best, maxProd)
		if r > best {
			best = r
		}
	}
	if allGE(resources, costs.ore) && production.ore < maxProd.ore {
		extra := Resources{1, 0, 0, 0}
		newResources := addResource(resources, production)
		newResources = removeResource(newResources, costs.ore)
		r := bestGeode(time-1, newResources, addResource(production, extra), costs, best, maxProd)
		if r > best {
			best = r
		}
	}
	if allGE(resources, costs.clay) && production.clay < maxProd.clay {
		extra := Resources{0, 1, 0, 0}
		newResources := addResource(resources, production)
		newResources = removeResource(newResources, costs.clay)
		r := bestGeode(time-1, newResources, addResource(production, extra), costs, best, maxProd)
		if r > best {
			best = r
		}
	}

	r := bestGeode(time-1, addResource(resources, production), production, costs, best, maxProd)
	if r > best {
		best = r
	}

	return best
}

func readInput() []Costs {
	f, err := os.Open("inp")
	defer func() {
		err = f.Close()
		logErr(err)
	}()

	res := make([]Costs, 0)

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		line = strings.Split(line, ": ")[1]
		data := strings.Split(line, ". ")

		// ORE
		oreOreCost, err := strconv.Atoi(strings.Split(data[0], " ")[4])
		logErr(err)
		oreCost := Resources{uint16(oreOreCost), 0, 0, 0}

		// clay
		clayOreCost, err := strconv.Atoi(strings.Split(data[1], " ")[4])
		logErr(err)
		clayCost := Resources{uint16(clayOreCost), 0, 0, 0}

		// obsidian
		obsidianOreCost, err := strconv.Atoi(strings.Split(data[2], " ")[4])
		logErr(err)
		obsidianClayCost, err := strconv.Atoi(strings.Split(data[2], " ")[7])
		logErr(err)
		obsidianCost := Resources{uint16(obsidianOreCost), uint16(obsidianClayCost), 0, 0}

		// geode
		geodeOreCost, err := strconv.Atoi(strings.Split(data[3], " ")[4])
		logErr(err)
		geodeObsidianCost, err := strconv.Atoi(strings.Split(data[3], " ")[7])
		logErr(err)
		geodeCost := Resources{uint16(geodeOreCost), 0, uint16(geodeObsidianCost), 0}

		res = append(res, Costs{oreCost, clayCost, obsidianCost, geodeCost})
	}

	return res
}

func main() {
	costs := readInput()
	acc := 1
	total := time.Now()
	for i, cost := range costs {
		start := time.Now()
		// cache, _ := lru.New[CacheEntry, int](1_000_000)
		maxProd := getMaxProd(cost)
		visited = 0
		res := bestGeode(24, Resources{0, 0, 0, 0}, Resources{1, 0, 0, 0}, cost, 0, maxProd)
		fmt.Println(i, res)
		fmt.Println("Took:", time.Since(start), "Visited", visited)
		acc *= int(res)
	}

	fmt.Println("ACC", acc, "Total time:", time.Since(total))
}

func logErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
