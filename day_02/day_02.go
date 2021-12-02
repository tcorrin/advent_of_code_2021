package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type instruction struct {
	direction string
	magnitude int
}

func load_file(filename string) []instruction {

	file, err := os.Open(filename)
	instructions := make([]instruction, 0)

	if err == nil {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line_slice := strings.Split(scanner.Text(), " ")
			magnitude, _ := strconv.Atoi(line_slice[1])
			instructions = append(instructions, instruction{direction: line_slice[0], magnitude: magnitude})
		}
	} else {
		fmt.Println("Error opening file")
		fmt.Println(err)
	}

	return instructions
}

func process_file(filename string) {

	instructions := load_file(filename)
	x := 0
	y := 0

	for _, instruction := range instructions {
		if instruction.direction == "forward" {
			x += instruction.magnitude
		} else if instruction.direction == "up" {
			y -= instruction.magnitude
		} else if instruction.direction == "down" {
			y += instruction.magnitude
		}
	}

	fmt.Println("Part 1 result for ", filename, " x: ", x, " y: ", y, " x*y: ", x*y)

	x = 0
	y = 0
	aim := 0

	for _, instruction := range instructions {
		if instruction.direction == "forward" {
			x += instruction.magnitude
			y += instruction.magnitude * aim
		} else if instruction.direction == "up" {
			aim -= instruction.magnitude
		} else if instruction.direction == "down" {
			aim += instruction.magnitude
		}
	}

	fmt.Println("Part 2 result for ", filename, " x: ", x, " y: ", y, " x*y: ", x*y)
}

func main() {
	process_file("day_02_test.txt")
	process_file("day_02.txt")
}
