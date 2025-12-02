package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Direction int

const (
	Left Direction = iota
	Right
)

type Step struct {
	Direction Direction
	Distance  int
}

func (d Direction) String() string {
	switch d {
	case Left:
		return "L"
	case Right:
		return "R"
	default:
		return "?"
	}
}

func (s Step) String() string {
	distanceStr := strconv.Itoa(s.Distance)
	if s.Direction == Left {
		return "-" + distanceStr
	}
	return "+" + distanceStr
}

func main() {
	lines := ReadLines("input.txt")
	instructions := ParseLines(lines)
	fmt.Printf("Part 1: %d\n", CountStopsAtZero(instructions, 50))
	fmt.Printf("Part 2: %d\n", CountPassesThroughZero(instructions, 50))
}

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

func ParseLines(lines []string) []Step {
	var result = make([]Step, len(lines))
	for i, line := range lines {
		dir, dist := ParseLine(line)
		result[i] = Step{Direction: dir, Distance: dist}
	}
	return result
}

func ParseLine(line string) (Direction, int) {
	var direction Direction
	switch line[0] {
	case 'L':
		direction = Left
	case 'R':
		direction = Right
	default:
		panic(fmt.Sprintf("Unknown direction %c", line[0]))
	}

	distance, err := strconv.Atoi(line[1:])
	if err != nil {
		panic(err)
	}

	return direction, distance
}

func CountStopsAtZero(instructions []Step, startingValue int) int {
	if startingValue < 0 || startingValue > 99 {
		panic(fmt.Sprintf("Illegal starting number %d", startingValue))
	}

	zeroCounter := 0
	current := startingValue
	for _, elem := range instructions {
		current, _ = RunStep(elem, current)
		if current == 0 {
			zeroCounter++
		}
	}

	return zeroCounter
}

func CountPassesThroughZero(instructions []Step, startingValue int) int {
	if startingValue < 0 || startingValue > 99 {
		panic(fmt.Sprintf("Illegal starting number %d", startingValue))
	}

	zeroCounter := 0
	current := startingValue
	for _, elem := range instructions {
		var zeroCount int
		current, zeroCount = RunStep(elem, current)
		zeroCounter += zeroCount
	}

	return zeroCounter
}

func RunStep(step Step, start int) (int, int) {
	result := start
	adjustedDistance := step.Distance % 100
	zeroCounter := 0
	fullTurns := AbsInt(step.Distance / 100)
	zeroCounter += fullTurns

	if adjustedDistance == 0 {
		return result, zeroCounter
	}

	switch step.Direction {
	case Left:
		result -= adjustedDistance
		if result == 0 {
			zeroCounter++
		} else if result < 0 {
			result += 100
			if start != 0 {
				zeroCounter++
			}
		}
	case Right:
		result += adjustedDistance
		if result >= 100 {
			result -= 100
			zeroCounter++
		}
	}

	return result, zeroCounter
}

func Sign(x int) int {
	if x > 0 {
		return 1
	} else if x < 0 {
		return -1
	}
	return 0
}

func AbsInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
