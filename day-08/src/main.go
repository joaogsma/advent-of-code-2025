package main

import (
	"cmp"
	"fmt"
	"math"
	"slices"
	"time"
)

type Tuple2[A any, B any] struct {
	_0 A
	_1 B
}

func main() {
	lines := ReadLines("input.txt")
	points := ParseLines(lines)

	fmt.Printf("Part 1: %d\n", Part1(points, 1000))
	fmt.Printf("Part 2: %d\n", Part2(points))
}

func Part1(boxes []Point, iterations int) int {
	ts := time.Now()
	defer func() {
		fmt.Printf("Part 1 took %s\n", time.Since(ts))
	}()

	var circuits, _ = BuildCircuits(boxes, iterations)

	orderedCircuits := circuits.ToSlice()
	slices.SortFunc(
		orderedCircuits,
		func(a, b *Set[Point]) int {
			return cmp.Compare(a.Len(), b.Len())
		})

	var result int = 1
	for _, circuit := range orderedCircuits[max(0, len(orderedCircuits)-3):] {
		result *= circuit.Len()
	}

	return result
}

func Part2(boxes []Point) int {
	ts := time.Now()
	defer func() {
		fmt.Printf("Part 2 took %s\n", time.Since(ts))
	}()

	var _, lastPair = BuildCircuits(boxes, math.MaxInt)
	return lastPair._0.X * lastPair._1.X
}

func BuildCircuits(boxes []Point, iterations int) (circuits *Set[*Set[Point]], lastPair Tuple2[Point, Point]) {
	var boxPairs []Tuple2[Point, Point] = BuildOrderedDistances(boxes)

	circuits = NewSet[*Set[Point]]()
	circuitsByBox := make(map[Point]*Set[Point])
	for _, box := range boxes {
		circuitsByBox[box] = NewSet(box)
		circuits.Add(circuitsByBox[box])
	}

	for i, pair := range boxPairs {
		if i >= iterations {
			break
		}
		var aCircuit, bCircuit *Set[Point] = circuitsByBox[pair._0], circuitsByBox[pair._1]
		if aCircuit == bCircuit {
			continue
		}
		lastPair = pair
		var newCircuit *Set[Point] = aCircuit.Union(bCircuit)
		for _, box := range newCircuit.Values() {
			circuitsByBox[box] = newCircuit
		}
		circuits.Remove(aCircuit)
		circuits.Remove(bCircuit)
		circuits.Add(newCircuit)
	}

	return circuits, lastPair
}

func CompareByDistance(a, b Tuple2[Point, Point]) int {
	aDistance := a._0.DistanceTo(a._1)
	bDistance := b._0.DistanceTo(b._1)
	return cmp.Compare(aDistance, bDistance)
}

func BuildOrderedDistances(boxes []Point) []Tuple2[Point, Point] {
	var pairs []Tuple2[Point, Point]
	for i := range boxes {
		for j := i + 1; j < len(boxes); j++ {
			a := boxes[i]
			b := boxes[j]
			pairs = append(pairs, Tuple2[Point, Point]{a, b})
		}
	}
	slices.SortFunc(pairs, CompareByDistance)
	return pairs
}
