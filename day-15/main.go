package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	warehouseMap, movements, err := readLines("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(partOne(warehouseMap, movements))
	fmt.Println(partTwo(warehouseMap, movements))
}

// partOne solves part one of the puzzle.
func partOne(warehouseMap [][]string, movements []string) int {
	for _, movement := range movements {
		warehouseMap = iterateWarehouse(warehouseMap, movement)
	}
	return sumCoordinates(warehouseMap)
}

// partTwo solves part two of the puzzle.
func partTwo(warehouseMap [][]string, movements []string) int {
	resizedMap := resizeMap(warehouseMap)
	for _, movement := range movements {
		resizedMap = iterateResizedWarehouse(resizedMap, movement)
	}
	return sumResizedCoordinates(resizedMap)
}

// iterateWarehouse performs the movement.
func iterateWarehouse(warehouseMap [][]string, movement string) [][]string {
	warehouseCopy := copyWarehouse(warehouseMap)
	var robotPosition [2]int
	for i, row := range warehouseCopy {
		for j, position := range row {
			if position == "@" {
				robotPosition = [2]int{i, j}
			}
		}
	}
	positionsToMove := getPositionsToMove(warehouseMap, robotPosition, movement)
	for i := len(positionsToMove) - 1; i >= 0; i-- {
		nextPosition := nextPosition(positionsToMove[i], movement)
		warehouseCopy[nextPosition[0]][nextPosition[1]] = warehouseCopy[positionsToMove[i][0]][positionsToMove[i][1]]
		warehouseCopy[positionsToMove[i][0]][positionsToMove[i][1]] = "."
	}
	return warehouseCopy
}

// copyWarehouse creates a copy of the warehouse.
func copyWarehouse(warehouseMap [][]string) [][]string {
	warehouseCopy := [][]string{}
	for _, row := range warehouseMap {
		rowCopy := make([]string, len(row))
		copy(rowCopy, row)
		warehouseCopy = append(warehouseCopy, rowCopy)
	}
	return warehouseCopy
}

// getPositionsToMove gets all positions in the given movement direction up to the first wall or empty space.
func getPositionsToMove(warehouseMap [][]string, robotPosition [2]int, movement string) [][2]int {
	positions := [][2]int{robotPosition}
	for {
		nextPosition := nextPosition(positions[len(positions)-1], movement)
		if warehouseMap[nextPosition[0]][nextPosition[1]] == "." {
			return positions
		} else if warehouseMap[nextPosition[0]][nextPosition[1]] == "#" {
			return [][2]int{}
		}
		positions = append(positions, nextPosition)
	}
}

// nextPosition gets the next position in the movement direction.
func nextPosition(currentPosition [2]int, movement string) [2]int {
	if movement == "^" {
		return [2]int{currentPosition[0] - 1, currentPosition[1]}
	} else if movement == ">" {
		return [2]int{currentPosition[0], currentPosition[1] + 1}
	} else if movement == "v" {
		return [2]int{currentPosition[0] + 1, currentPosition[1]}
	} else {
		return [2]int{currentPosition[0], currentPosition[1] - 1}
	}
}

// sumCoordinates sums the coordinates of the items in the warehouse.
func sumCoordinates(warehouseMap [][]string) int {
	total := 0
	for i, row := range warehouseMap {
		for j, position := range row {
			if position == "O" {
				total += i*100 + j
			}
		}
	}
	return total
}

// resizeMap converts a warehouse map to the larger size.
func resizeMap(warehouseMap [][]string) [][]string {
	newMap := [][]string{}
	for _, row := range warehouseMap {
		newRow := []string{}
		for _, position := range row {
			if position == "#" {
				newRow = append(newRow, "#", "#")
			} else if position == "." {
				newRow = append(newRow, ".", ".")
			} else if position == "O" {
				newRow = append(newRow, "[", "]")
			} else {
				newRow = append(newRow, "@", ".")
			}
		}
		newMap = append(newMap, newRow)
	}
	return newMap
}

// iterateResizedWarehouse performs the movement on the resized warehouse.
func iterateResizedWarehouse(warehouseMap [][]string, movement string) [][]string {
	warehouseCopy := copyWarehouse(warehouseMap)
	var robotPosition [2]int
	for i, row := range warehouseCopy {
		for j, position := range row {
			if position == "@" {
				robotPosition = [2]int{i, j}
			}
		}
	}
	positionsToMove := getResizedPositionsToMove(warehouseMap, robotPosition, movement)
	for i := len(positionsToMove) - 1; i >= 0; i-- {
		nextPosition := nextPosition(positionsToMove[i], movement)
		warehouseCopy[nextPosition[0]][nextPosition[1]] = warehouseCopy[positionsToMove[i][0]][positionsToMove[i][1]]
		warehouseCopy[positionsToMove[i][0]][positionsToMove[i][1]] = "."
	}

	return warehouseCopy
}

// getResizedPositionsToMove gets all positions in the given movement direction up to the first wall or empty space.
func getResizedPositionsToMove(warehouseMap [][]string, robotPosition [2]int, movement string) [][2]int {
	positions := [][2]int{robotPosition}
	positionsToConsider := map[[2]int]bool{robotPosition: true}
	for {
		newPositions := map[[2]int]bool{}
		empty := true
		for positionToConsider := range positionsToConsider {
			nextPosition := nextPosition(positionToConsider, movement)
			if warehouseMap[nextPosition[0]][nextPosition[1]] == "#" {
				return [][2]int{}
			}
			if warehouseMap[nextPosition[0]][nextPosition[1]] != "." {
				empty = false
				newPositions[nextPosition] = true
			}
			if (movement == "^" || movement == "v") && warehouseMap[nextPosition[0]][nextPosition[1]] == "[" {
				newPositions[[2]int{nextPosition[0], nextPosition[1] + 1}] = true
			} else if (movement == "^" || movement == "v") && warehouseMap[nextPosition[0]][nextPosition[1]] == "]" {
				newPositions[[2]int{nextPosition[0], nextPosition[1] - 1}] = true
			}
		}
		if empty {
			return positions
		}

		for newPosition := range newPositions {
			positions = append(positions, newPosition)
		}
		positionsToConsider = newPositions
	}
}

// sumResizedCoordinates sums the coordinates of the items in the resized warehouse.
func sumResizedCoordinates(warehouseMap [][]string) int {
	total := 0
	for i, row := range warehouseMap {
		for j, position := range row {
			if position == "[" {
				total += i*100 + j
			}
		}
	}
	return total
}

// readLines converts the information from the file into a usable form.
func readLines(fileName string) ([][]string, []string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, nil, err
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)

	warehouseMap, movements := [][]string{}, []string{}
	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}
		warehouseMap = append(warehouseMap, strings.Split(scanner.Text(), ""))
	}
	for scanner.Scan() {
		movements = append(movements, strings.Split(scanner.Text(), "")...)
	}
	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}

	return warehouseMap, movements, nil
}
