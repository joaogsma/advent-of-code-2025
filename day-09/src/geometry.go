package main

import (
	"cmp"
	"math"
	"slices"
)

type Point2 struct {
	X int64
	Y int64
}

type Point3 struct {
	X int64
	Y int64
	Z int64
}

func (p Point2) Plus(v Vector2) Point2 {
	return Point2{p.X + v.X, p.Y + v.Y}
}

func (a Point2) Minus(b Point2) Vector2 {
	return Vector2{a.X - b.X, a.Y - b.Y}
}

func (a Point2) DistanceTo(b Point2) float64 {
	return a.Minus(b).Magnitude()
}

type Vector2 struct {
	X int64
	Y int64
}

type Vector3 struct {
	X int64
	Y int64
	Z int64
}

func (v Vector2) DotProduct(b Vector2) int64 {
	return v.X*b.X + v.Y*b.Y
}

func (v Vector2) CrossProduct(b Vector2) Vector3 {
	return Vector3{v.X, v.Y, 0}.CrossProduct(Vector3{b.X, b.Y, 0})
}

func (a Vector3) CrossProduct(b Vector3) Vector3 {
	x := a.Y*b.Z - a.Z*b.Y
	y := a.Z*b.X - a.X*b.Z
	z := a.X*b.Y - a.Y*b.X
	return Vector3{x, y, z}
}

func (v Vector2) Magnitude() float64 {
	return math.Sqrt(float64(v.DotProduct(v)))
}

func (v Vector2) Mult(n int64) Vector2 {
	return Vector2{v.X * n, v.Y * n}
}

func (a Vector2) AngleWith(b Vector2) float64 {
	return math.Acos(float64(a.DotProduct(b)) / (a.Magnitude() * b.Magnitude()))
}

func RectangleArea(a, b Point2) int64 {
	deltaX := float64(a.X - b.X)
	deltaY := float64(a.Y - b.Y)
	return int64((math.Abs(deltaX) + 1) * (math.Abs(deltaY) + 1))
}

func IsCounterClockwise(a, b, c Point2) bool {
	v0 := b.Minus(a)
	v1 := c.Minus(a)
	return v0.CrossProduct(v1).Z > 0
}

func IsClockwise(a, b, c Point2) bool {
	v0 := b.Minus(a)
	v1 := c.Minus(a)
	return v0.CrossProduct(v1).Z < 0
}

func BuildConvexHull(points []Point2) []Point2 {
	var leftmostBottomPoint Point2 = points[0]
	for _, p := range points[1:] {
		if p.Y < leftmostBottomPoint.Y || (p.Y == leftmostBottomPoint.Y && p.X < leftmostBottomPoint.X) {
			leftmostBottomPoint = p
		}
	}

	orderedPoints := slices.Clone(points)
	slices.SortFunc(orderedPoints, PolarAngle(leftmostBottomPoint))

	convexHull := EmptyStack[Point2]()
	for _, p := range orderedPoints {
		for convexHull.Len() > 1 && !IsCounterClockwise(*convexHull.Peek(1), *convexHull.Peek(0), p) {
			convexHull.Pop()
		}
		convexHull.Push(p)
	}

	return convexHull.ToSlice()
}

func PolarAngle(origin Point2) func(Point2, Point2) int {
	return func(a, b Point2) int {
		var aAngle float64 = Vector2{1, 0}.AngleWith(a.Minus(origin))
		var bAngle float64 = Vector2{1, 0}.AngleWith(b.Minus(origin))
		if colinear := math.Abs(aAngle-bAngle) < 1e-8; !colinear {
			return cmp.Compare(aAngle, bAngle)
		}
		return cmp.Compare(origin.DistanceTo(a), origin.DistanceTo(b))
	}
}

func IsInsideConvex(p Point2, polygon []Point2) bool {
	for i := 0; i < len(polygon); i++ {
		edgeStart := polygon[i]
		edgeEnd := polygon[(i+1)%len(polygon)]
		if IsClockwise(edgeStart, edgeEnd, p) {
			return false
		}
	}
	return true
}
