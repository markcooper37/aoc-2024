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
	rules, updates, err := readLines("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(partOne(rules, updates))
	fmt.Println(partTwo(rules, updates))
}

// partOne solves part one of the puzzle.
func partOne(rules [][2]int, updates [][]int) int {
	ruleMap := map[[2]int]bool{}
	for _, rule := range rules {
		ruleMap[rule] = true
	}
	total := 0
	for _, update := range updates {
		if updateValid(ruleMap, update) {
			total += update[(len(update)-1)/2]
		}
	}
	return total
}

// partTwo solves part two of the puzzle.
func partTwo(rules [][2]int, updates [][]int) int {
	ruleMap := map[[2]int]bool{}
	for _, rule := range rules {
		ruleMap[rule] = true
	}
	total := 0
	for _, update := range updates {
		if !updateValid(ruleMap, update) {
			reorderedUpdate := reorderUpdate(ruleMap, update)
			total += reorderedUpdate[(len(reorderedUpdate)-1)/2]
		}
	}
	return total
}

// updateValid checks whether an update satisfies all rules.
func updateValid(ruleMap map[[2]int]bool, update []int) bool {
	for i := 0; i < len(update)-1; i++ {
		for j := i + 1; j < len(update); j++ {
			if ruleMap[[2]int{update[j], update[i]}] {
				return false
			}
		}
	}
	return true
}

// reorderUpdate correctly orders an update.
func reorderUpdate(ruleMap map[[2]int]bool, update []int) []int {
	newUpdate := make([]int, len(update))
	copy(newUpdate, update)
	slices.SortFunc(newUpdate, func(i, j int) int {
		if ruleMap[[2]int{i, j}] {
			return -1
		} else if ruleMap[[2]int{j, i}] {
			return 1
		} else {
			return 0
		}
	})
	return newUpdate
}

// readLines converts the information from the file into a usable form.
func readLines(fileName string) ([][2]int, [][]int, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, nil, err
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)

	rules, updates := [][2]int{}, [][]int{}
	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}
		ruleStr := strings.Split(scanner.Text(), "|")
		first, err := strconv.Atoi(ruleStr[0])
		if err != nil {
			return nil, nil, err
		}

		second, err := strconv.Atoi(ruleStr[1])
		if err != nil {
			return nil, nil, err
		}

		rules = append(rules, [2]int{first, second})
	}
	for scanner.Scan() {
		updateStr := strings.Split(scanner.Text(), ",")
		update := []int{}
		for _, valueStr := range updateStr {
			value, err := strconv.Atoi(valueStr)
			if err != nil {
				return nil, nil, err
			}

			update = append(update, value)
		}
		updates = append(updates, update)
	}
	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}

	return rules, updates, nil
}
