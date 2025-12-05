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

func ParseLines(lines []string) ([]Range[uint64], []uint64) {
	var ranges []Range[uint64]
	var ids []uint64

	parsingRanges := true
	for _, line := range lines {
		if line == "" {
			parsingRanges = false
			continue
		}
		if parsingRanges {
			ranges = append(ranges, ParseRange(line))
		} else {
			ids = append(ids, ParseIds(line))
		}
	}

	return ranges, ids
}

func ParseRange(line string) Range[uint64] {
	parts := strings.Split(line, "-")
	begin, err := strconv.ParseUint(parts[0], 10, 64)
	if err != nil {
		panic("Invalid range start " + parts[0])
	}
	end, err := strconv.ParseUint(parts[1], 10, 64)
	if err != nil {
		panic("Invalid range end " + parts[1])
	}
	return Range[uint64]{begin, end + 1}
}

func ParseIds(line string) uint64 {
	id, err := strconv.ParseUint(line, 10, 64)
	if err != nil {
		panic("Invalid id " + line)
	}
	return id
}
