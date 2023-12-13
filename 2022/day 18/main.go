package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

type Point struct {
	x int
	y int
	z int
}

func logErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func readInput() map[Point]bool {
	f, err := os.Open("inp")
	defer func() {
		err = f.Close()
		logErr(err)
	}()

	logErr(err)

	scanner := bufio.NewScanner(f)

	volcano := make(map[Point]bool)

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		data := strings.Split(line, ",")

		x, err := strconv.Atoi(data[0])
		logErr(err)
		y, err := strconv.Atoi(data[1])
		logErr(err)
		z, err := strconv.Atoi(data[2])
		logErr(err)
		p := Point{x, y, z}
		volcano[p] = true
	}
	return volcano
}

var changes = []Point{
	{1, 0, 0},
	{0, 1, 0},
	{0, 0, 1},
	{-1, 0, 0},
	{0, -1, 0},
	{0, 0, -1},
}

func inBounds(p Point, maxX int, maxY int, maxZ int, minX int, minY int, minZ int) bool {
	return p.x >= minX && p.x <= maxX && p.y >= minY && p.y <= maxY && p.z >= minZ && p.z <= maxZ
}

func getTrapped(volcano map[Point]bool, maxX int, maxY int, maxZ int, minX int, minY int, minZ int) map[Point]bool {
	air := make(map[Point]bool)
	outside := make(map[Point]bool)
	trapped := make(map[Point]bool)

	for i := minX; i <= maxX; i++ {
		for j := minY; j <= maxY; j++ {
			for k := minZ; k <= maxZ; k++ {
				p := Point{i, j, k}
				if !volcano[p] {
					air[p] = true
					if i == minX || i == maxX || j == minY || j == maxY || k == maxZ || k == minZ {
						outside[p] = true
					}
				}

			}

		}
	}
	fmt.Println(len(outside))
	i := 0
	for p, _ := range air {
		i++
		reached := make(map[Point]bool)
		// fmt.Println(i)
		queue := make([]Point, 0)
		queue = append(queue, p)
		for len(queue) > 0 {
			curr := queue[0]
			queue = queue[1:]
			// fmt.Println(i, len(queue))
			// fmt.Println(curr, reached, queue)
			for _, change := range changes {
				other := Point{curr.x + change.x, curr.y + change.y, curr.z + change.z}
				if !reached[other] && inBounds(other, maxX, maxY, maxZ, minX, minY, minZ) && !volcano[other] {
					// fmt.Println("setting", curr)
					if outside[other] {
						outside[p] = true
						break
					}
					if trapped[other] {
						trapped[p] = true
						break
					}
					reached[other] = true
					queue = append(queue, other)
				}
			}
		}

		if !outside[p] {
			trapped[p] = true
		}

	}
	return trapped
}

func isTrapped(trapped map[Point]bool, point Point, volcano map[Point]bool) bool {
	if volcano[point] {
		return false
	}
	if trapped[point] {
		return true
	}

	for _, change := range changes {
		other := Point{point.x + change.x, point.y + change.y, point.z + change.z}
		if !volcano[other] {
			return false
		}
	}

	trapped[point] = true
	return true
}

func getSurface(volcano map[Point]bool) int {
	surface := 0

	for point, _ := range volcano {
		for _, change := range changes {
			other := Point{point.x + change.x, point.y + change.y, point.z + change.z}
			if !volcano[other] {
				surface++
			}
		}
	}
	return surface
}

func Max[T comparable](arr map[T]bool, key func(T) int) int {
	best := 0
	for v, _ := range arr {
		k := key(v)
		if best < k {
			best = k
		}
	}
	return best
}

func Min[T comparable](arr map[T]bool, key func(T) int) int {
	best := math.MaxInt32
	for v, _ := range arr {
		k := key(v)
		if best > k {
			best = k
		}
	}
	return best
}

func main() {
	volcano := readInput()

	maxX := Max[Point](volcano, func(p Point) int { return p.x })
	maxY := Max[Point](volcano, func(p Point) int { return p.y })
	maxZ := Max[Point](volcano, func(p Point) int { return p.z })
	fmt.Printf("MAXES: %d, %d, %d\n", maxX, maxY, maxZ)

	minX := Min[Point](volcano, func(p Point) int { return p.x })
	minY := Min[Point](volcano, func(p Point) int { return p.y })
	minZ := Min[Point](volcano, func(p Point) int { return p.z })
	fmt.Printf("Mins: %d, %d, %d\n", minX, minY, minZ)

	start := time.Now()
	trapped := getTrapped(volcano, maxX, maxY, maxZ, minX, minY, minZ)
	r := getSurface(volcano)
	ts := getSurface(trapped)
	fmt.Println("result", r-ts, time.Since(start))
}
