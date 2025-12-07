package main

type Point struct {
	X int
	Y int
}

func (c Point) Up() Point {
	return Point{c.X, c.Y + 1}
}

func (c Point) Down() Point {
	return Point{c.X, c.Y - 1}
}

func (c Point) Left() Point {
	return Point{c.X - 1, c.Y}
}

func (c Point) Right() Point {
	return Point{c.X + 1, c.Y}
}
