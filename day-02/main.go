package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	reports, err := readLines("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(partOne(reports))
	fmt.Println(partTwo(reports))
}

// partOne solves part one of the puzzle.
func partOne(reports [][]int) int {
	safeCount := 0
	for _, report := range reports {
		if isSafe(report) {
			safeCount++
		}
	}

	return safeCount
}

// partTwo solves part two of the puzzle.
func partTwo(reports [][]int) int {
	safeCount := 0
	for _, report := range reports {
		if isSafe(report) {
			safeCount++
			continue
		}
		for i := range report {
			firstPart, secondPart := make([]int, i), make([]int, len(report) - i - 1)
			copy(firstPart, report[:i])
			copy(secondPart, report[i+1:])
			reducedReport := append(firstPart, secondPart...)
			if isSafe(reducedReport) {
				safeCount++
				break
			}
		}
	}

	return safeCount
}

// isSafe checks if a report is safe.
func isSafe(report []int) bool {
	inc := false
	for i := 0; i < len(report)-1; i++ {
		if report[i] == report[i+1] {
			return false
		} else if difference(report[i], report[i+1]) > 3 {
			return false
		}
		if i == 0 {
			inc = report[0] < report[1]
		} else {
			if (report[i] < report[i+1]) != inc {
				return false
			}
		}
	}
	return true
}

// difference returns the difference between two integers.
func difference(x, y int) int {
	diff := x - y
	if diff < 0 {
		return -diff
	}
	return diff
}

// readLines converts the information from the file into a usable form.
func readLines(fileName string) ([][]int, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)

	reports := [][]int{}
	for scanner.Scan() {
		strLevels := strings.Split(scanner.Text(), " ")
		levels := []int{}
		for _, strLevel := range strLevels {
			level, err := strconv.Atoi(strLevel)
			if err != nil {
				return nil, err
			}

			levels = append(levels, level)
		}

		reports = append(reports, levels)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return reports, nil
}
