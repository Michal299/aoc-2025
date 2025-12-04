package dayXY

import (
	"testing"
)

func TestPart1(t *testing.T) {
	t.Run("Day XY example", func(t *testing.T) {
		input := []string{}
		expected := -1
		got := Part1(input)
		if got != expected {
			t.Errorf("Expected %d but got %d", expected, got)
		}
	})
}

func TestPart2(t *testing.T) {
	t.Run("Day XY example", func(t *testing.T) {
		input := []string{}
		expected := -1
		got := Part2(input)
		if got != expected {
			t.Errorf("Expected %d but got %d", expected, got)
		}
	})
}
