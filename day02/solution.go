package day02

import (
	"strconv"
	"strings"
)

type IdRange struct {
	firstId int
	lastId  int
}

func Part1(line string) int {
	ranges := parseInput(line)
	return len(ranges)
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
