package main

import (
	"cmp"
	"fmt"
	"slices"
	"time"
)

type Tuple2[A any, B any] struct {
	_0 A
	_1 B
}

func main() {
	lines := ReadLines("input.txt")
	Points := ParseLines(lines)

	fmt.Printf("Part 1: %d\n", Part1(Points))
	fmt.Printf("Part 2: %d\n", Part2(Points))
}

func Part1(points []Point2) int64 {
	ts := time.Now()
	defer func() {
		fmt.Printf("Part 1 took %s\n", time.Since(ts))
	}()

	convexHull := BuildConvexHull(points)
	furthestApart := slices.MaxFunc(AllPairs(convexHull), CompareByRectangleArea)
	return RectangleArea(furthestApart._0, furthestApart._1)
}

func Part2(points []Point2) int64 {
	ts := time.Now()
	defer func() {
		fmt.Printf("Part 2 took %s\n", time.Since(ts))
	}()

	// Input is given in clockwise order
	polygon := slices.Clone(points)
	slices.Reverse(polygon)

	candidates :=
		slices.DeleteFunc(
			AllPairs(points),
			func(t Tuple2[Point2, Point2]) bool {
				return !IsRectangleInside(t._0, t._1, polygon)
			})

	furthestApart := slices.MaxFunc(candidates, CompareByRectangleArea)

	return RectangleArea(furthestApart._0, furthestApart._1)
}

func CompareByRectangleArea(a, b Tuple2[Point2, Point2]) int {
	aDistance := RectangleArea(a._0, a._1)
	bDistance := RectangleArea(b._0, b._1)
	return cmp.Compare(aDistance, bDistance)
}

func FindFurthestApart(points []Point2) Tuple2[Point2, Point2] {
	return slices.MaxFunc(AllPairs(points), CompareByRectangleArea)
}

func AllPairs(points []Point2) []Tuple2[Point2, Point2] {
	var pairs []Tuple2[Point2, Point2]
	for i := range points {
		for j := i + 1; j < len(points); j++ {
			a := points[i]
			b := points[j]
			pairs = append(pairs, Tuple2[Point2, Point2]{a, b})
		}
	}
	return pairs
}

func IsRectangleInside(a Point2, b Point2, polygon []Point2) bool {
	minX := min(a.X, b.X)
	maxX := max(a.X, b.X)
	minY := min(a.Y, b.Y)
	maxY := max(a.Y, b.Y)
	rectanglePolygon := []Point2{{minX, minY}, {maxX, minY}, {maxX, maxY}, {minX, maxY}}
	for _, corner := range rectanglePolygon {
		if !RayCastIsInside(corner, polygon) {
			return false
		}
	}

	for i := range rectanglePolygon {
		recP0 := rectanglePolygon[i]
		recP1 := rectanglePolygon[(i+1)%4]

		for i := range polygon {
			polyP0 := polygon[i]
			polyP1 := polygon[(i+1)%len(polygon)]

			recEdge := recP1.Minus(recP0)
			polyEdge := polyP1.Minus(polyP0)

			cross1 := recEdge.CrossProduct(polyP0.Minus(recP0))
			cross2 := recEdge.CrossProduct(polyP1.Minus(recP0))
			cross3 := polyEdge.CrossProduct(recP0.Minus(polyP0))
			cross4 := polyEdge.CrossProduct(recP1.Minus(polyP0))
			if ((cross1.Z > 0 && cross2.Z < 0) || (cross1.Z < 0 && cross2.Z > 0)) &&
				((cross3.Z > 0 && cross4.Z < 0) || (cross3.Z < 0 && cross4.Z > 0)) {
				return false
			}
		}
	}
	return true
}

func RayCastIsInside(p Point2, polygon []Point2) bool {
	crossCount := 0
	for i := range polygon {
		e0 := polygon[i]
		e1 := polygon[(i+1)%len(polygon)]

		// Point contained in the edge
		if min(e0.X, e1.X) <= p.X && p.X <= max(e0.X, e1.X) &&
			min(e0.Y, e1.Y) <= p.Y && p.Y <= max(e0.Y, e1.Y) {
			return true
		}
		// Horizontal edge - skip
		if e0.Y == e1.Y {
			continue
		}
		// Half-open rule
		if (e0.Y > p.Y) == (e1.Y > p.Y) {
			continue
		}
		xIntersect := e0.X + (p.Y-e0.Y)*(e1.X-e0.X)/(e1.Y-e0.Y)
		if p.X < xIntersect {
			crossCount++
		}
	}
	return crossCount%2 == 1
}
