package main

import (
	"cmp"
	"fmt"
	"slices"
	"time"
)

func main() {
	lines := ReadLines("input.txt")
	ranges, ids := ParseLines(lines)
	fmt.Printf("Part 1: %d\n", FunPart1(ranges, ids))
	fmt.Printf("Part 2: %d\n", Part2(ranges))
}

func FunPart1(ranges []Range[uint64], ids []uint64) int {
	ts := time.Now()
	defer func() {
		fmt.Printf("Part 1 took %s\n", time.Since(ts))
	}()

	sum := 0
	tree := BuildBspTree(ranges, 1, RangeMedian)

	for _, id := range ids {
		matches := tree.Search(id, RangeContains)
		if matches.Len() > 0 {
			sum++
		}
	}
	return sum
}

func BoringPart1(ranges []Range[uint64], ids []uint64) int {
	ts := time.Now()
	defer func() {
		fmt.Printf("Part 1 took %s\n", time.Since(ts))
	}()

	sum := 0

	for _, id := range ids {
		for _, r := range ranges {
			if r.Contains(id) {
				sum++
				break
			}
		}
	}
	return sum
}

func Part2(ranges []Range[uint64]) uint64 {
	ts := time.Now()
	defer func() {
		fmt.Printf("Part 2 took %s\n", time.Since(ts))
	}()

	var sum uint64
	trimmedRanges := slices.Clone(ranges)
	slices.SortFunc(trimmedRanges, RangeCmpByEnd)

	for i := len(trimmedRanges) - 1; i > 0; i-- {
		if trimmedRanges[i].Intersects(trimmedRanges[i-1]) {
			trimmedRanges[i-1] = trimmedRanges[i-1].Union(trimmedRanges[i])
			trimmedRanges[i] = trimmedRanges[len(trimmedRanges)-1]
			trimmedRanges = trimmedRanges[:len(trimmedRanges)-1]
		}
	}

	for _, r := range trimmedRanges {
		sum += r.End - r.Begin
	}
	return sum
}

func RangeMedian(data []Range[uint64]) uint64 {
	if len(data) == 0 {
		panic("Median of empty slice")
	}
	var points []uint64
	for _, element := range data {
		points = append(points, (element.Begin+element.End)/2)
	}
	slices.Sort(points)
	if len(points)%2 == 1 {
		return points[len(points)/2]
	}
	return (points[len(points)/2-1] + points[len(points)/2]) / 2
}

func RangeContains(r Range[uint64], v uint64) bool { return r.Contains(v) }

func RangeCmpByEnd(a Range[uint64], b Range[uint64]) int {
	return cmp.Compare(a.End, b.End)
}
