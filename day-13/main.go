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
	machines, err := readLines("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(partOne(machines))
	fmt.Println(partTwo(machines))
}

// partOne solves part one of the puzzle.
func partOne(machines []Machine) int {
	total := 0
	for _, machine := range machines {
		minTokens := 401
		for i := 0; i <= 100; i++ {
			for j := 0; j <= 100; j++ {
				if i*machine.ButtonAMovements[0]+j*machine.ButtonBMovements[0] == machine.PrizePosition[0] &&
					i*machine.ButtonAMovements[1]+j*machine.ButtonBMovements[1] == machine.PrizePosition[1] &&
					3*i+j < minTokens {
					minTokens = 3*i + j
				}
			}
		}
		if minTokens < 401 {
			total += minTokens
		}
	}
	return total
}

// partTwo solves part two of the puzzle.
func partTwo(machines []Machine) int {
	for i, machine := range machines {
		machines[i].PrizePosition = [2]int{machine.PrizePosition[0] + 10000000000000, machine.PrizePosition[1] + 10000000000000}
	}
	total := 0
	for _, machine := range machines {
		numerator := machine.PrizePosition[0]*machine.ButtonAMovements[1] - machine.PrizePosition[1]*machine.ButtonAMovements[0]
		denominator := machine.ButtonBMovements[0]*machine.ButtonAMovements[1] - machine.ButtonAMovements[0]*machine.ButtonBMovements[1]
		if numerator%denominator == 0 {
			b := numerator / denominator
			numerator := machine.PrizePosition[0] - b*machine.ButtonBMovements[0]
			if numerator%machine.ButtonAMovements[0] == 0 {
				a := numerator / machine.ButtonAMovements[0]
				if a >= 0 && b >= 0 {
					total += 3*a + b
				}
			}
		}
	}
	return total
}

type Machine struct {
	ButtonAMovements [2]int
	ButtonBMovements [2]int
	PrizePosition    [2]int
}

// readLines converts the information from the file into a usable form.
func readLines(fileName string) ([]Machine, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)

	machines := []Machine{}
	newMachine := Machine{}
	for scanner.Scan() {
		if scanner.Text() == "" {
			machines = append(machines, newMachine)
			newMachine = Machine{}
		} else {
			components := strings.Split(scanner.Text(), ": ")
			xAndY := strings.Split(components[1], ", ")
			x, err := strconv.Atoi(xAndY[0][2:])
			if err != nil {
				return nil, err
			}

			y, err := strconv.Atoi(xAndY[1][2:])
			if err != nil {
				return nil, err
			}

			if components[0] == "Button A" {
				newMachine.ButtonAMovements = [2]int{x, y}
			} else if components[0] == "Button B" {
				newMachine.ButtonBMovements = [2]int{x, y}
			} else {
				newMachine.PrizePosition = [2]int{x, y}
			}
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	machines = append(machines, newMachine)

	return machines, nil
}
