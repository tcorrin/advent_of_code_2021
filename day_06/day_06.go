package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func load_file(filename string) map[int]int {
	number_map := make(map[int]int)
	for i := 0; i < 9; i++ {
		number_map[i] = 0
	}

	file, err := os.Open(filename)
	if err == nil {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			split_list := strings.Split(scanner.Text(), ",")
			for _, number := range split_list {
				number_int, _ := strconv.Atoi(number)
				number_map[number_int]++
			}
		}
	} else {
		fmt.Println("Error opening file")
		fmt.Println(err)
	}

	return number_map
}

func process_day(number_map map[int]int) map[int]int {
	new_number_map := make(map[int]int)
	for i := 9; i > -1; i-- {
		if i == 0 {
			new_number_map[8] = number_map[i]
			new_number_map[6] += number_map[i]
		} else {
			new_number_map[i-1] = number_map[i]
		}
	}
	return new_number_map
}

func process_file(filename string, number_of_days int) {
	number_map := load_file(filename)
	for i := 0; i < number_of_days; i++ {
		number_map = process_day(number_map)
	}

	result := 0

	for i := 0; i < 9; i++ {
		result += number_map[i]
	}

	fmt.Println("Final fish count: ", result)
}

func main() {
	process_file("day_06_test.txt", 18)
	process_file("day_06_test.txt", 80)
	process_file("day_06_test.txt", 256)
	process_file("day_06.txt", 80)
	process_file("day_06.txt", 256)
}
