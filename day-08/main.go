package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	antennaMap, err := readLines("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(partOne(antennaMap))
	fmt.Println(partTwo(antennaMap))
}

// partOne solves part one of the puzzle.
func partOne(antennaMap [][]string) int {
	antennaLocations := map[string][][2]int{}
	for i, row := range antennaMap {
		for j, antenna := range row {
			if antenna != "." {
				antennaLocations[antenna] = append(antennaLocations[antenna], [2]int{i, j})
			}
		}
	}
	antinodes := map[[2]int]bool{}
	for _, locations := range antennaLocations {
		for i := 0; i < len(locations)-1; i++ {
			for j := i + 1; j < len(locations); j++ {
				diff := [2]int{locations[i][0] - locations[j][0], locations[i][1] - locations[j][1]}
				first := [2]int{locations[i][0] + diff[0], locations[i][1] + diff[1]}
				second := [2]int{locations[j][0] - diff[0], locations[j][1] - diff[1]}
				if first[0] >= 0 && first[0] < len(antennaMap) && first[1] >= 0 && first[1] < len(antennaMap[0]) {
					antinodes[first] = true
				}
				if second[0] >= 0 && second[0] < len(antennaMap) && second[1] >= 0 && second[1] < len(antennaMap[0]) {
					antinodes[second] = true
				}
			}
		}
	}
	return len(antinodes)
}

// partTwo solves part two of the puzzle.
func partTwo(antennaMap [][]string) int {
	antennaLocations := map[string][][2]int{}
	for i, row := range antennaMap {
		for j, antenna := range row {
			if antenna != "." {
				antennaLocations[antenna] = append(antennaLocations[antenna], [2]int{i, j})
			}
		}
	}
	antinodes := map[[2]int]bool{}
	for _, locations := range antennaLocations {
		for i := 0; i < len(locations)-1; i++ {
			for j := i + 1; j < len(locations); j++ {
				diff := [2]int{locations[i][0] - locations[j][0], locations[i][1] - locations[j][1]}
				first := [2]int{locations[i][0], locations[i][1]}
				second := [2]int{locations[j][0], locations[j][1]}
				for {
					if first[0] >= 0 && first[0] < len(antennaMap) && first[1] >= 0 && first[1] < len(antennaMap[0]) {
						antinodes[first] = true
					} else {
						break
					}
					first = [2]int{first[0] + diff[0], first[1] + diff[1]}
				}
				for {
					if second[0] >= 0 && second[0] < len(antennaMap) && second[1] >= 0 && second[1] < len(antennaMap[0]) {
						antinodes[second] = true
					} else {
						break
					}
					second = [2]int{second[0] - diff[0], second[1] - diff[1]}
				}
			}
		}
	}
	return len(antinodes)
}

// readLines converts the information from the file into a usable form.
func readLines(fileName string) ([][]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)

	antennaMap := [][]string{}
	for scanner.Scan() {
		antennaMap = append(antennaMap, strings.Split(scanner.Text(), ""))
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return antennaMap, nil
}
