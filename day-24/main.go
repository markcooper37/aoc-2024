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
	startWires, gates, err := readLines("fixed.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(partOne(startWires, gates))
	fmt.Println(partTwo(startWires, gates))
}

// partOne solves part one of the puzzle.
func partOne(startWires map[string]int, gates []Gate) int {
	allWires := allWires(startWires, gates)
	completedWires := map[string]int{}
	for wire, value := range startWires {
		completedWires[wire] = value
	}
	for len(completedWires) < len(allWires) {
		for _, gate := range gates {
			if input1, input1Exists := completedWires[gate.Inputs[0]]; input1Exists {
				if input2, input2Exists := completedWires[gate.Inputs[1]]; input2Exists {
					completedWires[gate.Output] = calculateOutput([2]int{input1, input2}, gate.Operation)
				}
			}
		}
	}
	zWires := zWires(allWires)
	output := 0
	for index, wire := range zWires {
		output = output | (completedWires[wire] << index)
	}

	return output
}

// partTwo solves part two of the puzzle.
func partTwo(startWires map[string]int, gates []Gate) int {
	// z XOR (a XOR b) -> ((z AND (a XOR b)) OR (a AND b)) XOR (c XOR d)
	// z28 <> tfb, z08 <> vvr, mqh <> z39, bkr<>rnq
	stuff := findStuff(startWires, gates)
	for index, row := range stuff {
		fmt.Println("index", index)
		fmt.Println(len(row))
		fmt.Println(row)
	}
	return 0
}

type Gate struct {
	Inputs    [2]string
	Operation string
	Output    string
}

func allWires(startWires map[string]int, gates []Gate) map[string]bool {
	allWires := map[string]bool{}
	for wire := range startWires {
		allWires[wire] = true
	}
	for _, gate := range gates {
		allWires[gate.Inputs[0]] = true
		allWires[gate.Inputs[1]] = true
		allWires[gate.Output] = true
	}
	return allWires
}

func zWires(allWires map[string]bool) []string {
	zWires := []string{}
	for wire := range allWires {
		if wire[0] == 'z' {
			zWires = append(zWires, wire)
		}
	}
	slices.Sort(zWires)
	return zWires
}

func calculateOutput(inputs [2]int, operation string) int {
	switch operation {
	case "AND":
		return inputs[0] & inputs[1]
	case "OR":
		return inputs[0] | inputs[1]
	case "XOR":
		return inputs[0] ^ inputs[1]
	}
	return 0
}

func findStuff(startWires map[string]int, gates []Gate) [][]string {
	allWires := allWires(startWires, gates)
	reverseGates := map[string][]string{}
	for _, gate := range gates {
		reverseGates[gate.Output] = []string{gate.Inputs[0], gate.Operation, gate.Inputs[1]}
	}
	zWires := zWires(allWires)
	stuff := [][]string{}
	for _, zWire := range zWires {
		start := reverseGates[zWire]
		found := true
		for found {
			found = false
			for i, value := range start {
				if gate, ok := reverseGates[value]; ok {
					startFirst := make([]string, i)
					copy(startFirst, start[:i])
					startEnd := make([]string, len(start)-i-1)
					copy(startEnd, start[i+1:])
					new := append(startFirst, gate...)
					new = append(new, startEnd...)
					start = new
					found = true
					break
				}
			}
		}
		stuff = append(stuff, start)
	}
	return stuff
}

// readLines converts the information from the file into a usable form.
func readLines(fileName string) (map[string]int, []Gate, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, nil, err
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)

	startWires := map[string]int{}
	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}
		parts := strings.Split(scanner.Text(), ": ")
		value, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, nil, err
		}

		startWires[parts[0]] = value
	}
	gates := []Gate{}
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " ")
		gates = append(gates, Gate{
			Inputs:    [2]string{parts[0], parts[2]},
			Operation: parts[1],
			Output:    parts[4],
		})
	}
	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}

	return startWires, gates, nil
}
