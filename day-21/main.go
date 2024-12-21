package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var numericalKeypad = [][]string{{"7", "8", "9"}, {"4", "5", "6"}, {"1", "2", "3"}, {"", "0", "A"}}

func main() {
	codes, err := readLines("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(partOne(codes))
	fmt.Println(partTwo(codes))
}

// partOne solves part one of the puzzle.
func partOne(codes [][]string) int {
	total := 0
	for _, code := range codes {
		shortestSequenceLength := findShortestSequenceLength(code, 2)
		total += shortestSequenceLength * numericalPart(code)
	}
	return total
}

// partTwo solves part two of the puzzle.
func partTwo(codes [][]string) int {
	total := 0
	for _, code := range codes {
		shortestSequenceLength := findShortestSequenceLength(code, 25)
		total += shortestSequenceLength * numericalPart(code)
	}
	return total
}

// findShortestSequenceLength finds the shortest sequence required to input the code.
func findShortestSequenceLength(code []string, directionalRobots int) int {
	length := 0
	current := "A"
	for i := 0; i < len(code); i++ {
		next := code[i]
		length += findMinimumDirectionPresses(current, next, directionalRobots)
		current = next
	}

	return length
}

// findMinimumDirectionPresses finds the fewest presses required to input the next digit.
func findMinimumDirectionPresses(current, next string, directionalRobots int) int {
	directionalSequences := numericalToDirectional(current, next)
	minPresses := -1
	for _, sequence := range directionalSequences {
		pairs := map[[2]string]int{}
		current := "A"
		for i := 0; i < len(sequence); i++ {
			next := sequence[i]
			pairs[[2]string{current, next}]++
			current = next
		}
		for i := 0; i < directionalRobots; i++ {
			newPairs := map[[2]string]int{}
			for pair, count := range pairs {
				newDirectionalSequence := directionalToDirectional(pair[0], pair[1])
				current = "A"
				for i := 0; i < len(newDirectionalSequence); i++ {
					next := newDirectionalSequence[i]
					newPairs[[2]string{current, next}] += count
					current = next
				}
			}
			pairs = newPairs
		}
		pressCount := 0
		for _, count := range pairs {
			pressCount += count
		}
		if minPresses == -1 || pressCount < minPresses {
			minPresses = pressCount
		}
	}

	return minPresses
}

// directionalToDirectional finds the best directional sequence that gets from current to next on the directional keypad.
func directionalToDirectional(current, next string) []string {
	switch current {
	case "A":
		switch next {
		case "A":
			return []string{"A"}
		case "<":
			return []string{"v", "<", "<", "A"}
		case ">":
			return []string{"v", "A"}
		case "v":
			return []string{"<", "v", "A"}
		case "^":
			return []string{"<", "A"}
		}
	case "<":
		switch next {
		case "A":
			return []string{">", ">", "^", "A"}
		case "<":
			return []string{"A"}
		case ">":
			return []string{">", ">", "A"}
		case "v":
			return []string{">", "A"}
		case "^":
			return []string{">", "^", "A"}
		}
	case ">":
		switch next {
		case "A":
			return []string{"^", "A"}
		case "<":
			return []string{"<", "<", "A"}
		case ">":
			return []string{"A"}
		case "v":
			return []string{"<", "A"}
		case "^":
			return []string{"<", "^", "A"}
		}
	case "v":
		switch next {
		case "A":
			return []string{"^", ">", "A"}
		case "<":
			return []string{"<", "A"}
		case ">":
			return []string{">", "A"}
		case "v":
			return []string{"A"}
		case "^":
			return []string{"^", "A"}
		}
	case "^":
		switch next {
		case "A":
			return []string{">", "A"}
		case "<":
			return []string{"v", "<", "A"}
		case ">":
			return []string{"v", ">", "A"}
		case "v":
			return []string{"v", "A"}
		case "^":
			return []string{"A"}
		}
	}
	return nil
}

// numericalToDirectional finds all directional sequences that get from current to next on the numerical keypad.
func numericalToDirectional(current, next string) [][]string {
	currentPosition := positionOnNumericalKeypad(current)
	nextPosition := positionOnNumericalKeypad(next)
	journeys := findJourneysBetweenPositions(currentPosition, nextPosition)
	directional := [][]string{}
	for _, journey := range journeys {
		if validNumericalJourney(journey) {
			arrows := convertJourneyToArrows(journey)
			arrows = append(arrows, "A")
			directional = append(directional, arrows)
		}
	}
	return directional
}

// positionOnNumericalKeypad finds the position of a character on the numerical keypad.
func positionOnNumericalKeypad(value string) [2]int {
	for i, row := range numericalKeypad {
		for j, position := range row {
			if position == value {
				return [2]int{i, j}
			}
		}
	}
	return [2]int{-1, -1}
}

// findJourneysBetweenPositions finds all possible journeys between positions.
func findJourneysBetweenPositions(start, end [2]int) [][][2]int {
	totalDifference := difference(start[0], end[0]) + difference(start[1], end[1])
	journeys := [][][2]int{{start}}
	for i := 1; i <= totalDifference; i++ {
		newJourneys := [][][2]int{}
		for _, journey := range journeys {
			lastPosition := journey[len(journey)-1]
			if lastPosition[0] < end[0] {
				journeyCopy := copyJourney(journey)
				journeyCopy = append(journeyCopy, [2]int{lastPosition[0] + 1, lastPosition[1]})
				newJourneys = append(newJourneys, journeyCopy)
			}
			if lastPosition[0] > end[0] {
				journeyCopy := copyJourney(journey)
				journeyCopy = append(journeyCopy, [2]int{lastPosition[0] - 1, lastPosition[1]})
				newJourneys = append(newJourneys, journeyCopy)
			}
			if lastPosition[1] < end[1] {
				journeyCopy := copyJourney(journey)
				journeyCopy = append(journeyCopy, [2]int{lastPosition[0], lastPosition[1] + 1})
				newJourneys = append(newJourneys, journeyCopy)
			}
			if lastPosition[1] > end[1] {
				journeyCopy := copyJourney(journey)
				journeyCopy = append(journeyCopy, [2]int{lastPosition[0], lastPosition[1] - 1})
				newJourneys = append(newJourneys, journeyCopy)
			}
		}
		journeys = newJourneys
	}
	return journeys
}

// validNumericalJourney checks whether a journey on a numerical keypad is valid.
func validNumericalJourney(positions [][2]int) bool {
	for _, position := range positions {
		if numericalKeypad[position[0]][position[1]] == "" {
			return false
		}
	}
	return true
}

// copyJourney creates a copy of a journey.
func copyJourney(journey [][2]int) [][2]int {
	newJourney := make([][2]int, len(journey))
	copy(newJourney, journey)
	return newJourney
}

// convertJourneyToArrows converts a journey to arrows.
func convertJourneyToArrows(journey [][2]int) []string {
	arrows := []string{}
	for i := 1; i < len(journey); i++ {
		rowDiff, colDiff := journey[i][0]-journey[i-1][0], journey[i][1]-journey[i-1][1]
		if rowDiff == 1 {
			arrows = append(arrows, "v")
		} else if rowDiff == -1 {
			arrows = append(arrows, "^")
		} else if colDiff == 1 {
			arrows = append(arrows, ">")
		} else {
			arrows = append(arrows, "<")
		}
	}
	return arrows
}

// difference finds the difference between two integers.
func difference(x, y int) int {
	diff := x - y
	if diff < 0 {
		return -diff
	}
	return diff
}

// numericalPart finds the numerical part of a code.
func numericalPart(code []string) int {
	numbersStr := code[:len(code)-1]
	numericalPart := 0
	for _, numberStr := range numbersStr {
		number, err := strconv.Atoi(numberStr)
		if err != nil {
			log.Fatal(err)
		}

		numericalPart *= 10
		numericalPart += number
	}
	return numericalPart
}

// readLines converts the information from the file into a usable form.
func readLines(fileName string) ([][]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)

	codes := [][]string{}
	for scanner.Scan() {
		codes = append(codes, strings.Split(scanner.Text(), ""))
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return codes, nil
}
