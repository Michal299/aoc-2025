package day01

import (
	"strconv"
)

func Part1(lines []string) int {
	position := 50
	count := 0
	for _, line := range lines {
		direction, steps := parseInputLine(line)
		position, _ = rotate(position, direction, steps)
		if position == 0 {
			count++
		}
	}
	return count
}

func Part2(lines []string) int {
	position := 50
	count := 0

	for _, line := range lines {
		direction, steps := parseInputLine(line)
		resets := 0
		position, resets = rotate(position, direction, steps)
		count += resets
	}
	return count
}

func parseInputLine(line string) (byte, int) {
	direction := line[0]
	steps, err := strconv.Atoi(line[1:])
	if err != nil {
		panic("Unexpected input. The number of steps cannot be converted to int")
	}
	return direction, steps
}

func rotate(position int, direction byte, steps int) (int, int) {
	sign := 1
	switch direction {
	case 'R':
		sign = 1
	case 'L':
		sign = -1
	}

	actualSteps := steps % 100

	newPosition := position + (sign * actualSteps)

	resets := steps / 100

	if newPosition == 0 {
		resets++
	} else if newPosition < 0 {
		newPosition = 100 + newPosition
		if position != 0 {
			resets++
		}
	} else if newPosition >= 100 {
		newPosition = newPosition - 100
		resets++
	}
	return newPosition, resets
}
