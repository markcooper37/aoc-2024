package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	rows, err := readLines("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(partOne(rows))
	fmt.Println(partTwo(rows))
}

// partOne solves part one of the puzzle.
func partOne(rows [][]string) int {
	total := 0

	// Check horizontal lines
	for i := 0; i < len(rows); i++ {
		for j := 0; j < len(rows[0])-3; j++ {
			if rows[i][j] == "X" && rows[i][j+1] == "M" && rows[i][j+2] == "A" && rows[i][j+3] == "S" {
				total++
			} else if rows[i][j] == "S" && rows[i][j+1] == "A" && rows[i][j+2] == "M" && rows[i][j+3] == "X" {
				total++
			}
		}
	}

	// Check vertical lines
	for j := 0; j < len(rows[0]); j++ {
		for i := 0; i < len(rows)-3; i++ {
			if rows[i][j] == "X" && rows[i+1][j] == "M" && rows[i+2][j] == "A" && rows[i+3][j] == "S" {
				total++
			} else if rows[i][j] == "S" && rows[i+1][j] == "A" && rows[i+2][j] == "M" && rows[i+3][j] == "X" {
				total++
			}
		}
	}

	// Check diagonals
	for i := 0; i < len(rows)-3; i++ {
		for j := 0; j < len(rows[0])-3; j++ {
			if rows[i][j] == "X" && rows[i+1][j+1] == "M" && rows[i+2][j+2] == "A" && rows[i+3][j+3] == "S" {
				total++
			} else if rows[i][j] == "S" && rows[i+1][j+1] == "A" && rows[i+2][j+2] == "M" && rows[i+3][j+3] == "X" {
				total++
			}

			if rows[i][j+3] == "X" && rows[i+1][j+2] == "M" && rows[i+2][j+1] == "A" && rows[i+3][j] == "S" {
				total++
			} else if rows[i][j+3] == "S" && rows[i+1][j+2] == "A" && rows[i+2][j+1] == "M" && rows[i+3][j] == "X" {
				total++
			}
		}
	}

	return total
}

// partTwo solves part two of the puzzle.
func partTwo(rows [][]string) int {
	total := 0
	for i := 0; i < len(rows)-2; i++ {
		for j := 0; j < len(rows[0])-2; j++ {
			if rows[i][j] == "M" && rows[i][j+2] == "M" && rows[i+1][j+1] == "A" && rows[i+2][j] == "S" && rows[i+2][j+2] == "S" {
				total++
			} else if rows[i][j] == "M" && rows[i][j+2] == "S" && rows[i+1][j+1] == "A" && rows[i+2][j] == "M" && rows[i+2][j+2] == "S" {
				total++
			} else if rows[i][j] == "S" && rows[i][j+2] == "M" && rows[i+1][j+1] == "A" && rows[i+2][j] == "S" && rows[i+2][j+2] == "M" {
				total++
			} else if rows[i][j] == "S" && rows[i][j+2] == "S" && rows[i+1][j+1] == "A" && rows[i+2][j] == "M" && rows[i+2][j+2] == "M" {
				total++
			}
		}
	}

	return total
}

// readLines converts the information from the file into a usable form.
func readLines(fileName string) ([][]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)

	rows := [][]string{}
	for scanner.Scan() {
		rows = append(rows, strings.Split(scanner.Text(), ""))
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return rows, nil
}
