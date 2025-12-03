package day02

import (
	"fmt"
	"strconv"
	"strings"
)

type IdRange struct {
	firstId int
	lastId  int
}

func Part1(line string) int {
	totalSumOfInvalidIds := 0
	ranges := parseInput(line)
	for _, r := range ranges {
		totalSumOfInvalidIds += sumInvalidIds(r)
	}
	return totalSumOfInvalidIds
}

func Part2(line string) int {
	return 0
}

func parseInput(line string) []IdRange {
	rangesAsString := strings.Split(line, ",")
	ranges := make([]IdRange, len(rangesAsString))
	for id, rangeString := range rangesAsString {
		ranges[id] = parseRange(rangeString)
	}
	return ranges
}

func parseRange(rangeString string) IdRange {
	ids := strings.Split(rangeString, "-")
	first, _ := strconv.Atoi(ids[0])
	last, _ := strconv.Atoi(ids[1])
	idRange := IdRange{firstId: first, lastId: last}
	return idRange
}

func sumInvalidIds(iRange IdRange) int {
	sum := 0

	startString := fmt.Sprint(iRange.firstId)
	l := len(startString)
	patternLen := l/2
	pattern := startString[:patternLen]
	valueStr := pattern + pattern
	value, _ := strconv.Atoi(valueStr)
	
	for value <= iRange.lastId {
		if value >= iRange.firstId && value <= iRange.lastId {
			sum  += value
		}
		patternValue, _ := strconv.Atoi(pattern)
		pattern = fmt.Sprint(patternValue + 1)
		valueStr = pattern + pattern
		value, _ = strconv.Atoi(valueStr)
	}
	return sum
}
