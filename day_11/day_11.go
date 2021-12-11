package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type position struct {
	x int
	y int
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

func contains_position(list []position, position_to_find position) bool {
	for _, value := range list {
		if value == position_to_find {
			return true
		}
	}
	return false
}

func load_file(filename string) [][]int {
	octo_grid := make([][]int, 0)

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
			octo_grid = append(octo_grid, row)
		}
	} else {
		fmt.Println("Error opening file")
		fmt.Println(err)
		os.Exit(23)
	}

	return octo_grid
}

func increment_octo(x int, y int, octo_grid_pointer *[][]int, flash_list_pointer *[]position, processed_flashes []position) {
	octo_grid := *octo_grid_pointer
	value := octo_grid[y][x]
	if value < 9 {
		octo_grid[y][x]++
	}

	if value == 9 &&
		!contains_position(*flash_list_pointer, position{x: x, y: y}) &&
		!contains_position(processed_flashes, position{x: x, y: y}) {
		*flash_list_pointer = append(*flash_list_pointer, position{x: x, y: y})
	}
}

func process_flash(flash_list_pointer *[]position, octo_grid_pointer *[][]int, processed_flashes []position) {
	flash_list := *flash_list_pointer
	flash_position := flash_list[0]
	// increment flash position neighbours
	x := flash_position.x
	y := flash_position.y
	if x+1 < 10 {
		increment_octo(x+1, y, octo_grid_pointer, flash_list_pointer, processed_flashes)
		if y+1 < 10 {
			increment_octo(x+1, y+1, octo_grid_pointer, flash_list_pointer, processed_flashes)
		}
		if y-1 > -1 {
			increment_octo(x+1, y-1, octo_grid_pointer, flash_list_pointer, processed_flashes)
		}
	}
	if x-1 > -1 {
		increment_octo(x-1, y, octo_grid_pointer, flash_list_pointer, processed_flashes)
		if y+1 < 10 {
			increment_octo(x-1, y+1, octo_grid_pointer, flash_list_pointer, processed_flashes)
		}
		if y-1 > -1 {
			increment_octo(x-1, y-1, octo_grid_pointer, flash_list_pointer, processed_flashes)
		}
	}
	if y-1 > -1 {
		increment_octo(x, y-1, octo_grid_pointer, flash_list_pointer, processed_flashes)
	}
	if y+1 < 10 {
		increment_octo(x, y+1, octo_grid_pointer, flash_list_pointer, processed_flashes)
	}

	*flash_list_pointer = remove_position(*flash_list_pointer, flash_position)
}

func run_step(octo_grid_pointer *[][]int) int {
	flash_list := make([]position, 0)
	processed_flashes := make([]position, 0)

	//inital increment
	for y := 0; y < 10; y++ {
		for x := 0; x < 10; x++ {
			increment_octo(x, y, octo_grid_pointer, &flash_list, processed_flashes)
		}
	}

	//process flash list
	for len(flash_list) > 0 {
		processed_flashes = append(processed_flashes, flash_list[0])
		process_flash(&flash_list, octo_grid_pointer, processed_flashes)
	}

	octo_grid := *octo_grid_pointer
	//set 9s to 0s
	for _, flash := range processed_flashes {
		octo_grid[flash.y][flash.x] = 0
	}

	return len(processed_flashes)
}

func process_file(filename string, steps int) {
	octo_grid := load_file(filename)
	flash_count := 0

	for i := 0; i < steps; i++ {
		flash_count += run_step(&octo_grid)
	}

	fmt.Println("Number of flashes: ", flash_count)

	octo_grid = load_file(filename)
	step_flash_count := 0
	i := 0
	for step_flash_count != 100 {
		step_flash_count = run_step(&octo_grid)
		i++
	}
	fmt.Println("Number of steps till big flash: ", i)

}

func main() {
	process_file("day_11_test.txt", 100)
	process_file("day_11.txt", 100)
}
