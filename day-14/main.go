package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	robots, err := readLines("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(partOne(robots, [2]int{101, 103}))
	fmt.Println(partTwo(robots, [2]int{101, 103}))
}

// partOne solves part one of the puzzle.
func partOne(robots []Robot, dimensions [2]int) int {
	newRobots := []Robot{}
	for _, robot := range robots {
		newRobots = append(newRobots, Robot{
			Position: [2]int{((robot.Position[0]+100*robot.Velocity[0])%dimensions[0] + dimensions[0]) % dimensions[0],
				((robot.Position[1]+100*robot.Velocity[1])%dimensions[1] + dimensions[1]) % dimensions[1]},
			Velocity: robot.Velocity,
		})
	}
	quadrants := map[int]int{}
	for _, robot := range newRobots {
		if robot.Position[0] < (dimensions[0]-1)/2 && robot.Position[1] < (dimensions[1]-1)/2 {
			quadrants[1]++
		} else if robot.Position[0] < (dimensions[0]-1)/2 && robot.Position[1] > (dimensions[1]-1)/2 {
			quadrants[2]++
		} else if robot.Position[0] > (dimensions[0]-1)/2 && robot.Position[1] < (dimensions[1]-1)/2 {
			quadrants[3]++
		} else if robot.Position[0] > (dimensions[0]-1)/2 && robot.Position[1] > (dimensions[1]-1)/2 {
			quadrants[4]++
		}
	}
	return quadrants[1] * quadrants[2] * quadrants[3] * quadrants[4]
}

// partTwo solves part two of the puzzle.
func partTwo(robots []Robot, dimensions [2]int) error {
	file, err := os.Create("pictures.csv")
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()
	for i := 1; i <= dimensions[0] * dimensions[1]; i++ {
		newRobots := []Robot{}
		for _, robot := range robots {
			newRobots = append(newRobots, Robot{
				Position: [2]int{((robot.Position[0]+robot.Velocity[0])%dimensions[0] + dimensions[0]) % dimensions[0],
					((robot.Position[1]+robot.Velocity[1])%dimensions[1] + dimensions[1]) % dimensions[1]},
				Velocity: robot.Velocity,
			})
		}
		picture := constructPicture(newRobots, [2]int{dimensions[0], dimensions[1]})
		err := writer.Write([]string{strconv.Itoa(i)})
		if err != nil {
			return err
		}
		for _, row := range picture {
			err := writer.Write([]string{row})
			if err != nil {
				return err
			}
		}
		robots = newRobots
	}
	return nil
}

type Robot struct {
	Position [2]int
	Velocity [2]int
}

// constructPicture returns a picture of the positions of all robots.
func constructPicture(robots []Robot, dimensions [2]int) []string {
	positionMap := map[[2]int]bool{}
	for _, robot := range robots {
		positionMap[robot.Position] = true
	}
	picture := []string{}
	for i := 0; i < dimensions[1]; i++ {
		var sb strings.Builder
		for j := 0; j < dimensions[0]; j++ {
			if positionMap[[2]int{j, i}] {
				sb.WriteString("*")
			} else {
				sb.WriteString(".")
			}
		}
		picture = append(picture, sb.String())
	}
	return picture
}

// readLines converts the information from the file into a usable form.
func readLines(fileName string) ([]Robot, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)

	robots := []Robot{}
	for scanner.Scan() {
		components := strings.Split(scanner.Text(), " ")
		positionComponents := strings.Split(strings.TrimLeft(components[0], "p="), ",")
		positionX, err := strconv.Atoi(positionComponents[0])
		if err != nil {
			return nil, err
		}

		positionY, err := strconv.Atoi(positionComponents[1])
		if err != nil {
			return nil, err
		}

		velocityComponents := strings.Split(strings.TrimLeft(components[1], "v="), ",")
		velocityX, err := strconv.Atoi(velocityComponents[0])
		if err != nil {
			return nil, err
		}

		velocityY, err := strconv.Atoi(velocityComponents[1])
		if err != nil {
			return nil, err
		}
		robots = append(robots, Robot{Position: [2]int{positionX, positionY}, Velocity: [2]int{velocityX, velocityY}})
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return robots, nil
}
