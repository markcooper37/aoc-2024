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
	trailMap, err := readLines("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(partOne(trailMap))
	fmt.Println(partTwo(trailMap))
}

// partOne solves part one of the puzzle.
func partOne(trailMap [][]int) int {
	trailHeads := [][2]int{}
	for i, row := range trailMap {
		for j, value := range row {
			if value == 0 {
				trailHeads = append(trailHeads, [2]int{i, j})
			}
		}
	}

	total := 0
	for _, trailHead := range trailHeads {
		total += score(trailHead, trailMap, false)
	}
	return total
}

// partTwo solves part two of the puzzle.
func partTwo(trailMap [][]int) int {
	trailHeads := [][2]int{}
	for i, row := range trailMap {
		for j, value := range row {
			if value == 0 {
				trailHeads = append(trailHeads, [2]int{i, j})
			}
		}
	}

	total := 0
	for _, trailHead := range trailHeads {
		total += score(trailHead, trailMap, true)
	}
	return total
}

// score determines the score for a trailhead.
func score(trailheadPosition [2]int, trailMap [][]int, countAllTrails bool) int {
	trails := [][][2]int{{trailheadPosition}}
	for i := 1; i <= 9; i++ {
		newTrails := [][][2]int{}
		for _, trail := range trails {
			validNeighbours := validNeighbours(trail[len(trail)-1], trailMap)
			for _, validNeighbour := range validNeighbours {
				extendedTrail := make([][2]int, len(trail)+1)
				copy(extendedTrail, trail)
				extendedTrail[len(extendedTrail)-1] = validNeighbour
				newTrails = append(newTrails, extendedTrail)
			}
		}
		trails = newTrails
	}
	if countAllTrails {
		return len(trails)
	} else {
		endPoints := map[[2]int]bool{}
		for _, trail := range trails {
			endPoints[trail[len(trail)-1]] = true
		}
		return len(endPoints)
	}
}

// validNeighbours finds all valid neighbouring positions.
func validNeighbours(position [2]int, trailMap [][]int) [][2]int {
	validNeighours := [][2]int{}
	neighbours := [][2]int{{position[0] + 1, position[1]}, {position[0] - 1, position[1]},
		{position[0], position[1] + 1}, {position[0], position[1] - 1}}
	for _, neighbour := range neighbours {
		if neighbour[0] >= 0 && neighbour[0] < len(trailMap) && neighbour[1] >= 0 && neighbour[1] < len(trailMap[0]) {
			if trailMap[neighbour[0]][neighbour[1]] == trailMap[position[0]][position[1]]+1 {
				validNeighours = append(validNeighours, neighbour)
			}
		}
	}
	return validNeighours
}

// readLines converts the information from the file into a usable form.
func readLines(fileName string) ([][]int, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)

	trailMap := [][]int{}
	for scanner.Scan() {
		valuesStr := strings.Split(scanner.Text(), "")
		values := []int{}
		for _, valueStr := range valuesStr {
			value, err := strconv.Atoi(valueStr)
			if err != nil {
				return nil, err
			}
			values = append(values, value)
		}
		trailMap = append(trailMap, values)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return trailMap, nil
}
