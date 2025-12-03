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
		for _, patternId := range FindPatternInputs(e, IsPatternAtLeastTwice) {
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

func IsPatternAtLeastTwice(value uint64) bool {
	totalDigits := NumDigits(value)
	previousPiece := uint64(0)
	for digits, mask := uint(1), uint64(10); digits <= totalDigits/2; digits, mask = digits+1, mask*10 {
		piece := value % mask
		newDigitIsZero := piece == previousPiece
		if newDigitIsZero {
			continue
		}
		isPattern := true
		// Dividing by the mask shifts the number right by that many digits
		for current := value / mask; current > 0 && isPattern; current /= mask {
			isPattern = isPattern && (current%mask) == piece
		}
		if isPattern {
			return true
		}
		previousPiece = piece
	}
	return false
}

func IsPatternAtLeastTwiceAlt(value uint64) bool {
	totalDigits := NumDigits(value)
	previousPiece := uint64(0)
	for digits, mask := uint(1), uint64(10); digits <= totalDigits/2; digits, mask = digits+1, mask*10 {
		piece := value % mask
		newDigitIsZero := piece == previousPiece
		if newDigitIsZero {
			continue
		}
		expected := piece
		for ; expected < value; expected = mask*expected + piece {
		}
		if value == expected {
			return true
		}
		previousPiece = piece
	}
	return false
}

func NumDigits(value uint64) uint {
	return 1 + uint(math.Log10(float64(value)))
}
