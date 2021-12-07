package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const MaxInt = int(^uint(0) >> 1)

func load_file(filename string) ([]int, int, int) {
	number_list := make([]int, 0)
	max := 0
	min := MaxInt

	file, err := os.Open(filename)
	if err == nil {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			split_list := strings.Split(scanner.Text(), ",")
			for _, number := range split_list {
				number_int, _ := strconv.Atoi(number)
				if number_int > max {
					max = number_int
				} else if number_int < min {
					min = number_int
				}
				number_list = append(number_list, number_int)
			}
		}
	} else {
		fmt.Println("Error opening file")
		fmt.Println(err)
	}

	return number_list, max, min
}

func calc_fuel_consumption(number_list []int, i int) int {
	result := 0
	for _, number := range number_list {
		if number > i {
			result += number - i
		} else if number < i {
			result += i - number
		}
	}
	return result
}

func incremental_fuel_burn(distance int) int {
	result := 0
	for i := 1; i <= distance; i++ {
		result += i
	}
	return result
}

func calc_modified_fuel_consumption(number_list []int, i int) int {
	result := 0
	for _, number := range number_list {
		if number > i {
			result += incremental_fuel_burn(number - i)
		} else if number < i {
			result += incremental_fuel_burn(i - number)
		}
	}
	return result
}

func process_file(filename string, modified_fuel_consumption bool) {
	number_list, max, min := load_file(filename)
	min_fuel_consumption := MaxInt
	fuel_consumption := 0

	for i := min; i <= max; i++ {
		if modified_fuel_consumption {
			fuel_consumption = calc_modified_fuel_consumption(number_list, i)
		} else {
			fuel_consumption = calc_fuel_consumption(number_list, i)
		}
		if fuel_consumption < min_fuel_consumption {
			min_fuel_consumption = fuel_consumption
		}
	}

	fmt.Println("Fuel Consumption: ", min_fuel_consumption)
}

func main() {
	process_file("day_07_test.txt", false)
	process_file("day_07.txt", false)
	process_file("day_07_test.txt", true)
	process_file("day_07.txt", true)
}
