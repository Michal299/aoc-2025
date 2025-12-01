package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func ReadInput(filePath string) []string {

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return lines
}

func printDayHeader(day int) {
	fmt.Printf("======== Day %d ========\n", day)
}

func printPart(part int, result int) {
	fmt.Printf("Part %d -> %d\n", part, result)
}

func PrintDay(day int, firstResult int, secondResult int) {
	printDayHeader(day)
	printPart(1, firstResult)
	printPart(2, secondResult)
	fmt.Println()
}
