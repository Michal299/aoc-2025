package main

import (
	"aoc-2025/day01"
	"aoc-2025/utils"
)

func main() {

	dayOneInput := utils.ReadInput("./input/day01/part1.txt")
	utils.PrintDay(1, day01.Part1(dayOneInput), day01.Part2(dayOneInput))

}
