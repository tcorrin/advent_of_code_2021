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

func count_number_of_increases(depth_data []int, window_size int) {

	increase_count := 0

	for index, _ := range depth_data {
		if index > window_size-1 {
			if calculate_window(depth_data, index, window_size) > calculate_window(depth_data, index-1, window_size) {
				increase_count += 1
			}
		}
	}

	fmt.Println("Number of increases: ", increase_count)
}

func calculate_window(depth_data []int, index int, window_size int) int {
	result := 0
	for x := 0; x < window_size; x++ {
		result += depth_data[index-x]
	}
	return result
}

func main() {
	depth_data := load_file("day_01/day_01.txt")
	count_number_of_increases(depth_data, 1)
	count_number_of_increases(depth_data, 3)
}
