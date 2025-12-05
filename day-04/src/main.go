package main

import (
	"fmt"
)

type Cell int

const (
	Empty Cell = iota
	Occupied
)

func Rune(c Cell) rune {
	switch c {
	case Empty:
		return '.'
	case Occupied:
		return '@'
	default:
		panic("Unknown cell value")
	}
}

func main() {
	lines := ReadLines("input.txt")
	cells := ParseLines(lines)
	grid := NewGrid(cells, uint(len(cells)), uint(len(cells[0])), Rune)
	fmt.Printf("Part 1: %d\n", Part1(grid))
	fmt.Printf("Part 2: %d\n", Part2(grid))
}

func Part1(grid Grid[Cell]) int {
	sum := 0
	for row := uint(0); row < grid.Rows; row++ {
		for col := uint(0); col < grid.Cols; col++ {
			if grid.Get(row, col) == Empty {
				continue
			}
			adjacentRolls := 0
			coord := Coordinate{row, col}
			for _, neighbour := range coord.Neighbours() {
				if grid.ContainsCoord(neighbour) && grid.GetCoord(neighbour) == Occupied {
					adjacentRolls++
				}
			}
			if adjacentRolls < 4 {
				sum++
			}
		}
	}
	return sum
}

func Part2(grid Grid[Cell]) int {
	totalRemoved := 0

	reacheableRolls := EmptyQueue[Coordinate]()
	for _, pos := range grid.AllCoordinates() {
		if IsReacheableRoll(grid, pos) {
			reacheableRolls.Push(pos)
		}
	}

	for !reacheableRolls.IsEmpty() {
		pos, _ := reacheableRolls.Pop()
		if grid.GetCoord(pos) == Empty {
			continue
		}
		grid.SetCoord(pos, Empty)
		totalRemoved++
		for _, neighbour := range pos.Neighbours() {
			if grid.ContainsCoord(neighbour) && IsReacheableRoll(grid, neighbour) {
				reacheableRolls.Push(neighbour)
			}
		}
	}

	return totalRemoved
}

func IsReacheableRoll(grid Grid[Cell], pos Coordinate) bool {
	if grid.GetCoord(pos) == Empty {
		return false
	}
	adjacentRolls := 0
	for _, neighbour := range pos.Neighbours() {
		if grid.ContainsCoord(neighbour) && grid.GetCoord(neighbour) == Occupied {
			adjacentRolls++
		}
	}
	return adjacentRolls < 4
}
