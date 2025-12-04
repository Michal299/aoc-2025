package day04

import (
	"errors"
)

type Map [][]byte

type StoreMap struct {
	surface [][]byte
	width   int
	height  int
}

type Position struct {
	x int
	y int
}

const (
	EMPTY = '.'
	PAPER = '@'
)

func Part1(lines []string) int {
	m := parseMap(lines)
	return len(getAllAccessible(m))
}

func Part2(lines []string) int {
	m := parseMap(lines)
	totalRemoved := 0
	for {
		available := getAllAccessible(m)
		if len(available) == 0 {
			break
		}
		removeFromMap(m, available)
		totalRemoved += len(available)
	}
	return totalRemoved
}

func parseMap(lines []string) *StoreMap {
	result := make([][]byte, len(lines))
	for i, line := range lines {
		result[i] = []byte(line)
	}

	return &StoreMap{surface: result, width: len(result[0]), height: len(result)}
}

func (m *StoreMap) getValue(x, y int) (byte, error) {
	if x < 0 || x >= m.width || y < 0 || y >= m.height {
		return 0, errors.New("requested value out of the map")
	}
	return m.surface[y][x], nil
}

func (m *StoreMap) remove(x, y int) (bool, error) {
	if x < 0 || x >= m.width || y < 0 || y >= m.height {
		return false, errors.New("requested value out of the map")
	}
	m.surface[y][x] = EMPTY
	return true, nil
}

func getAllAccessible(m *StoreMap) []Position {
	var positions []Position
	for y := range m.height {
		for x := range m.width {
			current, _ := m.getValue(x, y)
			if current == PAPER && isAccessible(x, y, m) {
				positions = append(positions, Position{x: x, y: y})
			}
		}
	}
	return positions
}

func removeFromMap(m *StoreMap, itemsPosition []Position) {
	for _, item := range itemsPosition {
		m.remove(item.x, item.y)
	}
}

func isAccessible(x, y int, m *StoreMap) bool {
	return countNonEmptyNeighbours(x, y, m) < 4
}

func countNonEmptyNeighbours(x, y int, m *StoreMap) int {
	count := 0
	for _, xTarget := range []int{-1, 0, 1} {
		for _, yTarget := range []int{-1, 0, 1} {
			if xTarget == 0 && yTarget == 0 {
				continue
			}
			value, err := m.getValue(x+xTarget, y+yTarget)
			if err == nil && value == PAPER {
				count++
			}
		}
	}
	return count
}
