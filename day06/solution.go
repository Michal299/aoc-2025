package day06

import (
	"strconv"
	"strings"
)

type Sheet struct {
	columns []Column
}

type Operation byte

const (
	PLUS Operation = iota
	MULTIPLY
	UNKNOWN
)

type Column struct {
	values    []int
	operation Operation
}

func Part1(lines []string) int {
	sheet := parseSheet(lines)
	return sheet.evaluate()
}

func Part2(lines []string) int {
	return 0
}

func (s *Sheet) evaluate() (r int) {
	for _, column := range s.columns {
		r += column.evaluate()
	}
	return
}

func (c *Column) evaluate() int {
	plusFunc := func(aggr int, val int) int {
		return aggr + val
	}
	multFunc := func(aggr int, val int) int {
		return aggr * val
	}

	var aggrFunc func(int, int) int
	switch c.operation {
	case PLUS:
		aggrFunc = plusFunc
	case MULTIPLY:
		aggrFunc = multFunc
	}

	result := 0
	for i, v := range c.values {
		if i == 0 {
			result = v
		} else {
			result = aggrFunc(result, v)
		}
	}
	return result
}

func parseSheet(lines []string) *Sheet {
	columns := len(filterWhiteSpaces(strings.Split(lines[0], " ")))

	values := len(lines) - 1

	sheet := Sheet{columns: make([]Column, columns)}

	for c := range columns {
		sheet.columns[c] = Column{values: make([]int, values), operation: UNKNOWN}
	}

	for rowIndex, line := range lines[:values] {
		valuesInRow := filterWhiteSpaces(strings.Split(line, " "))
		for columnIndex, v := range valuesInRow {
			valueAsInt, _ := strconv.Atoi(v)
			sheet.columns[columnIndex].values[rowIndex] = valueAsInt
		}
	}
	for columnIndex, operation := range filterWhiteSpaces(strings.Split(lines[values], " ")) {
		opTrimmed := strings.Trim(operation, " ")
		sheet.columns[columnIndex].operation = getOperation(opTrimmed)
	}
	return &sheet
}

func filterWhiteSpaces(s []string) []string {
	var result []string
	for _, check := range s {
		if check == "" {
			continue
		}
		result = append(result, check)
	}
	return result
}

func getOperation(opString string) Operation {
	switch opString {
	case "+":
		return PLUS
	case "*":
		return MULTIPLY
	default:
		return UNKNOWN
	}
}
