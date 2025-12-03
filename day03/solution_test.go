package day03

import (
	"testing"
)

func TestPart1(t *testing.T) {
	t.Run("Day 3 example", func(t *testing.T) {
		input := []string{
			"987654321111111",
			"811111111111119",
			"234234234234278",
			"818181911112111",
		}
		expected := 357
		got := Part1(input)
		if got != expected {
			t.Errorf("Expected %d but got %d", expected, got)
		}
	})
}

func TestPart2(t *testing.T) {
	t.Run("Day 3 example", func(t *testing.T) {
		input := []string{
			"987654321111111",
			"811111111111119",
			"234234234234278",
			"818181911112111",
		}
		expected := 3121910778619
		got := Part2(input)
		if got != expected {
			t.Errorf("Expected %d but got %d", expected, got)
		}
	})
}
