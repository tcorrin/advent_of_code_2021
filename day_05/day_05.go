package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type position struct {
	x int
	y int
}

func process_position_string(input string) position {
	split_list := strings.Split(input, ",")
	x, _ := strconv.Atoi(split_list[0])
	y, _ := strconv.Atoi(split_list[1])
	return position{x: x, y: y}
}

func calculate_max_values(list []position, x_max *int, y_max *int) {
	for _, p := range list {
		if p.x > *x_max {
			*x_max = p.x
		}
		if p.y > *y_max {
			*y_max = p.y
		}
	}
}

func initialise_grid(x_max int, y_max int) [][]int {
	floor_grid := make([][]int, y_max)
	for i := 0; i < y_max; i++ {
		floor_grid[i] = make([]int, x_max)
		for j := 0; j < x_max; j++ {
			floor_grid[i][j] = 0
		}
	}
	return floor_grid
}

func load_file(filename string) ([]position, []position) {
	start_list := make([]position, 0)
	end_list := make([]position, 0)
	re := regexp.MustCompile(" -> ")

	file, err := os.Open(filename)
	if err == nil {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			split_list := re.Split(line, -1)
			start_list = append(start_list, process_position_string(split_list[0]))
			end_list = append(end_list, process_position_string(split_list[1]))
		}
	}

	return start_list, end_list
}

func filter_list(start_list []position, end_list []position, keep_list []int) ([]position, []position) {
	filtered_start_list := make([]position, 0)
	filtered_end_list := make([]position, 0)
	for i := 0; i < len(start_list); i++ {
		for _, v := range keep_list {
			if i == v {
				filtered_start_list = append(filtered_start_list, start_list[i])
				filtered_end_list = append(filtered_end_list, end_list[i])
			}
		}
	}
	return filtered_start_list, filtered_end_list
}

func calc_direction(start int, end int) int {
	if start < end {
		return 1
	} else if start > end {
		return -1
	} else {
		return 0
	}
}

func draw_line(start position, end position, fg *[][]int) {
	floor_grid := *fg
	x_direction := calc_direction(start.x, end.x)
	y_direction := calc_direction(start.y, end.y)
	magnitude := 0

	if x_direction == 0 {
		magnitude = (end.y - start.y + y_direction) * y_direction
	} else {
		magnitude = (end.x - start.x + x_direction) * x_direction
	}

	for pos := 0; pos < magnitude; pos++ {
		floor_grid[start.y+(pos*y_direction)][start.x+(pos*x_direction)]++
	}

}

func process_file(filename string, diagonal bool) {
	start_list, end_list := load_file(filename)
	keep_list := make([]int, 0)
	if !diagonal {
		for i := 0; i < len(start_list); i++ {
			if start_list[i].y == end_list[i].y || start_list[i].x == end_list[i].x {
				keep_list = append(keep_list, i)
			}
		}

		start_list, end_list = filter_list(start_list, end_list, keep_list)
	}

	x_max := 0
	y_max := 0

	calculate_max_values(start_list, &x_max, &y_max)
	calculate_max_values(end_list, &x_max, &y_max)

	floor_grid := initialise_grid(x_max+1, y_max+1)

	for i := 0; i < len(start_list); i++ {
		draw_line(start_list[i], end_list[i], &floor_grid)
	}

	danger_count := 0

	for _, row := range floor_grid {
		for _, col := range row {
			if col > 1 {
				danger_count++
			}
		}
	}

	fmt.Println("Danger Count: ", danger_count)
}

func main() {
	process_file("day_05_test.txt", false)
	process_file("day_05.txt", false)
	process_file("day_05_test.txt", true)
	process_file("day_05.txt", true)
}
