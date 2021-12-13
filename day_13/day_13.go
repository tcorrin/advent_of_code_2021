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

type fold struct {
	axis     string
	location int
}

func count_dots(dot_grid [][]string) int {
	count := 0
	for _, row := range dot_grid {
		for _, pos := range row {
			if pos == "#" {
				count++
			}
		}
	}
	return count
}

func calc_size_of_grid(dot_positions []position) (int, int) {
	x_max := 0
	y_max := 0
	for _, pos := range dot_positions {
		if pos.x > x_max {
			x_max = pos.x
		}
		if pos.y > y_max {
			y_max = pos.y
		}
	}
	return x_max, y_max
}

func build_empty_grid(x_max int, y_max int) [][]string {
	dot_grid := make([][]string, 0)
	for y := 0; y <= y_max; y++ {
		row := make([]string, 0)
		for x := 0; x <= x_max; x++ {
			row = append(row, ".")
		}
		dot_grid = append(dot_grid, row)
	}
	return dot_grid
}

func map_dots_onto_grid(dot_grid_pointer *[][]string, dot_positions []position) {
	dot_grid := *dot_grid_pointer
	for _, pos := range dot_positions {
		dot_grid[pos.y][pos.x] = "#"
	}
}

func print_dot_grid(dot_grid [][]string) {
	for _, row := range dot_grid {
		fmt.Println(row)
	}
}

func fold_grid(dot_grid [][]string, f fold) [][]string {
	if f.axis == "y" {
		new_dot_grid := build_empty_grid(len(dot_grid[0])-1, f.location-1)
		for y, row := range dot_grid {
			for x, pos := range row {
				if y < f.location {
					new_dot_grid[y][x] = pos
				} else if y > f.location && pos == "#" {
					new_y := y - (y-f.location)*2
					new_dot_grid[new_y][x] = pos
				}
			}
		}
		return new_dot_grid
	} else if f.axis == "x" {
		new_dot_grid := build_empty_grid(f.location-1, len(dot_grid)-1)
		for y, row := range dot_grid {
			for x, pos := range row {
				if x < f.location {
					new_dot_grid[y][x] = pos
				} else if x > f.location && pos == "#" {
					new_x := x - (x-f.location)*2
					new_dot_grid[y][new_x] = pos
				}
			}
		}
		return new_dot_grid
	}
	return make([][]string, 0)
}

func load_file(filename string) ([]position, []fold) {

	dot_positions := make([]position, 0)
	fold_list := make([]fold, 0)

	file, err := os.Open(filename)
	if err == nil {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			if strings.Contains(line, ",") {
				split_list := strings.Split(line, ",")
				x, _ := strconv.Atoi(split_list[0])
				y, _ := strconv.Atoi(split_list[1])
				dot_positions = append(dot_positions, position{x: x, y: y})
			} else if strings.Contains(line, "=") {
				split_list := strings.Split(strings.Fields(line)[2], "=")
				location, _ := strconv.Atoi(split_list[1])
				fold_list = append(fold_list, fold{axis: split_list[0], location: location})
			}
		}
	} else {
		fmt.Println("Error opening file")
		fmt.Println(err)
		os.Exit(23)
	}

	return dot_positions, fold_list
}

func process_file(filename string) {
	dot_positions, fold_list := load_file(filename)
	x_max, y_max := calc_size_of_grid(dot_positions)
	dot_grid := build_empty_grid(x_max, y_max)
	map_dots_onto_grid(&dot_grid, dot_positions)
	fmt.Println("~~~")
	for _, fold := range fold_list {
		dot_grid = fold_grid(dot_grid, fold)
		print_dot_grid(dot_grid)
		fmt.Println("Number of dots: ", count_dots(dot_grid))
		fmt.Println("~~~")
	}
}

func main() {
	process_file("day_13_test.txt")
	process_file("day_13.txt")
}
