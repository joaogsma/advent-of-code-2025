package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Coordinate struct {
	Row uint
	Col uint
}

func (c Coordinate) Up() Coordinate {
	return Coordinate{c.Row + 1, c.Col}
}

func (c Coordinate) Down() Coordinate {
	return Coordinate{c.Row - 1, c.Col}
}

func (c Coordinate) Left() Coordinate {
	return Coordinate{c.Row, c.Col - 1}
}

func (c Coordinate) Right() Coordinate {
	return Coordinate{c.Row, c.Col + 1}
}

func (c Coordinate) Neighbours() [8]Coordinate {
	up := c.Up()
	down := c.Down()
	left := c.Left()
	right := c.Right()
	return [8]Coordinate{up.Left(), up, up.Right(), right, right.Down(), down, down.Left(), left}
}

type Grid[T any] struct {
	data   []T
	Rows   uint
	Cols   uint
	toRune func(t T) rune
}

func NewGrid[T any](data [][]T, rows uint, cols uint, toRune func(t T) rune) Grid[T] {
	if uint(len(data)) != rows {
		panic("Wrong grid shape")
	}
	for _, row := range data {
		if uint(len(row)) != cols {
			panic("Wrong grid shape")
		}
	}

	singleSliceData := []T{}
	for i := int64(rows - 1); i >= 0; i-- {
		singleSliceData = append(singleSliceData, data[i]...)
	}

	return Grid[T]{singleSliceData, rows, cols, toRune}
}

func (g Grid[T]) AllCoordinates() []Coordinate {
	result := []Coordinate{}
	for row := uint(0); row < g.Rows; row++ {
		for col := uint(0); col < g.Cols; col++ {
			result = append(result, Coordinate{row, col})
		}
	}
	return result
}

func (g Grid[T]) Get(row uint, col uint) T {
	if !g.Contains(row, col) {
		panic(fmt.Sprintf("Coordinates (%d,%d) out of bounds for grid with shape %dx%d", row, col, g.Rows, g.Cols))
	}
	index := row*g.Cols + col
	return g.data[index]
}

func (g Grid[T]) GetCoord(c Coordinate) T {
	return g.Get(c.Row, c.Col)
}

func (g Grid[T]) Set(row uint, col uint, value T) {
	if !g.Contains(row, col) {
		panic(fmt.Sprintf("Coordinates (%d,%d) out of bounds for grid with shape %dx%d", row, col, g.Rows, g.Cols))
	}
	index := row*g.Cols + col
	g.data[index] = value
}

func (g Grid[T]) SetCoord(c Coordinate, value T) {
	g.Set(c.Row, c.Col, value)
}

func (g Grid[T]) Contains(row uint, col uint) bool {
	return row < g.Rows && col < g.Cols
}

func (g Grid[T]) ContainsCoord(c Coordinate) bool {
	return g.Contains(c.Row, c.Col)
}

func (g Grid[T]) String() string {
	var strb strings.Builder
	strb.WriteString(strconv.FormatUint(uint64(g.Rows), 10))
	strb.WriteRune('x')
	strb.WriteString(strconv.FormatUint(uint64(g.Cols), 10))
	strb.WriteRune('\n')
	for i := uint(0); i < g.Rows; i++ {
		row := g.Rows - 1 - i
		for col := uint(0); col < g.Cols; col++ {
			char := g.toRune(g.Get(row, col))
			strb.WriteRune(char)
		}
		if row > 0 {
			strb.WriteString("\n")
		}
	}
	return strb.String()
}
