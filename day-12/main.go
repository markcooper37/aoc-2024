package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	gardenMap, err := readLines("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(partOne(gardenMap))
	fmt.Println(partTwo(gardenMap))
}

// partOne solves part one of the puzzle.
func partOne(gardenMap [][]string) int {
	consideredPlots := map[[2]int]bool{}
	total := 0
	for i, row := range gardenMap {
		for j := range row {
			if consideredPlots[[2]int{i, j}] {
				continue
			}
			region := constructRegion([2]int{i, j}, gardenMap)
			for plot := range region {
				consideredPlots[plot] = true
			}
			total += perimeter(region) * len(region)
		}
	}
	return total
}

// partTwo solves part two of the puzzle.
func partTwo(gardenMap [][]string) int {
	consideredPlots := map[[2]int]bool{}
	total := 0
	for i, row := range gardenMap {
		for j := range row {
			if consideredPlots[[2]int{i, j}] {
				continue
			}
			region := constructRegion([2]int{i, j}, gardenMap)
			for plot := range region {
				consideredPlots[plot] = true
			}
			total += sides(region) * len(region)
		}
	}
	return total
}

// constructRegion creates a region of all connected plots of the same type.
func constructRegion(plot [2]int, gardenMap [][]string) map[[2]int]bool {
	region := map[[2]int]bool{plot: true}
	plotsToConsider := map[[2]int]bool{plot: true}
	for len(plotsToConsider) > 0 {
		newPlots := map[[2]int]bool{}
		for plot := range plotsToConsider {
			adjacentPlots := adjacentPlots(plot, gardenMap)
			for _, adjacentPlot := range adjacentPlots {
				if gardenMap[adjacentPlot[0]][adjacentPlot[1]] == gardenMap[plot[0]][plot[1]] {
					if !region[adjacentPlot] {
						newPlots[adjacentPlot] = true
					}
					region[adjacentPlot] = true
				}
			}
		}
		plotsToConsider = newPlots
	}
	return region
}

// adjacentPlots gets all adjacent plots.
func adjacentPlots(plot [2]int, gardenMap [][]string) [][2]int {
	adjacentPlots := [][2]int{}
	possibilities := adjacentPositions(plot)
	for _, possibility := range possibilities {
		if possibility[0] >= 0 && possibility[0] < len(gardenMap) && possibility[1] >= 0 && possibility[1] < len(gardenMap[0]) {
			adjacentPlots = append(adjacentPlots, possibility)
		}
	}
	return adjacentPlots
}

// adjacentPositions gets all adjacent positions in the order t (top), r (right), b (bottom), l (left).
func adjacentPositions(position [2]int) [][2]int {
	return [][2]int{{position[0] - 1, position[1]}, {position[0], position[1] + 1}, {position[0] + 1, position[1]}, {position[0], position[1] - 1}}
}

// perimeter calculates the perimeter of a region.
func perimeter(region map[[2]int]bool) int {
	perimeter := 0
	for plot := range region {
		adjacentPositions := adjacentPositions(plot)
		for _, adjacentPosition := range adjacentPositions {
			if !region[adjacentPosition] {
				perimeter++
			}
		}
	}
	return perimeter
}

type Fence struct {
	Plot      [2]int
	Direction int // top: 0, right: 1, bottom: 2, left: 3
}

// sides counts the sides of a region.
func sides(region map[[2]int]bool) int {
	fences := map[Fence]bool{}
	for plot := range region {
		adjacentPositions := adjacentPositions(plot)
		for i, adjacentPosition := range adjacentPositions {
			if !region[adjacentPosition] {
				fences[Fence{Plot: plot, Direction: i}] = true
			}
		}
	}
	sides := 0
	consideredFences := map[Fence]bool{}
	for fence := range fences {
		if consideredFences[fence] {
			continue
		}
		sides++
		consideredFences[fence] = true
		fencesToConsider := map[Fence]bool{fence: true}
		for len(fencesToConsider) > 0 {
			newFences := map[Fence]bool{}
			for fenceToConsider := range fencesToConsider {
				if fenceToConsider.Direction == 0 || fenceToConsider.Direction == 2 {
					adjacent1 := Fence{Plot: [2]int{fenceToConsider.Plot[0], fenceToConsider.Plot[1] - 1}, Direction: fenceToConsider.Direction}
					if fences[adjacent1] {
						if !consideredFences[adjacent1] {
							newFences[adjacent1] = true
						}
						consideredFences[adjacent1] = true
					}
					adjacent2 := Fence{Plot: [2]int{fenceToConsider.Plot[0], fenceToConsider.Plot[1] + 1}, Direction: fenceToConsider.Direction}
					if fences[adjacent2] {
						if !consideredFences[adjacent2] {
							newFences[adjacent2] = true
						}
						consideredFences[adjacent2] = true
					}
				} else {
					adjacent1 := Fence{Plot: [2]int{fenceToConsider.Plot[0] - 1, fenceToConsider.Plot[1]}, Direction: fenceToConsider.Direction}
					if fences[adjacent1] {
						if !consideredFences[adjacent1] {
							newFences[adjacent1] = true
						}
						consideredFences[adjacent1] = true
					}
					adjacent2 := Fence{Plot: [2]int{fenceToConsider.Plot[0] + 1, fenceToConsider.Plot[1]}, Direction: fenceToConsider.Direction}
					if fences[adjacent2] {
						if !consideredFences[adjacent2] {
							newFences[adjacent2] = true
						}
						consideredFences[adjacent2] = true
					}
				}
			}
			fencesToConsider = newFences
		}
	}
	return sides
}

// readLines converts the information from the file into a usable form.
func readLines(fileName string) ([][]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)

	gardenMap := [][]string{}
	for scanner.Scan() {
		gardenMap = append(gardenMap, strings.Split(scanner.Text(), ""))
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return gardenMap, nil
}
