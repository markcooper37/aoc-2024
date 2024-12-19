package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	patterns, designs, err := readLines("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(partOne(patterns, designs))
	fmt.Println(partTwo(patterns, designs))
}

// partOne solves part one of the puzzle.
func partOne(patterns, designs []string) int {
	total := 0
	for _, design := range designs {
		if designPossible(design, patterns) {
			total++
		}
	}
	return total
}

// partTwo solves part two of the puzzle.
func partTwo(patterns, designs []string) int {
	total := 0
	for _, design := range designs {
		total += countArrangements(design, patterns)
	}
	return total
}

// designPossible checks whether a design can be made with the given patterns.
func designPossible(design string, patterns []string) bool {
	remainders := map[string]bool{design: true}
	for len(remainders) > 0 {
		newRemainders := map[string]bool{}
		for remainder := range remainders {
			for _, pattern := range patterns {
				if newRemainder := getRemainder(remainder, pattern); newRemainder != nil {
					if *newRemainder == "" {
						return true
					} else {
						newRemainders[*newRemainder] = true
					}
				}
			}
		}
		remainders = newRemainders
	}
	return false
}

// getRemainder returns the remainder of the design that appears after the pattern, or nil if the pattern
// does not appear at the start of the design.
func getRemainder(design, pattern string) *string {
	if len(design) < len(pattern) {
		return nil
	}
	for i := 0; i < len(pattern); i++ {
		if design[i] != pattern[i] {
			return nil
		}
	}
	remainder := design[len(pattern):]
	return &remainder
}

// countArrangements counts all the ways the design can be made with the patterns.
func countArrangements(design string, patterns []string) int {
	remainders := map[string]int{design: 1}
	total := 0
	for len(remainders) > 0 {
		newRemainders := map[string]int{}
		for remainder, count := range remainders {
			for _, pattern := range patterns {
				if newRemainder := getRemainder(remainder, pattern); newRemainder != nil {
					if *newRemainder == "" {
						total += count
					} else {
						newRemainders[*newRemainder] += count
					}
				}
			}
		}
		remainders = newRemainders
	}
	return total
}

// readLines converts the information from the file into a usable form.
func readLines(fileName string) ([]string, []string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, nil, err
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	patterns := strings.Split(scanner.Text(), ", ")
	scanner.Scan()
	designs := []string{}
	for scanner.Scan() {
		designs = append(designs, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}

	return patterns, designs, nil
}
