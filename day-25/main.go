package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	schematics, err := readLines("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(partOne(schematics))
}

// partOne solves part one of the puzzle.
func partOne(schematics [][][]string) int {
	locks := [][][]string{}
	keys := [][][]string{}
	for _, schematic := range schematics {
		if schematic[0][0] == "#" {
			locks = append(locks, schematic)
		} else {
			keys = append(keys, schematic)
		}
	}

	allLockHeights := [][]int{}
	for _, lock := range locks {
		allLockHeights = append(allLockHeights, getLockHeights(lock))
	}

	allKeyHeights := [][]int{}
	for _, key := range keys {
		allKeyHeights = append(allKeyHeights, getKeyHeights(key))
	}

	total := 0
	for _, lockHeights := range allLockHeights {
		for _, keyHeights := range allKeyHeights {
			if !overlap(lockHeights, keyHeights) {
				total++
			}
		}
	}
	return total
}

func getKeyHeights(key [][]string) []int {
	heights := []int{}
	for j := range key[0] {
		for i := len(key) - 2; i >= 0; i-- {
			if key[i][j] == "." {
				heights = append(heights, len(key)-2-i)
				break
			}
		}
	}
	return heights
}

func getLockHeights(key [][]string) []int {
	heights := []int{}
	for j := range key[0] {
		for i := 1; i < len(key); i++ {
			if key[i][j] == "." {
				heights = append(heights, i-1)
				break
			}
		}
	}
	return heights
}

func overlap(lockHeights, keyHeights []int) bool {
	for i, height := range lockHeights {
		if height + keyHeights[i] > 5 {
			return true
		}
	}
	return false
}

// readLines converts the information from the file into a usable form.
func readLines(fileName string) ([][][]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)

	schematics := [][][]string{}
	newSchematic := [][]string{}
	for scanner.Scan() {
		if scanner.Text() == "" {
			continue
		}
		newSchematic = append(newSchematic, strings.Split(scanner.Text(), ""))
		if len(newSchematic) == 7 {
			schematics = append(schematics, newSchematic)
			newSchematic = [][]string{}
		}

	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return schematics, nil
}
