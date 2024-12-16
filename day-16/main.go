package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	maze, err := readLines("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(partOne(maze))
	fmt.Println(partTwo(maze))
}

type Location struct {
	Position  [2]int
	Direction int // 0-3 from north clockwise to west
}

type Route struct {
	Positions        map[[2]int]bool
	CurrentPosition  [2]int
	CurrentDirection int
	Points           int
}

// partOne solves part one of the puzzle.
func partOne(maze [][]string) int {
	startLocation := Location{
		Position:  findPosition(maze, "S"),
		Direction: 1,
	}
	lowestPoints := lowestPoints(startLocation, maze)
	return lowestPointsAtPosition(lowestPoints, findPosition(maze, "E"))
}

// partTwo solves part two of the puzzle.
func partTwo(maze [][]string) int {
	startPosition := findPosition(maze, "S")
	startLocation := Location{
		Position:  startPosition,
		Direction: 1,
	}
	lowestPoints := lowestPoints(startLocation, maze)
	startRoute := Route{
		Positions:        map[[2]int]bool{startPosition: true},
		CurrentPosition:  startPosition,
		CurrentDirection: 1,
		Points:           0,
	}
	bestRoutes := findBestRoutes(startRoute, lowestPoints, maze)
	positions := map[[2]int]bool{}
	for _, route := range bestRoutes {
		for position := range route.Positions {
			positions[position] = true
		}
	}
	return len(positions)
}

// lowestPoints finds the lowest score achievable for each location.
func lowestPoints(startLocation Location, maze [][]string) map[Location]int {
	endPosition := findPosition(maze, "E")
	locationsToConsider := []Location{startLocation}
	bestPoints := map[Location]int{startLocation: 0}
	for len(locationsToConsider) > 0 {
		newLocationsToConsider := []Location{}
		for _, location := range locationsToConsider {
			adjacentPositons := adjacentPositions(location.Position)
			for i, adjacentPosition := range adjacentPositons {
				if maze[adjacentPosition[0]][adjacentPosition[1]] != "#" {
					newPoints := bestPoints[location] + 1 + 1000*rotationsRequired(location.Direction, i)
					if best, ok := bestPoints[Location{Position: adjacentPosition, Direction: i}]; !ok || newPoints < best {
						bestPoints[Location{Position: adjacentPosition, Direction: i}] = newPoints
						if adjacentPosition != endPosition {
							newLocationsToConsider = append(newLocationsToConsider, Location{Position: adjacentPosition, Direction: i})
						}
					}
				}
			}
		}
		locationsToConsider = newLocationsToConsider
	}
	return bestPoints
}

// lowestPointsAtPosition finds the lowest possible points needed to reach a certain position.
func lowestPointsAtPosition(lowestPoints map[Location]int, position [2]int) int {
	lowestScore := -1
	for i := 0; i <= 3; i++ {
		if val, ok := lowestPoints[Location{Position: position, Direction: i}]; ok && (lowestScore == -1 || val < lowestScore) {
			lowestScore = val
		}
	}
	return lowestScore
}

// findPosition finds a position in a maze.
func findPosition(maze [][]string, desiredPosition string) [2]int {
	for i, row := range maze {
		for j, position := range row {
			if position == desiredPosition {
				return [2]int{i, j}
			}
		}
	}
	return [2]int{-1, -1}
}

// rotationsRequired calculates the rotations needed to switch to the desired direction.
func rotationsRequired(currentDirection, newDirection int) int {
	difference := abs(newDirection - currentDirection)
	if difference == 3 {
		return 1
	}
	return difference
}

// abd calculates the absolute value of an integer.
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// adjacentPositions returns all adjacent positions.
func adjacentPositions(position [2]int) [][2]int {
	return [][2]int{{position[0] - 1, position[1]}, {position[0], position[1] + 1},
		{position[0] + 1, position[1]}, {position[0], position[1] - 1}}
}

// findBestRoutes constructs the bestRoutes from the given start route.
func findBestRoutes(route Route, lowestPoints map[Location]int, maze [][]string) []Route {
	routesToConsider := []Route{route}
	bestRoutes := []Route{}
	endPosition := findPosition(maze, "E")
	for len(routesToConsider) > 0 {
		newRoutesToConsider := []Route{}
		for _, routeToConsider := range routesToConsider {
			adjacentPositions := adjacentPositions(routeToConsider.CurrentPosition)
			for i, adjacentPosition := range adjacentPositions {
				if maze[adjacentPosition[0]][adjacentPosition[1]] != "#" {
					newPoints := routeToConsider.Points + 1 + 1000*rotationsRequired(routeToConsider.CurrentDirection, i)
					if newPoints == lowestPoints[Location{Position: adjacentPosition, Direction: i}] {
						newPositions := copyMap(routeToConsider.Positions)
						newPositions[adjacentPosition] = true
						newRoute := Route{
							Positions:        newPositions,
							CurrentPosition:  adjacentPosition,
							CurrentDirection: i,
							Points:           newPoints,
						}
						if adjacentPosition != endPosition {
							newRoutesToConsider = append(newRoutesToConsider, newRoute)
						} else if newPoints == lowestPointsAtPosition(lowestPoints, endPosition) {
							bestRoutes = append(bestRoutes, newRoute)
						}

					}
				}
			}
		}
		routesToConsider = newRoutesToConsider
	}
	return bestRoutes
}

// copyMap creates a copy of a map.
func copyMap(oldMap map[[2]int]bool) map[[2]int]bool {
	newMap := map[[2]int]bool{}
	for k, v := range oldMap {
		newMap[k] = v
	}
	return newMap
}

// readLines converts the information from the file into a usable form.
func readLines(fileName string) ([][]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)

	maze := [][]string{}
	for scanner.Scan() {
		maze = append(maze, strings.Split(scanner.Text(), ""))
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return maze, nil
}
