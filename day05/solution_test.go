package day05

import (
	"testing"
)

func TestPart1(t *testing.T) {
	t.Run("Part1 example", func(t *testing.T) {
		input := []string{
			"3-5",
			"10-14",
			"16-20",
			"12-18",
			"",
			"1",
			"5",
			"8",
			"11",
			"17",
			"32",
		}
		expected := 3
		got := Part1(input)
		if got != expected {
			t.Errorf("Expected %d but got %d", expected, got)
		}
	})
}

func TestPart2(t *testing.T) {
	t.Run("Part2 example", func(t *testing.T) {
		input := []string{
			"3-5",
			"10-14",
			"16-20",
			"12-18",
		}
		expected := 14
		got := Part2(input)
		if got != expected {
			t.Errorf("Expected %d but got %d", expected, got)
		}
	})
}
