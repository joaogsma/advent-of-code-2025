package main

import (
	"fmt"
	"time"
)

func main() {
	lines := ReadLines("input.txt")
	numbers, operators := ParseLinesPart1(lines)

	fmt.Printf("Part 1: %d\n", Part1(numbers, operators))
	fmt.Printf("Part 2: %d\n", Part2(lines))
}

func Part1(numbers [][]int, operators []func(int, int) int) int {
	ts := time.Now()
	defer func() {
		fmt.Printf("Part 1 took %s\n", time.Since(ts))
	}()

	var cols int = len(numbers[0])
	var results []int
	for col := 0; col < cols; col++ {
		var operands []int
		for row := 0; row < len(numbers); row++ {
			operands = append(operands, numbers[row][col])
		}
		results = append(results, Reduce(operands, operators[col]))
	}
	return Reduce(results, Plus)
}

func Part2(lines []string) int {
	ts := time.Now()
	defer func() {
		fmt.Printf("Part 2 took %s\n", time.Since(ts))
	}()

	var results []int
	var numberLines []string = lines[:len(lines)-1]
	var operatorLine string = lines[len(lines)-1]

	var operands []int
	for col := len(lines[0]) - 1; col >= 0; col-- {
		var acc int = 0
		for row := range numberLines {
			if char := numberLines[row][col]; char != ' ' {
				acc = 10*acc + int(char-'0')
			}
		}
		operands = append(operands, acc)

		switch operatorLine[col] {
		case '*':
			results = append(results, Reduce(operands, Mult))
			operands = []int{}
			col--
		case '+':
			results = append(results, Reduce(operands, Plus))
			operands = []int{}
			col--
		}
	}

	return Reduce(results, Plus)
}

func Reduce[T any](values []T, fn func(T, T) T) T {
	var acc T = values[0]
	for _, v := range values[1:] {
		acc = fn(acc, v)
	}
	return acc
}

func Plus(a int, b int) int { return a + b }
func Mult(a int, b int) int { return a * b }
