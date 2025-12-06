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

func ParseLinesPart1(lines []string) (numbers [][]int, operators []func(int, int) int) {
	for i, line := range lines {
		var isLastLine bool = i == len(lines)-1
		var parts []string = strings.Split(line, " ")
		var lineOperands []int
		for _, elem := range parts {
			if len(elem) == 0 {
				continue
			}
			if !isLastLine {
				number, err := strconv.Atoi(elem)
				if err != nil {
					panic("Not a number: " + elem)
				}
				lineOperands = append(lineOperands, number)
				continue
			}
			switch elem {
			case "*":
				operators = append(operators, Mult)
			case "+":
				operators = append(operators, Plus)
			default:
				panic("Unknown operator " + elem)
			}
		}
		if !isLastLine {
			numbers = append(numbers, lineOperands)
		}
	}

	return numbers, operators
}
