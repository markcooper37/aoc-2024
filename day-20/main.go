package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	racetrack, err := readLines("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(partOne(racetrack, 100))
	fmt.Println(partTwo(racetrack, 100))
}

// partOne solves part one of the puzzle.
func partOne(racetrack [][]string, picosecondsToSave int) int {
	route := constructRoute(racetrack)
	total := 0
	for position, picoseconds := range route {
		adjacentToPosition := adjacentPositions(position)
		for _, adjacentPosition := range adjacentToPosition {
			if racetrack[adjacentPosition[0]][adjacentPosition[1]] == "#" {
				adjacentToWall := adjacentPositions(adjacentPosition)
				for _, adjacentPositionToWall := range adjacentToWall {
					if withinBoundaries(adjacentPositionToWall, racetrack) && adjacentPositionToWall != position &&
						racetrack[adjacentPositionToWall[0]][adjacentPositionToWall[1]] != "#" {
						timeSaved := route[adjacentPositionToWall] - picoseconds - 2
						if timeSaved >= picosecondsToSave {
							total++
						}
					}
				}
			}
		}
	}
	return total
}

// partTwo solves part two of the puzzle.
func partTwo(racetrack [][]string, picosecondsToSave int) int {
	route := constructRoute(racetrack)
	total := 0
	for position, picoseconds := range route {
		cheatEnds := map[[2]int]bool{}
		positionsToConsider := map[[2]int]bool{position: true}
		for i := 1; i <= 20; i++ {
			allAdjacentPositions := map[[2]int]bool{}
			for position := range positionsToConsider {
				adjacentToPosition := adjacentPositions(position)
				for _, adjacentPosition := range adjacentToPosition {
					allAdjacentPositions[adjacentPosition] = true
				}
			}
			newPositionsToConsider := map[[2]int]bool{}
			for adjacentPosition := range allAdjacentPositions {
				if !withinBoundaries(adjacentPosition, racetrack) {
					continue
				}
				newPositionsToConsider[adjacentPosition] = true
				if !cheatEnds[adjacentPosition] {
					cheatEnds[adjacentPosition] = true
					timeSaved := route[adjacentPosition] - picoseconds - i
					if timeSaved >= picosecondsToSave {
						total++
					}
				}
			}
			positionsToConsider = newPositionsToConsider
		}
	}
	return total
}

// constructRoute constructs a route from the start to the end.
func constructRoute(racetrack [][]string) map[[2]int]int {
	start := findPosition("S", racetrack)
	route := map[[2]int]int{start: 0}
	newPosition := start
	end := findPosition("E", racetrack)
	for {
		adjacentPositions := adjacentPositions(newPosition)
		for _, adjacentPosition := range adjacentPositions {
			if _, ok := route[adjacentPosition]; !ok &&
				(racetrack[adjacentPosition[0]][adjacentPosition[1]] == "." || racetrack[adjacentPosition[0]][adjacentPosition[1]] == "E") {
				route[adjacentPosition] = route[newPosition] + 1
				if adjacentPosition == end {
					return route
				}
				newPosition = adjacentPosition
				break
			}
		}
	}
}

// findPosition finds the desired position.
func findPosition(position string, racetrack [][]string) [2]int {
	for i, row := range racetrack {
		for j, pos := range row {
			if pos == position {
				return [2]int{i, j}
			}
		}
	}
	return [2]int{-1, -1}
}

// adjacentPositions gets all adjacent positions.
func adjacentPositions(position [2]int) [][2]int {
	return [][2]int{{position[0] - 1, position[1]}, {position[0], position[1] + 1},
		{position[0] + 1, position[1]}, {position[0], position[1] - 1}}
}

// withinBoundaries checks if the position is within the boundaries of the racetrack.
func withinBoundaries(position [2]int, racetrack [][]string) bool {
	return position[0] >= 0 && position[0] < len(racetrack) && position[1] >= 0 && position[1] < len(racetrack[0])
}

// readLines converts the information from the file into a usable form.
func readLines(fileName string) ([][]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)

	racetrack := [][]string{}
	for scanner.Scan() {
		racetrack = append(racetrack, strings.Split(scanner.Text(), ""))
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return racetrack, nil
}
