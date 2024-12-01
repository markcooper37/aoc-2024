package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	firstColumn, secondColumn, err := readLines("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(partOne(firstColumn, secondColumn))
	fmt.Println(partTwo(firstColumn, secondColumn))
}

// partOne solves part one of the puzzle.
func partOne(firstColumn, secondColumn []int) int {
	slices.Sort(firstColumn)
	slices.Sort(secondColumn)
	total := 0
	for i, value := range firstColumn {
		total += difference(value, secondColumn[i])
	}
	return total
}

// partTwo solves part two of the puzzle.
func partTwo(firstColumn, secondColumn []int) int {
	firstColumnMap, secondColumnMap := map[int]int{}, map[int]int{}
	for _, value := range firstColumn {
		firstColumnMap[value]++
	}

	for _, value := range secondColumn {
		secondColumnMap[value]++
	}

	total := 0
	for key, value := range firstColumnMap {
		total += key * value * secondColumnMap[key]
	}
	return total
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
func readLines(fileName string) ([]int, []int, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, nil, err
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)

	firstColumn, secondColumn := []int{}, []int{}
	for scanner.Scan() {
		pair := strings.Split(scanner.Text(), "   ")
		first, err :=  strconv.Atoi(pair[0])
		if err != nil {
			return nil, nil, err
		}

		firstColumn = append(firstColumn, first)

		second, err :=  strconv.Atoi(pair[1])
		if err != nil {
			return nil, nil, err
		}

		secondColumn = append(secondColumn, second)
	}
	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}

	return firstColumn, secondColumn, nil
}
