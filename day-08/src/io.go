package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
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

func ParseLines(lines []string) []Point {
	var points []Point

	for _, line := range lines {
		parts := strings.Split(line, ",")
		x, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(err)
		}
		y, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}
		z, err := strconv.Atoi(parts[2])
		if err != nil {
			panic(err)
		}
		points = append(points, Point{x, y, z})
	}

	return points
}
