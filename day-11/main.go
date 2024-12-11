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
	stones, err := readLines("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(partOne(stones))
	fmt.Println(partTwo(stones))
}

// partOne solves part one of the puzzle.
func partOne(stones []int) int {
	return iterateStones(stones, 25)
}

// partTwo solves part two of the puzzle.
func partTwo(stones []int) int {
	return iterateStones(stones, 75)
}

// iterateStones performs iterations on the stones and returns the final stone count
func iterateStones(stones []int, iterations int) int {
	stonesMap := map[int]int{}
	for _, stone := range stones {
		stonesMap[stone]++
	}
	for i := 1; i <= iterations; i++ {
		newStonesMap := map[int]int{}
		for stone, count := range stonesMap {
			if stone == 0 {
				newStonesMap[1] += count
			} else if digitCount := countDigits(stone); digitCount%2 == 0 {
				stone1, stone2 := splitNumber(stone)
				newStonesMap[stone1] += count
				newStonesMap[stone2] += count
			} else {
				newStonesMap[stone*2024] += count
			}
		}
		stonesMap = newStonesMap
	}
	total := 0
	for _, count := range stonesMap {
		total += count
	}
	return total
}

// countDigits counts the number of digits that a number has.
func countDigits(number int) int {
	digits := 1
	for {
		number = number / 10
		if number == 0 {
			return digits
		}
		digits++
	}
}

// splitNumber splits a number in two.
func splitNumber(number int) (int, int) {
	digitCount := countDigits(number)
	powerOfTen := 10
	for i := 1; i < digitCount/2; i++ {
		powerOfTen *= 10
	}
	return number / powerOfTen, number % powerOfTen
}

// readLines converts the information from the file into a usable form.
func readLines(fileName string) ([]int, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)

	stones := []int{}
	for scanner.Scan() {
		stonesStr := strings.Split(scanner.Text(), " ")
		for _, stoneStr := range stonesStr {
			stone, err := strconv.Atoi(stoneStr)
			if err != nil {
				return nil, err
			}
			stones = append(stones, stone)
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return stones, nil
}
