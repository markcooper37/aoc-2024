package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	guardMap, err := readLines("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(partOne(guardMap))
	fmt.Println(partTwo(guardMap))
}

// partOne solves part one of the puzzle.
func partOne(guardMap [][]string) int {
	var currentPosition [2]int
	var currentDirection string
	for i := range guardMap {
		for j := range guardMap[0] {
			if guardMap[i][j] == "^" || guardMap[i][j] == ">" || guardMap[i][j] == "v" || guardMap[i][j] == "<" {
				currentPosition = [2]int{i, j}
				currentDirection = guardMap[i][j]
			}
		}
	}
	visitedPositions := map[[2]int]bool{}
	for {
		visitedPositions[currentPosition] = true
		nextPosition := findNextPosition(currentPosition, currentDirection)
		if nextPosition[0] < 0 || nextPosition[0] >= len(guardMap) || nextPosition[1] < 0 || nextPosition[1] >= len(guardMap[1]) {
			break
		} else if guardMap[nextPosition[0]][nextPosition[1]] == "#" {
			currentDirection = rotate(currentDirection)
		} else {
			currentPosition = nextPosition
		}
	}
	return len(visitedPositions)
}

// partTwo solves part two of the puzzle.
func partTwo(guardMap [][]string) int {
	var startPosition [2]int
	var startDirection string
	for i := range guardMap {
		for j := range guardMap[0] {
			if guardMap[i][j] == "^" || guardMap[i][j] == ">" || guardMap[i][j] == "v" || guardMap[i][j] == "<" {
				startPosition = [2]int{i, j}
				startDirection = guardMap[i][j]
			}
		}
	}
	total := 0
	for i := range guardMap {
		for j := range guardMap[0] {
			if [2]int{i, j} == startPosition || guardMap[i][j] == "#" {
				continue
			}
			mapCopy := copyMap(guardMap)
			mapCopy[i][j] = "#"
			currentPosition, currentDirection := startPosition, startDirection
			visitedPositions := map[[2]int]int{}
			for {
				visitedPositions[currentPosition]++
				if visitedPositions[currentPosition] >= 5 {
					total++
					break
				}
				nextPosition := findNextPosition(currentPosition, currentDirection)
				if nextPosition[0] < 0 || nextPosition[0] >= len(mapCopy) || nextPosition[1] < 0 || nextPosition[1] >= len(mapCopy[1]) {
					break
				} else if mapCopy[nextPosition[0]][nextPosition[1]] == "#" {
					currentDirection = rotate(currentDirection)
				} else {
					currentPosition = nextPosition
				}
			}
		}
	}
	return total
}

// findNextPosition finds the indices on step in the current direction.
func findNextPosition(currentPosition [2]int, direction string) [2]int {
	switch direction {
	case "^":
		return [2]int{currentPosition[0] - 1, currentPosition[1]}
	case ">":
		return [2]int{currentPosition[0], currentPosition[1] + 1}
	case "v":
		return [2]int{currentPosition[0] + 1, currentPosition[1]}
	case "<":
		return [2]int{currentPosition[0], currentPosition[1] - 1}
	}
	return [2]int{-1, -1}
}

// rotate rotates the direction 90 degrees to the right.
func rotate(direction string) string {
	switch direction {
	case "^":
		return ">"
	case ">":
		return "v"
	case "v":
		return "<"
	case "<":
		return "^"
	}
	return ""
}

// copyMap creates a copy of the guard map.
func copyMap(guardMap [][]string) [][]string {
	mapCopy := [][]string{}
	for i := range guardMap {
		rowCopy := make([]string, len(guardMap[i]))
		copy(rowCopy, guardMap[i])
		mapCopy = append(mapCopy, rowCopy)
	}
	return mapCopy
}

// readLines converts the information from the file into a usable form.
func readLines(fileName string) ([][]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)

	guardMap := [][]string{}
	for scanner.Scan() {
		guardMap = append(guardMap, strings.Split(scanner.Text(), ""))
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return guardMap, nil
}
