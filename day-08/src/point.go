package main

import "math"

type Point struct {
	X int
	Y int
	Z int
}

func (a Point) Plus(v Vector) Point {
	return Point{a.X + v.X, a.Y + v.Y, a.Z + v.Z}
}

func (a Point) Minus(b Point) Vector {
	return Vector{a.X - b.X, a.Y - b.Y, a.Z - b.Z}
}

func (a Point) DistanceTo(b Point) float64 {
	return a.Minus(b).Magnitude()
}

type Vector struct {
	X int
	Y int
	Z int
}

func (v Vector) DotProduct(other Vector) float64 {
	return math.Sqrt(float64(v.X*other.X + v.Y*other.Y + v.Z*other.Z))
}

func (v Vector) Magnitude() float64 {
	return v.DotProduct(v)
}

func (v Vector) Mult(n int) Vector {
	return Vector{v.X * n, v.Y * n, v.Z * n}
}
