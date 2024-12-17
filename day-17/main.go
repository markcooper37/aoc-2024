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
	computer, err := readLines("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(partOne(computer))
	fmt.Println(partTwo(computer))
}

type Computer struct {
	Registers          [3]int
	Program            []int
	InstructionPointer int
	Outputs            []int
}

// partOne solves part one of the puzzle.
func partOne(computer Computer) string {
	c := computer.copy()
	for c.InstructionPointer < len(c.Program) {
		c.performOperation()
	}

	return constructOutputString(c.Outputs)
}

// partTwo solves part two of the puzzle.
func partTwo(computer Computer) int {
	computers := []Computer{}
	for i := 0; i < 8; i++ {
		computers = append(computers, findValidStarts([3]int{0, 0, i}, computer.Program)...)
	}

	min := -1
	for _, comp := range computers {
		if min == -1 || comp.Registers[0] < min {
			min = comp.Registers[0]
		}
	}
	return min
}

// copy creates a copy of the computer.
func (c *Computer) copy() Computer {
	program := make([]int, len(c.Program))
	copy(program, c.Program)
	return Computer{
		Program:   program,
		Registers: c.Registers,
	}
}

// performOperation performs a computer operation.
func (c *Computer) performOperation() {
	switch c.Program[c.InstructionPointer] {
	case 0:
		c.adv()
	case 1:
		c.bxl()
	case 2:
		c.bst()
	case 3:
		c.jnz()
	case 4:
		c.bxc()
	case 5:
		c.out()
	case 6:
		c.bdv()
	case 7:
		c.cdv()
	}
}

// adv performs the adv operation.
func (c *Computer) adv() {
	numerator := c.Registers[0]
	operand := c.comboOperand()
	denominator := 1 << operand
	c.Registers[0] = numerator / denominator
	c.InstructionPointer += 2
}

// bxl performs the bxl operation.
func (c *Computer) bxl() {
	c.Registers[1] = c.Registers[1] ^ c.Program[c.InstructionPointer+1]
	c.InstructionPointer += 2
}

// bst performs the bst operation.
func (c *Computer) bst() {
	c.Registers[1] = c.comboOperand() % 8
	c.InstructionPointer += 2
}

// jnz performs the jnz operation.
func (c *Computer) jnz() {
	if c.Registers[0] != 0 {
		c.InstructionPointer = c.Program[c.InstructionPointer+1]
	} else {
		c.InstructionPointer += 2
	}
}

// bxc performs the bxc operation.
func (c *Computer) bxc() {
	c.Registers[1] = c.Registers[1] ^ c.Registers[2]
	c.InstructionPointer += 2
}

// out performs the out operation.
func (c *Computer) out() {
	operand := c.comboOperand() % 8
	c.Outputs = append(c.Outputs, operand)
	c.InstructionPointer += 2
}

// bdv performs the bdv operation.
func (c *Computer) bdv() {
	numerator := c.Registers[0]
	operand := c.comboOperand()
	denominator := 1 << operand
	c.Registers[1] = numerator / denominator
	c.InstructionPointer += 2
}

// cdv performs the cdv operation.
func (c *Computer) cdv() {
	numerator := c.Registers[0]
	operand := c.comboOperand()
	denominator := 1 << operand
	c.Registers[2] = numerator / denominator
	c.InstructionPointer += 2
}

// comboOperand calculates the combo operand.
func (c *Computer) comboOperand() int {
	operand := c.Program[c.InstructionPointer+1]
	if operand == 4 {
		operand = c.Registers[0]
	} else if operand == 5 {
		operand = c.Registers[1]
	} else if operand == 6 {
		operand = c.Registers[2]
	}
	return operand
}

// constructOutputString constructs a comma-separated string from a computer output.
func constructOutputString(program []int) string {
	outputStr := []string{}
	for _, value := range program {
		outputStr = append(outputStr, strconv.Itoa(value))
	}
	return strings.Join(outputStr, ",")
}

// findValidStarts finds all computers that produce the given end registers and output program.
// It assumes that the last step of a program jumps back to the first step, and in each run through,
// all other registers are derived from A and A is divided by 8.
func findValidStarts(endRegisters [3]int, program []int) []Computer {
	computers := []Computer{{Registers: endRegisters, Program: program}}
	for i := len(program) - 1; i >= 0; i-- {
		newComputers := []Computer{}
		for _, computer := range computers {
			for offset := 0; offset < 8; offset++ {
				testComputer := Computer{Registers: [3]int{computer.Registers[0]*8 + offset, 0, 0}, Program: program}
				for step := 0; step < len(testComputer.Program)/2; step++ {
					testComputer.performOperation()
				}
				if testComputer.Outputs[0] == program[i] {
					newComputers = append(newComputers, Computer{Registers: [3]int{computer.Registers[0]*8 + offset, 0, 0}, Program: program})
				}
			}
		}
		computers = newComputers
	}
	return computers
}

// readLines converts the information from the file into a usable form.
func readLines(fileName string) (Computer, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return Computer{}, err
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)

	computer := Computer{Registers: [3]int{}, Program: []int{}}
	index := 0
	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}
		splitValues := strings.Split(scanner.Text(), ": ")
		value, err := strconv.Atoi(splitValues[1])
		if err != nil {
			return Computer{}, err
		}

		computer.Registers[index] = value
		index++
	}
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		return Computer{}, err
	}

	splitValues := strings.Split(scanner.Text(), ": ")
	programValues := strings.Split(splitValues[1], ",")
	for _, stringValue := range programValues {
		value, err := strconv.Atoi(stringValue)
		if err != nil {
			return Computer{}, nil
		}
		computer.Program = append(computer.Program, value)
	}

	return computer, nil
}
