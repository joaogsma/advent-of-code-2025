package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Range struct {
	Begin uint64
	End   uint64
}

func (r Range) String() string {
	return strconv.FormatUint(r.Begin, 10) + "-" + strconv.FormatUint(r.End, 10)
}

func main() {
	lines := ReadLines("input.txt")
	if len(lines) > 1 {
		panic("More than one line detected")
	}
	ranges := ParseLine(lines[0])
	fmt.Printf("Part 1: %d\n", Part1(ranges))
	fmt.Printf("Part 2: %d\n", Part2(ranges))
}

func Part1(ranges []Range) uint64 {
	sum := uint64(0)
	for _, e := range ranges {
		for _, patternId := range FindPatternInputs(e, IsPatternTwice) {
			sum += patternId
		}
	}
	return sum
}

func Part2(ranges []Range) uint64 {
	sum := uint64(0)
	for _, e := range ranges {
		for _, patternId := range FindPatternInputs(e, IsPatternAtLeastTwiceLog) {
			sum += patternId
		}
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

func ParseLine(line string) []Range {
	rangeStrs := strings.Split(line, ",")
	result := []Range{}
	for _, r := range rangeStrs {
		parts := strings.Split(r, "-")
		begin, err := strconv.ParseUint(parts[0], 10, 64)
		if err != nil {
			panic(err)
		}
		end, err := strconv.ParseUint(parts[1], 10, 64)
		if err != nil {
			panic(err)
		}
		result = append(result, Range{Begin: begin, End: end})
	}
	return result
}

func FindPatternInputs(r Range, isPattern func(x uint64) bool) []uint64 {
	result := []uint64{}
	for id := r.Begin; id <= r.End; id++ {
		if !isPattern(id) {
			continue
		}
		result = append(result, id)
	}
	return result
}

func IsPatternTwice(value uint64) bool {
	str := strconv.FormatUint(value, 10)
	if len(str)%2 == 1 {
		return false
	}
	if str[:len(str)/2] == str[len(str)/2:] {
		return true
	}
	return false
}

//// Naive solution
// func IsPatternAtLeastTwice(value uint64) bool {
// 	valueStr := strconv.FormatUint(value, 10)
// 	for i := 1; i <= len(valueStr)/2; i++ {
// 		piece := valueStr[:i]
// 		if strings.ReplaceAll(valueStr, piece, "") == "" {
// 			return true
// 		}
// 	}
// 	return false
// }

func IsPatternAtLeastTwice(value uint64) bool {
	for base := uint64(10); base < value; base *= 10 {
		piece := value % base
		// When the first number is 0 it will still work, because any number mod 1 is 0
		previousPiece := value % (base / 10)
		if piece == previousPiece {
			continue
		}

		isPattern := true
		for current := value / base; current > 0 && isPattern; current /= base {
			isPattern = isPattern && (current%base) == piece
		}
		if isPattern {
			return true
		}
	}
	return false
}

func IsPatternAtLeastTwiceLog(value uint64) bool {
	totalDigits := NumDigits(value)
	for digits, base := uint(1), uint64(10); digits <= totalDigits/2; digits, base = digits+1, base*10 {
		piece := value % base
		// When the first number is 0 it will still work, because any number mod 1 is 0
		previousPiece := value % (base / 10)
		if piece == previousPiece {
			continue
		}

		isPattern := true
		for current := value / base; current > 0 && isPattern; current /= base {
			isPattern = isPattern && (current%base) == piece
		}
		if isPattern {
			return true
		}
	}
	return false
}

func NumDigits(value uint64) uint {
	return 1 + uint(math.Log10(float64(value)))
}
