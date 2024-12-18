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
	bytes, err := readLines("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(partOne(bytes, 70, 1024))
	fmt.Println(partTwo(bytes, 70))
}

// partOne solves part one of the puzzle.
func partOne(bytes [][2]int, gridSize int, simulatedBytes int) int {
	grid := constructGrid(gridSize)
	for i := 0; i < simulatedBytes; i++ {
		grid[bytes[i][1]][bytes[i][0]] = "#"
	}
	return *findShortestRouteLength(grid)
}

// partTwo solves part two of the puzzle.
func partTwo(bytes [][2]int, gridSize int) string {
	grid := constructGrid(gridSize)
	for _, byte := range bytes {
		grid[byte[1]][byte[0]] = "#"
		shortestRouteLength := findShortestRouteLength(grid)
		if shortestRouteLength == nil {
			return strconv.Itoa(byte[0]) + "," + strconv.Itoa(byte[1])
		}
	}
	return ""
}

// constructGrid constructs the grid.
func constructGrid(gridSize int) [][]string {
	grid := [][]string{}
	for i := 0; i <= gridSize; i++ {
		row := []string{}
		for j := 0; j <= gridSize; j++ {
			row = append(row, ".")
		}
		grid = append(grid, row)
	}
	return grid
}

// findShortestRouteLength finds the length of the shortest route from the top left to the bottom right.
func findShortestRouteLength(grid [][]string) *int {
	distance := 0
	end := [2]int{len(grid)-1, len(grid)-1}
	allPositions := map[[2]int]bool{{0, 0}: true}
	positionsToConsider := map[[2]int]bool{{0, 0}: true}
	for len(positionsToConsider) > 0 {
		distance++
		newPositions := map[[2]int]bool{}
		for positionToConsider := range positionsToConsider {
			adjacentPositions := adjacentPositions(positionToConsider)
			for _, adjacentPosition := range adjacentPositions {
				if _, ok := allPositions[adjacentPosition]; withinBoundaries(adjacentPosition, grid) &&
					!ok && grid[adjacentPosition[0]][adjacentPosition[1]] == "." {
					if adjacentPosition == end {
						return &distance
					}
					newPositions[adjacentPosition] = true
					allPositions[adjacentPosition] = true
				}
			}
		}
		positionsToConsider = newPositions
	}
	return nil
}

// adjacentPositions gets all adjacent positions.
func adjacentPositions(position [2]int) [][2]int {
	return [][2]int{{position[0] - 1, position[1]}, {position[0], position[1] + 1},
		{position[0] + 1, position[1]}, {position[0], position[1] - 1}}
}

// withinBoundaries checks if the position is within the boundaries of the grid.
func withinBoundaries(position [2]int, grid [][]string) bool {
	return position[0] >= 0 && position[0] < len(grid) && position[1] >= 0 && position[1] < len(grid[0])
}

// readLines converts the information from the file into a usable form.
func readLines(fileName string) ([][2]int, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)

	bytes := [][2]int{}
	for scanner.Scan() {
		splitValues := strings.Split(scanner.Text(), ",")
		first, err := strconv.Atoi(splitValues[0])
		if err != nil {
			return nil, err
		}

		second, err := strconv.Atoi(splitValues[1])
		if err != nil {
			return nil, err
		}

		bytes = append(bytes, [2]int{first, second})
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return bytes, nil
}
