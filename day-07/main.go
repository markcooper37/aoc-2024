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
	equations, err := readLines("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(partOne(equations))
	fmt.Println(partTwo(equations))
}

type Equation struct {
	Value   int
	Numbers []int
}

// partOne solves part one of the puzzle.
func partOne(equations []Equation) int {
	total := 0
	for _, equation := range equations {
		if canSatisfyEquationPartOne(equation.Numbers[0], equation.Value, equation.Numbers[1:]) {
			total += equation.Value
		}
	}
	return total
}

// partTwo solves part two of the puzzle.
func partTwo(equations []Equation) int {
	total := 0
	for _, equation := range equations {
		if canSatisfyEquationPartTwo(equation.Numbers[0], equation.Value, equation.Numbers[1:]) {
			total += equation.Value
		}
	}
	return total
}

// canSatisfyEquationPartOne checks whether an equation can be satisfied by inserting * or + operators.
func canSatisfyEquationPartOne(currentValue, targetValue int, remainingNumbers []int) bool {
	if len(remainingNumbers) == 0 {
		return currentValue == targetValue
	}
	val1 := currentValue * remainingNumbers[0]
	val2 := currentValue + remainingNumbers[0]
	return canSatisfyEquationPartOne(val1, targetValue, remainingNumbers[1:]) || canSatisfyEquationPartOne(val2, targetValue, remainingNumbers[1:])
}

// canSatisfyEquationPartTwo checks whether an equation can be satisfied by inserting *, + or || operators.
func canSatisfyEquationPartTwo(currentValue, targetValue int, remainingNumbers []int) bool {
	if len(remainingNumbers) == 0 {
		return currentValue == targetValue
	}
	val1 := currentValue * remainingNumbers[0]
	val2 := currentValue + remainingNumbers[0]
	digits := countDigits(remainingNumbers[0])
	val3 := currentValue
	for i := 0; i < digits; i++ {
		val3 = val3 * 10
	}
	val3 += remainingNumbers[0]
	return canSatisfyEquationPartTwo(val1, targetValue, remainingNumbers[1:]) || canSatisfyEquationPartTwo(val2, targetValue, remainingNumbers[1:]) ||
		canSatisfyEquationPartTwo(val3, targetValue, remainingNumbers[1:])
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

// readLines converts the information from the file into a usable form.
func readLines(fileName string) ([]Equation, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)

	equations := []Equation{}
	for scanner.Scan() {
		equationStr := strings.Split(scanner.Text(), ": ")
		equation := Equation{}
		value, err := strconv.Atoi(equationStr[0])
		if err != nil {
			return nil, err
		}

		equation.Value = value
		numbers := []int{}
		numbersStr := strings.Split(equationStr[1], " ")
		for _, numberStr := range numbersStr {
			number, err := strconv.Atoi(numberStr)
			if err != nil {
				return nil, err
			}

			numbers = append(numbers, number)
		}
		equation.Numbers = numbers
		equations = append(equations, equation)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return equations, nil
}
