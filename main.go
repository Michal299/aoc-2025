package main

import (
	"aoc-2025/day01"
	"aoc-2025/day02"
	"aoc-2025/day03"
	"aoc-2025/day04"
	"aoc-2025/utils"
)

func main() {

	dayOneInput := utils.ReadInput("./input/day01/part1.txt")
	utils.PrintDay(1, day01.Part1(dayOneInput), day01.Part2(dayOneInput))

	dayTwoInput := utils.ReadInput("./input/day02/part1.txt")[0]
	utils.PrintDay(2, day02.Part1(dayTwoInput), day02.Part2(dayTwoInput))

	dayThreeInput := utils.ReadInput("./input/day03/part1.txt")
	utils.PrintDay(3, day03.Part1(dayThreeInput), day03.Part2(dayThreeInput))

	dayFourInput := utils.ReadInput("./input/day04/part1.txt")
	utils.PrintDay(4, day04.Part1(dayFourInput), day04.Part2(dayFourInput))

}
