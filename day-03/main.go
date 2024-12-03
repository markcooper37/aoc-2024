package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	reports, err := readLines("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(partOne(reports))
	fmt.Println(partTwo(reports))
}

// partOne solves part one of the puzzle.
func partOne(lines []string) (int, error) {
	regex := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
	total := 0
	for _, line := range lines {
		matches := regex.FindAllString(line, -1)
		for _, match := range matches {
			vals := strings.Split(strings.TrimLeft(strings.TrimRight(match, ")"), "mul("), ",")
			firstVal, err := strconv.Atoi(vals[0])
			if err != nil {
				return 0, err
			}

			secondVal, err := strconv.Atoi(vals[1])
			if err != nil {
				return 0, err
			}

			total += firstVal * secondVal
		}
	}
	return total, nil
}

// partTwo solves part two of the puzzle.
func partTwo(lines []string) (int, error) {
	regex := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)|do\(\)|don't\(\)`)
	total, do := 0, true
	for _, line := range lines {
		matches := regex.FindAllString(line, -1)
		for _, match := range matches {
			if match == "do()" {
				do = true
				continue
			} else if match == "don't()" {
				do = false
				continue
			} else if do {
				vals := strings.Split(strings.TrimLeft(strings.TrimRight(match, ")"), "mul("), ",")
				firstVal, err := strconv.Atoi(vals[0])
				if err != nil {
					return 0, err
				}

				secondVal, err := strconv.Atoi(vals[1])
				if err != nil {
					return 0, err
				}

				total += firstVal * secondVal
			}
		}
	}
	return total, nil
}

// readLines converts the information from the file into a usable form.
func readLines(fileName string) ([]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)

	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}
