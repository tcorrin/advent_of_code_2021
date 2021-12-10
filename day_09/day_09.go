package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type position struct {
	x int
	y int
}

func contains_position(list []position, position_to_find position) bool {
	for _, value := range list {
		if value == position_to_find {
			return true
		}
	}
	return false
}

func remove_position(list []position, position_to_remove position) []position {
	new_list := make([]position, 0)
	for _, value := range list {
		if value != position_to_remove {
			new_list = append(new_list, value)
		}
	}
	return new_list
}

func load_file(filename string) [][]int {
	floor_map := make([][]int, 0)

	file, err := os.Open(filename)
	if err == nil {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			row := make([]int, 0)
			split_list := strings.Split(scanner.Text(), "")
			for _, value := range split_list {
				value_int, _ := strconv.Atoi(value)
				row = append(row, value_int)
			}
			floor_map = append(floor_map, row)
		}
	} else {
		fmt.Println("Error opening file")
		fmt.Println(err)
		os.Exit(23)
	}

	return floor_map
}

func is_low_spot(x int, y int, floor_map [][]int, height int, width int) bool {

	if x+1 < width && floor_map[y][x+1] <= floor_map[y][x] {
		return false
	}

	if x-1 > -1 && floor_map[y][x-1] <= floor_map[y][x] {
		return false
	}

	if y+1 < height && floor_map[y+1][x] <= floor_map[y][x] {
		return false
	}

	if y-1 > -1 && floor_map[y-1][x] <= floor_map[y][x] {
		return false
	}

	return true

}

func calc_basin_size(x int, y int, floor_map [][]int, height int, width int) int {
	list_of_basin_positions := make([]position, 0)
	list_of_basin_positions = append(list_of_basin_positions, position{x: x, y: y})
	list_of_processed_positions := make([]position, 0)

	for len(list_of_basin_positions) > 0 {
		position_to_process := list_of_basin_positions[0]
		x := position_to_process.x
		y := position_to_process.y

		if x+1 < width && floor_map[y][x+1] < 9 {
			new_position := position{x + 1, y}
			if !contains_position(list_of_basin_positions, new_position) && !contains_position(list_of_processed_positions, new_position) {
				list_of_basin_positions = append(list_of_basin_positions, new_position)
			}
		}

		if x-1 > -1 && floor_map[y][x-1] < 9 {
			new_position := position{x - 1, y}
			if !contains_position(list_of_basin_positions, new_position) && !contains_position(list_of_processed_positions, new_position) {
				list_of_basin_positions = append(list_of_basin_positions, new_position)
			}
		}

		if y+1 < height && floor_map[y+1][x] < 9 {
			new_position := position{x, y + 1}
			if !contains_position(list_of_basin_positions, new_position) && !contains_position(list_of_processed_positions, new_position) {
				list_of_basin_positions = append(list_of_basin_positions, new_position)
			}
		}

		if y-1 > -1 && floor_map[y-1][x] < 9 {
			new_position := position{x, y - 1}
			if !contains_position(list_of_basin_positions, new_position) && !contains_position(list_of_processed_positions, new_position) {
				list_of_basin_positions = append(list_of_basin_positions, new_position)
			}
		}

		list_of_basin_positions = remove_position(list_of_basin_positions, position_to_process)
		list_of_processed_positions = append(list_of_processed_positions, position_to_process)
	}

	return len(list_of_processed_positions)
}

func process_file(filename string) {
	floor_map := load_file(filename)
	height := len(floor_map)
	width := len(floor_map[0])
	risk_level := 0
	basin_size_list := make([]int, 0)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if is_low_spot(x, y, floor_map, height, width) {
				risk_level += floor_map[y][x] + 1
				basin_size_list = append(basin_size_list, calc_basin_size(x, y, floor_map, height, width))
			}
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(basin_size_list)))

	fmt.Println("Risk Level: ", risk_level)
	fmt.Println("Basin size result: ", basin_size_list[0]*basin_size_list[1]*basin_size_list[2])
}

func main() {
	process_file("day_09_test.txt")
	process_file("day_09.txt")
}
