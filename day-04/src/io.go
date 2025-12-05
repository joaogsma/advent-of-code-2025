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

func ParseLines(lines []string) [][]Cell {
	result := [][]Cell{}
	for _, line := range lines {
		result = append(result, ParseLine(line))
	}
	return result
}

func ParseLine(line string) []Cell {
	result := []Cell{}
	for _, c := range line {
		var cell Cell
		if c == '.' {
			cell = Empty
		} else {
			cell = Occupied
		}
		result = append(result, cell)
	}
	return result
}
