package main

import (
	"bufio"
	"os"
)

func ReadLines(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	result := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return result
}

func ParseLines(lines []string) (start Point, splitters [][]Point) {
	splitters = make([][]Point, len(lines[0]))

	for i, line := range lines {
		var y int = len(lines) - 1 - i
		for x, char := range line {
			switch char {
			case '^':
				splitters[x] = append(splitters[x], Point{x, y})
			case 'S':
				start = Point{x, y}
			}
		}
	}

	return start, splitters
}
