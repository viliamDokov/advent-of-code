package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
	"time"
)

func check(err error) {
	if err != nil {
		log.Fatalf("Fatal error %s", err)
	}
}

type RangeMap struct {
	start  Range
	offset int64
}

type Range struct {
	start int64
	end   int64
}

func getOverlap(r1 Range, r2 Range) Range {
	return Range{max(r1.start, r2.start), min(r1.end, r2.end)}
}

func getDiff(r1 Range, r2 Range) []Range {
	diffs := make([]Range, 0)

	if r2.start <= r1.start {
		if r2.end <= r1.start {
			diffs = append(diffs, r1)
			return diffs
		} else if r2.end < r1.end {
			diffs = append(diffs, Range{r2.end, r1.end})
			return diffs
		} else {
			return diffs
		}
	} else if r2.start < r1.end {
		diffs = append(diffs, Range{r1.start, r2.start})
		if r2.end < r1.end {
			diffs = append(diffs, Range{r2.end, r1.end})
			return diffs
		} else {
			return diffs
		}
	} else if r2.start >= r1.end {
		diffs = append(diffs, r1)
		return diffs
	}

	log.Panicln("Never happen", r1, r2)
	return diffs
}

func mapRange(r Range, rm RangeMap) ([]Range, []Range) {
	mp := make([]Range, 0)
	ump := make([]Range, 0)
	overlap := getOverlap(r, rm.start)
	if overlap.start >= overlap.end {
		ump = append(ump, r)
		return ump, mp
	} else {
		mapped := Range{overlap.start + rm.offset, overlap.end + rm.offset}
		diff := getDiff(r, overlap)
		mp = append(mp, mapped)
		return diff, mp
	}

}

func remove(slice []Range, s int) []Range {
	return append(slice[:s], slice[s+1:]...)
}

func normalizeRanges(ranges []Range) []Range {
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].start < ranges[j].start
	})

	i := 0
	for i < len(ranges)-1 {

		if ranges[i+1].start <= ranges[i].end {
			e := max(ranges[i+1].end, ranges[i].end)
			ranges[i+1] = Range{ranges[i].start, e}
			ranges = remove(ranges, i)
		} else {
			i++
		}
	}
	return ranges
}

// func convertNumber(n int64, r RangeMap) (bool, int64) {
// 	if r.start <= n && n <= r.start+r.length-1 {
// 		d := n - r.start
// 		return true, r.dest + d
// 	}
// 	return false, n
// }

func solve(ranges []Range, allMaps [][]RangeMap) int64 {
	ranges = normalizeRanges(ranges)

	for _, rmaps := range allMaps {
		unmapped := slices.Clone(ranges)
		mapped := make([]Range, 0)
		for _, almp := range rmaps {
			new_unmapped := make([]Range, 0)
			for _, r := range unmapped {
				ump, mp := mapRange(r, almp)
				new_unmapped = append(new_unmapped, ump...)
				mapped = append(mapped, mp...)
			}
			unmapped = new_unmapped
		}

		mapped = append(mapped, unmapped...)
		ranges = normalizeRanges(mapped)
	}

	return ranges[0].start
}

func read() ([]Range, [][]RangeMap) {
	f, err := os.Open("input.txt")
	check(err)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	nums := make([]int64, 0)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	seeds := make([]Range, 0)
	allMaps := make([][]RangeMap, 0)

	readingMaps := false
	var rmaps []RangeMap
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "seeds: ") {
			line = line[7:]
			tmp := strings.Split(line, " ")
			for i := 0; i < len(tmp); i += 2 {
				n, err := strconv.ParseInt(strings.TrimSpace(tmp[i]), 10, 64)
				check(err)
				l, err2 := strconv.ParseInt(strings.TrimSpace(tmp[i+1]), 10, 64)
				check(err2)
				seeds = append(seeds, Range{n, n + l})
			}
		}

		if strings.TrimSpace(line) == "" {
			readingMaps = false
			allMaps = append(allMaps, rmaps)
		}

		if readingMaps {
			line = strings.TrimSpace(line)
			tmp := strings.Split(line, " ")
			dst, err := strconv.ParseInt(tmp[0], 10, 64)
			check(err)
			st, err := strconv.ParseInt(tmp[1], 10, 64)
			check(err)
			l, err := strconv.ParseInt(tmp[2], 10, 64)
			check(err)
			rmap := RangeMap{Range{st, st + l}, dst - st}
			rmaps = append(rmaps, rmap)
		}

		if strings.Contains(line, "map") {
			readingMaps = true
			rmaps = make([]RangeMap, 0)
		}

	}
	fmt.Println("NUMS:", len(nums))
	return seeds, allMaps[1:]
}

func tests() {
	ds := getDiff(Range{10, 20}, Range{1, 9})
	fmt.Println(ds)
	ds = getDiff(Range{10, 20}, Range{1, 10})
	fmt.Println(ds)

	ds = getDiff(Range{10, 20}, Range{1, 11})
	fmt.Println(ds)
	ds = getDiff(Range{10, 20}, Range{10, 11})
	fmt.Println(ds)
	ds = getDiff(Range{10, 20}, Range{11, 19})
	fmt.Println(ds)
	ds = getDiff(Range{10, 20}, Range{19, 20})
	fmt.Println(ds)
	ds = getDiff(Range{10, 20}, Range{20, 21})
	fmt.Println(ds)
	ds = getDiff(Range{10, 20}, Range{10, 20})
	fmt.Println(ds)
	ds = getDiff(Range{10, 20}, Range{1, 20})
	fmt.Println(ds)

}

func main() {
	start := time.Now()
	seeds, allMaps := read()
	fmt.Println(seeds, allMaps)
	s := solve(seeds, allMaps)
	fmt.Println(s)
	fmt.Println("Took", time.Since(start))
}
