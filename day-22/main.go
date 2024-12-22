package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	secretNumbers, err := readLines("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(partOne(secretNumbers))
	fmt.Println(partTwo(secretNumbers))
}

// partOne solves part one of the puzzle.
func partOne(secretNumbers []int) int {
	total := 0
	for _, secretNumber := range secretNumbers {
		for i := 1; i <= 2000; i++ {
			secretNumber = newSecretNumber(secretNumber)
		}
		total += secretNumber
	}
	return total
}

// partTwo solves part two of the puzzle.
func partTwo(secretNumbers []int) int {
	changeMaps := []map[[4]int]int{}
	for _, secretNumber := range secretNumbers {
		changeMap := map[[4]int]int{}
		numbers := []int{secretNumber}
		for i := 1; i <= 2000; i++ {
			numbers = append(numbers, newSecretNumber(numbers[len(numbers)-1]))
			if i >= 4 {
				changes := [4]int{(numbers[i-3] % 10) - (numbers[i-4] % 10), (numbers[i-2] % 10) - (numbers[i-3] % 10),
					(numbers[i-1] % 10) - (numbers[i-2] % 10), (numbers[i] % 10) - (numbers[i-1] % 10)}
				if _, ok := changeMap[changes]; !ok {
					changeMap[changes] = numbers[i] % 10
				}
			}
		}
		changeMaps = append(changeMaps, changeMap)
	}
	overallMap := map[[4]int]int{}
	for _, changeMap := range changeMaps {
		for key, value := range changeMap {
			overallMap[key] +=  value
		}
	}
	max := -1
	for _, value := range overallMap {
		if max == -1 || value > max {
			max = value
		}
	}
	return max
}

// newSecretNumber performs an evolution of a secret number.
func newSecretNumber(secretNumber int) int {
	secretNumber = mixAndPrune(secretNumber, secretNumber*64)
	secretNumber = mixAndPrune(secretNumber, secretNumber/32)
	return mixAndPrune(secretNumber, secretNumber*2048)
}

// mixAndPrune performs the mix and prune operations.
func mixAndPrune(secretNumber, secondValue int) int {
	secretNumber = secretNumber ^ secondValue
	secretNumber = secretNumber % 16777216
	return secretNumber
}

// readLines converts the information from the file into a usable form.
func readLines(fileName string) ([]int, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)

	secretNumbers := []int{}
	for scanner.Scan() {
		number, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}

		secretNumbers = append(secretNumbers, number)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return secretNumbers, nil
}
