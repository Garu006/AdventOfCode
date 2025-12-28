package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type interval struct {
	start int64
	end   int64
}

func parseInput(path string) ([]interval, []int64, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Buffer(make([]byte, 0, 1024), 1024*1024)

	var ranges []interval
	var ids []int64
	readingRanges := true

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			readingRanges = false
			continue
		}
		if readingRanges {
			parts := strings.Split(line, "-")
			if len(parts) != 2 {
				return nil, nil, fmt.Errorf("invalid range line: %q", line)
			}
			a, err1 := strconv.ParseInt(parts[0], 10, 64)
			b, err2 := strconv.ParseInt(parts[1], 10, 64)
			if err1 != nil || err2 != nil {
				return nil, nil, fmt.Errorf("invalid range numbers: %q", line)
			}
			if a > b {
				a, b = b, a
			}
			ranges = append(ranges, interval{start: a, end: b})
		} else {
			v, err := strconv.ParseInt(line, 10, 64)
			if err != nil {
				return nil, nil, fmt.Errorf("invalid ingredient id: %q", line)
			}
			ids = append(ids, v)
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}
	return ranges, ids, nil
}

func mergeIntervals(ranges []interval) []interval {
	if len(ranges) == 0 {
		return ranges
	}
	sort.Slice(ranges, func(i, j int) bool { return ranges[i].start < ranges[j].start })

	merged := make([]interval, 0, len(ranges))
	cur := ranges[0]
	for i := 1; i < len(ranges); i++ {
		r := ranges[i]
		if r.start <= cur.end+1 {
			if r.end > cur.end {
				cur.end = r.end
			}
		} else {
			merged = append(merged, cur)
			cur = r
		}
	}
	merged = append(merged, cur)
	return merged
}

func countFresh(ids []int64, merged []interval) int {
	if len(merged) == 0 {
		return 0
	}
	starts := make([]int64, len(merged))
	for i := range merged {
		starts[i] = merged[i].start
	}

	total := 0
	for _, id := range ids {
		idx := sort.Search(len(starts), func(i int) bool { return starts[i] > id })
		if idx > 0 {
			prev := merged[idx-1]
			if id >= prev.start && id <= prev.end {
				total++
			}
		}
	}
	return total
}

// countCovered loops ranges and sums r.end - r.start + 1 into an int64 total.

func countCovered(merged []interval) int64 {
	var total int64
	for _, r:= range merged {
		total += r.end - r.start + 1
	}
	return total
}

func main() {
	path := "Day5/input.txt"
	if len(os.Args) > 1 {
		path = os.Args[1]
	}

	ranges, ids, err := parseInput(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}

	merged := mergeIntervals(ranges)
	covered := countCovered(merged)
	fresh := countFresh(ids, merged)
	fmt.Println(fresh)
	fmt.Println(covered)
}