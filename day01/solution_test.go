package day01

import (
	"flag"
	"testing"
)

var inputFile = flag.String("input", "input.txt", "path to input file")

func TestPart1(t *testing.T) {
	t.Run("Day 1 example", func(t *testing.T) {
		input := []string{
			"L68",
			"L30",
			"R48",
			"L5",
			"R60",
			"L55",
			"L1",
			"L99",
			"R14",
			"L82",
		}
		expected := 3
		got := Part1(input)
		if got != expected {
			t.Errorf("Expected %d but got %d", expected, got)
		}
	})
}

func TestPart2(t *testing.T) {
	t.Run("Day 2 example", func(t *testing.T) {
		input := []string{
			"L68",
			"L30",
			"R48",
			"L5",
			"R60",
			"L55",
			"L1",
			"L99",
			"R14",
			"L82",
		}
		expected := 6
		got := Part2(input)
		if got != expected {
			t.Errorf("Expected %d but got %d", expected, got)
		}
	})

	t.Run("Overflow example", func(t *testing.T) {
		input := []string{
			"R1000",
		}
		expected := 10
		got := Part2(input)
		if got != expected {
			t.Errorf("Expected %d but got %d", expected, got)
		}
	})

	t.Run("Overflow example 2", func(t *testing.T) {
		input := []string{
			"R150",
		}
		expected := 2
		got := Part2(input)
		if got != expected {
			t.Errorf("Expected %d but got %d", expected, got)
		}
	})

		t.Run("Overflow example 3", func(t *testing.T) {
		input := []string{
			"L50",
		}
		expected := 1
		got := Part2(input)
		if got != expected {
			t.Errorf("Expected %d but got %d", expected, got)
		}
	})
}
