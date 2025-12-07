package main

import (
	"fmt"
	"time"
)

func main() {
	lines := ReadLines("input.txt")
	start, splitters := ParseLines(lines)

	fmt.Printf("Part 1: %d\n", Part1(start, splitters))
	fmt.Printf("Part 2: %d\n", Part2(start, splitters))
}

func Part1(start Point, splitters [][]Point) int {
	ts := time.Now()
	defer func() {
		fmt.Printf("Part 1 took %s\n", time.Since(ts))
	}()

	var totalSplits int = 0
	rays := FilledQueue([]Point{start})
	markedRays := NewSet[Point]()
	markedSplitters := NewSet[Point]()

	for ray, ok := rays.Pop(); ok; ray, ok = rays.Pop() {
		if markedRays.Contains(ray) {
			continue
		}
		markedRays.Add(ray)
		splitter, ok := findSplitter(ray, splitters)
		if !ok || markedSplitters.Contains(splitter) {
			continue
		}
		markedSplitters.Add(splitter)
		totalSplits++
		rays.Push(splitter.Left())
		rays.Push(splitter.Right())
	}

	return totalSplits
}

func Part2(start Point, splitters [][]Point) int {
	ts := time.Now()
	defer func() {
		fmt.Printf("Part 2 took %s\n", time.Since(ts))
	}()

	cache := make(map[Point]int)
	CountTimelines(start, splitters, cache)
	return cache[start]
}

func CountTimelines(start Point, splitters [][]Point, cache map[Point]int) {
	if _, ok := cache[start]; ok {
		return
	}
	splitter, ok := findSplitter(start, splitters)
	if !ok {
		cache[start] = 1
		return
	}
	leftRay := splitter.Left()
	rightRay := splitter.Right()
	CountTimelines(leftRay, splitters, cache)
	CountTimelines(rightRay, splitters, cache)
	cache[start] = cache[leftRay] + cache[rightRay]
}

func findSplitter(start Point, splitters [][]Point) (Point, bool) {
	for _, splitter := range splitters[start.X] {
		if splitter.Y < start.Y {
			return splitter, true
		}
	}
	return Point{}, false
}
