package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

func main() {
	connections, err := readLines("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(partOne(connections))
	fmt.Println(partTwo(connections))
}

// partOne solves part one of the puzzle.
func partOne(connections [][2]string) int {
	sets := interconnectedTrios(connections)
	total := 0
	for set := range sets {
		if set[0][0] == 't' || set[1][0] == 't' || set[2][0] == 't' {
			total++
		}
	}
	return total
}

// partTwo solves part two of the puzzle.
func partTwo(connections [][2]string) string {
	connectionMap := constructConnectionMap(connections)
	computerMap := constructComputerMap(connections)
	cliques := maximalCliques([]string{}, computerMap, map[string]bool{}, connectionMap)
	longestClique := longest(cliques)
	slices.Sort(longestClique)
	return strings.Join(longestClique, ",")
}

// interconnectedTrios creates a list of all interconnected trios.
func interconnectedTrios(connections [][2]string) map[[3]string]bool {
	connectionMap := constructConnectionMap(connections)
	validTrios := map[[3]string]bool{}
	for _, connection := range connections {
		for computer := range connectionMap[connection[0]] {
			if connectionMap[connection[1]][computer] {
				trio := []string{connection[0], connection[1], computer}
				slices.Sort(trio)
				validTrios[[3]string{trio[0], trio[1], trio[2]}] = true
			}
		}
	}
	return validTrios
}

// constructConnectionMap constructs a map of all connections.
func constructConnectionMap(connections [][2]string) map[string]map[string]bool {
	connectionMap := map[string]map[string]bool{}
	for _, connection := range connections {
		if _, ok := connectionMap[connection[0]]; !ok {
			connectionMap[connection[0]] = map[string]bool{}
		}
		connectionMap[connection[0]][connection[1]] = true
		if _, ok := connectionMap[connection[1]]; !ok {
			connectionMap[connection[1]] = map[string]bool{}
		}
		connectionMap[connection[1]][connection[0]] = true
	}
	return connectionMap
}

// constructComputerMap constructs a map of all computers.
func constructComputerMap(connections [][2]string) map[string]bool {
	computerMap := map[string]bool{}
	for _, connection := range connections {
		computerMap[connection[0]] = true
		computerMap[connection[1]] = true
	}
	return computerMap
}

// maximalCliques finds all maximal cliques.
func maximalCliques(clique []string, options, excluded map[string]bool, connectionMap map[string]map[string]bool) [][]string {
	if len(options) == 0 && len(excluded) == 0 {
		return [][]string{clique}
	}
	maxCliques := [][]string{}
	for vertex := range options {
		newClique := make([]string, len(clique))
		copy(newClique, clique)
		newClique = append(newClique, vertex)

		newOptions := map[string]bool{}
		for k := range options {
			if connectionMap[vertex][k] {
				newOptions[k] = true
			}
		}

		newExcluded := map[string]bool{}
		for k := range excluded {
			if connectionMap[vertex][k] {
				newExcluded[vertex] = true
			}
		}
		
		maxCliques = append(maxCliques, maximalCliques(newClique, newOptions, newExcluded, connectionMap)...)
		delete(options, vertex)
		excluded[vertex] = true
	}
	return maxCliques
}

// longest returns the longest slice.
func longest(slices [][]string) []string {
	longest := []string{}
	for _, slice := range slices {
		if len(slice) > len(longest) {
			longest = slice
		}
	}
	return longest
}

// readLines converts the information from the file into a usable form.
func readLines(fileName string) ([][2]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)

	connections := [][2]string{}
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), "-")
		connections = append(connections, [2]string{parts[0], parts[1]})
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return connections, nil
}
