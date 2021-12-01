package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func load_file(filename string) []int {

	file, err := os.Open(filename)
	line_list := make([]int, 0)

	if err == nil {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			depth_int, _ := strconv.Atoi(scanner.Text())
			line_list = append(line_list, depth_int)
		}
	} else {
		fmt.Println("Error opening file")
		fmt.Println(err)
	}

	return line_list
}

func part_one(depth_data []int) {

	increase_count := 0

	for index, depth_value := range depth_data {
		if index > 0 {
			if depth_value > depth_data[index-1] {
				increase_count += 1
			}
		}
	}

	fmt.Println("Pt 1 - Number of increases: ", increase_count)
}

func part_two(depth_data []int) {

	increase_count := 0

	for index, _ := range depth_data {
		if index > 2 {
			if calculate_window(depth_data, index) > calculate_window(depth_data, index-1) {
				increase_count += 1
			}
		}
	}

	fmt.Println("Pt 2 - Number of increases: ", increase_count)
}

func calculate_window(depth_data []int, index int) int {
	return depth_data[index] + depth_data[index-1] + depth_data[index-2]
}

func main() {
	depth_data := load_file("day_01/day_01.txt")
	part_one(depth_data)
	part_two(depth_data)
}
