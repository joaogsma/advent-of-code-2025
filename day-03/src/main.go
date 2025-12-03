package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	lines := ReadLines("input.txt")
	banks := ParseLines(lines)
	fmt.Printf("Part 1: %d\n", Part1(banks))
	fmt.Printf("Part 2: %d\n", Part2(banks))
}

func Part1(banks [][]int) int64 {
	return SumJoltage(banks, 2)
}

func Part2(banks [][]int) int64 {
	return SumJoltage(banks, 12)
}

func SumJoltage(banks [][]int, numBatteries int) int64 {
	sum := int64(0)
	for _, bank := range banks {
		start := 0
		joltage := int64(0)
		for digits := 1; digits <= numBatteries; digits++ {
			remainingDigits := numBatteries - digits
			max, index := Max(bank[start : len(bank)-remainingDigits])
			joltage = 10*joltage + int64(max)
			start += index + 1
		}
		sum += joltage
	}
	return sum
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

func ParseLines(lines []string) [][]int {
	result := [][]int{}
	for _, line := range lines {
		result = append(result, ParseLine(line))
	}
	return result
}

func ParseLine(line string) []int {
	result := []int{}
	for _, c := range line {
		result = append(result, int(c-'0'))
	}
	return result
}

func Max(bank []int) (int, int) {
	index := 0
	max := bank[0]
	for i, val := range bank[1:] {
		if val > max {
			max = val
			index = i + 1
		}
	}
	return max, index
}
