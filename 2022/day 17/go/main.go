package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"
	"time"
)

type Pattern struct {
	height    int
	iteration int
}

func overlaps(rock []uint8, cavern []uint8, x int, y int, h int) bool {
	if x < 0 {
		return true
	}

	for i, row := range rock {
		if y+i >= h+len(cavern) {
			if (row>>x)&1 != 0 {
				return true
			}
		} else {
			if (row>>x)&cavern[y+i-h] != 0 {
				return true
			}
		}

	}

	return false
}

func visualize(cavern []uint8, h int) {
	f, err := os.Create("out")
	defer f.Close()

	if err != nil {
		log.Fatal(err)
	}

	w := bufio.NewWriter(f)

	for i := len(cavern) - 1; i >= 0; i-- {
		row := cavern[i]
		s := fmt.Sprintf("%07b", row>>1)
		s = strings.Replace(s, "0", ".", -1)
		s = strings.Replace(s, "1", "#", -1)
		fmt.Print(fmt.Sprintf("|%s|\n", s))
		w.WriteString(fmt.Sprintf("|%s|\n", s))
	}
}

func simulate(moves []int) int {
	f, err := os.Create("out")
	defer f.Close()

	if err != nil {
		log.Fatal(err)
	}

	h := 0
	rocks := make([][]uint8, 5)
	rocks[0] = []uint8{240}
	rocks[1] = []uint8{64, 224, 64}
	rocks[2] = []uint8{224, 32, 32}
	rocks[3] = []uint8{128, 128, 128, 128}
	rocks[4] = []uint8{192, 192}

	seen := make([][]map[uint64]Pattern, len(moves))
	for i := range moves {
		seen[i] = make([]map[uint64]Pattern, len(rocks))
		for j := range rocks {
			seen[i][j] = make(map[uint64]Pattern)
		}
	}
	jumped := false

	cavern := make([]uint8, 1)
	cavern[0] = 255

	rock_i := uint8(0)
	dir_i := 0

	target := 1_000_000_000_000
	for i := 0; i < target; i++ {

		x, y := 2, h+len(cavern)+3
		if !overlaps(rocks[rock_i], cavern, x+moves[dir_i], y, h) {
			x = x + moves[dir_i]
		}

		dir_i++
		if dir_i >= len(moves) {
			dir_i = 0
		}

		for !overlaps(rocks[rock_i], cavern, x, y-1, h) {
			y = y - 1

			if !overlaps(rocks[rock_i], cavern, x+moves[dir_i], y, h) {
				x = x + moves[dir_i]
			}

			dir_i++
			if dir_i >= len(moves) {
				dir_i = 0
			}

		}

		// if seen[dir_i]&rock_i != 0 {
		// 	// fmt.Printf("REPEAT dir: %d  rock:%d at %d\n", dir_i, rock_i, i)
		// }

		for y+len(rocks[rock_i]) > h+len(cavern) {
			cavern = append(cavern, 1)
		}

		for j, row := range rocks[rock_i] {
			cavern[y+j-h] |= (row >> x)

			if cavern[y+j-h] == 255 && !jumped {
				cavern = cavern[y+j-h:]
				if len(cavern) < 9 {
					num := uint64(0)
					for k, layer := range cavern {
						num |= (uint64(layer) << (8 * k))
					}
					// fmt.Printf("At: %d\n", i)
					data := seen[dir_i][rock_i][num]
					if data.height != 0 && data.iteration != 0 {
						fmt.Printf("Found pattern!")
						fmt.Println(data)
						dh := h - data.height
						dt := i - data.iteration
						mul := (target - i) / dt
						h += dh * mul
						i += dt * mul
						y += dh * mul
						fmt.Printf("Jumped to: mul=%d,h=%d,i=%d", mul, h, i)
						jumped = true
					}

					seen[dir_i][rock_i][num] = Pattern{h, i}
				}
				h += y + j - h
			}
		}

		// fmt.Printf("Full block top at %d\n", i)
		// if cavern[len(cavern)-1] == 255 {
		// 	if seen[dir_i]&rock_i != 0 {
		// 		fmt.Printf("REPEAT dir: %d  rock:%d at %d\n", dir_i, rock_i, i)
		// 	}
		// 	seen[dir_i] |= uint8(rock_i)
		// }

		rock_i++
		if int(rock_i) >= len(rocks) {
			rock_i = 0
		}
	}
	// visualize(cavern)
	return h + len(cavern) - 1
}

func get_input() []int {

	f, err := os.ReadFile("inp")
	if err != nil {
		log.Fatal(err)
	}

	str := strings.TrimSpace(string(f))
	dirs := make([]int, len(str))
	for i, c := range str {
		if c == '>' {
			dirs[i] = 1
		} else {
			dirs[i] = -1
		}
	}
	return dirs
}

func main() {
	moves := get_input()
	s := time.Now()
	r := simulate(moves)
	fmt.Printf("Result: %d in: %s\n", r, time.Since(s))
	PrintMemUsage()
}

// PrintMemUsage outputs the current, total and OS memory being used. As well as the number
// of garage collection cycles completed.
func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
