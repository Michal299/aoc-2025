package day06

import (
	"testing"
)

func TestPart1(t *testing.T) {
	t.Run("Part1 example", func(t *testing.T) {
		input := []string{
			"123 328  51 64",
			"45 64  387 23",
			"6 98  215 314",
			"*   +   *   +",
		}
		expected := 4277556
		got := Part1(input)
		if got != expected {
			t.Errorf("Expected %d but got %d", expected, got)
		}
	})
}

func TestPart2(t *testing.T) {
	t.Run("Part2 example", func(t *testing.T) {
		input := []string{}
		expected := -1
		got := Part2(input)
		if got != expected {
			t.Errorf("Expected %d but got %d", expected, got)
		}
	})
}
