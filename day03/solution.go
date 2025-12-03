package day03

import (
	"fmt"
	"strconv"
)

func Part1(lines []string) int {
	totalJoltage := 0
	for _, line := range lines {
		totalJoltage += getMaximumJoltage(line)
	}
	return totalJoltage
}

func Part2(lines []string) int {
	totalJoltage := 0
	for i, line := range lines {
		fmt.Printf("Processing line %d\n", i)
		batteries := []byte(line)
		batteriesInt := make([]int, len(batteries))
		for i, b := range batteries {
			batteriesInt[i] = int(b - '0')
		}
		joltage, _ := strconv.Atoi(getMaximumJoltageForBatteries(batteriesInt, 12))
		totalJoltage += joltage
	}
	return totalJoltage
}

func getMaximumJoltage(bank string) int {
	maxJoltage := 0
	for firstDigitId := 0; firstDigitId < len(bank); firstDigitId++ {
		for secondDigitId := firstDigitId + 1; secondDigitId < len(bank); secondDigitId++ {
			joltage, _ := strconv.Atoi(string(bank[firstDigitId]) + string(bank[secondDigitId]))
			if maxJoltage < joltage {
				maxJoltage = joltage
			}
		}
	}
	return maxJoltage
}

func getMaximumJoltageForBatteries(bank []int, batteries int) string {
	if batteries == 1 {
		v, _ := findMax(bank)
		return fmt.Sprint(v)
	}
	bankLen := len(bank)
	bankToConsider := bank[:bankLen-batteries+1]

	maxVal, maxIndex := findMax(bankToConsider)

	remaining := getMaximumJoltageForBatteries(bank[maxIndex+1:], batteries-1)
	return fmt.Sprint(maxVal) + remaining
}

func findMax(bytes []int) (val int, index int) {
	val = 0
	index = 0
	for i, n := range bytes {
		if n > val {
			val = n
			index = i
		}
	}
	return
}
