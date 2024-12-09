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
	diskMap, err := readLines("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(partOne(diskMap))
	fmt.Println(partTwo(diskMap))
}

// partOne solves part one of the puzzle.
func partOne(diskMap []int) int {
	blocks := []int{}
	for i, value := range diskMap {
		blocks = append(blocks, getBlock(i, value)...)
	}
	firstFree := 0
	for index, value := range blocks {
		if value == -1 {
			firstFree = index
			break
		}
	}
	lastTaken := len(blocks) - 1
	for firstFree < lastTaken {
		blocks[firstFree] = blocks[lastTaken]
		blocks[lastTaken] = -1
		for i := firstFree + 1; i < len(blocks); i++ {
			if blocks[i] == -1 {
				firstFree = i
				break
			}
		}
		for i := lastTaken - 1; i >= 0; i-- {
			if blocks[i] != -1 {
				lastTaken = i
				break
			}
		}
	}
	return calculateChecksum(blocks)
}

type Block struct {
	Index  int
	Value  int
	Length int
}

// partTwo solves part two of the puzzle.
func partTwo(diskMap []int) int {
	blocks := []int{}
	for i, value := range diskMap {
		blocks = append(blocks, getBlock(i, value)...)
	}
	index := 0
	freeMap := map[int][]int{}
	fileBlocks := []Block{}
	for i, value := range diskMap {
		if i%2 == 0 {
			fileBlocks = append(fileBlocks, Block{Index: index, Length: value, Value: i / 2})
		} else {
			freeMap[value] = append(freeMap[value], index)
		}
		index += value
	}
	slices.Reverse(fileBlocks)
	for _, block := range fileBlocks {
		minIndex := block.Index
		gapSize := -1
		for i := block.Length; i <= 9; i++ {
			if len(freeMap[i]) != 0 && freeMap[i][0] < minIndex {
				minIndex = freeMap[i][0]
				gapSize = i
			}
		}
		if minIndex < block.Index {
			for i := 0; i < block.Length; i++ {
				blocks[minIndex+i] = block.Value
				blocks[block.Index+i] = -1
			}
			freeMap[gapSize] = freeMap[gapSize][1:]
			freeMap[gapSize-block.Length] = append(freeMap[gapSize-block.Length], minIndex+block.Length)
			slices.Sort(freeMap[gapSize-block.Length])

		}
	}
	return calculateChecksum(blocks)
}

// getBlock returns a block based on the given index and value.
func getBlock(index, value int) []int {
	block := make([]int, value)
	if index%2 != 0 {
		for i := 0; i < value; i++ {
			block[i] = -1
		}
	} else {
		number := index / 2
		for i := 0; i < value; i++ {
			block[i] = number
		}
	}
	return block
}

// calculateChecksum calculates the checksum for a set of blocks.
func calculateChecksum(blocks []int) int {
	total := 0
	for i, value := range blocks {
		if value == -1 {
			continue
		}
		total += i * value
	}
	return total
}

// readLines converts the information from the file into a usable form.
func readLines(fileName string) ([]int, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)

	diskMap := []int{}
	for scanner.Scan() {
		values := strings.Split(scanner.Text(), "")
		for _, valueStr := range values {
			value, err := strconv.Atoi(valueStr)
			if err != nil {
				return nil, err
			}
			diskMap = append(diskMap, value)
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return diskMap, nil
}
