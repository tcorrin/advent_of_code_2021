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

type prioritized_position struct {
	pos  position
	cost int
}

func load_file(filename string) [][]int {
	chiton_density := make([][]int, 0)

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
			chiton_density = append(chiton_density, row)
		}
	} else {
		fmt.Println("Error opening file")
		fmt.Println(err)
		os.Exit(23)
	}

	return chiton_density
}

func get_neighbours(lookup position, x_max int, y_max int) []position {
	result_list := make([]position, 0)
	if lookup.x < x_max {
		result_list = append(result_list, position{x: lookup.x + 1, y: lookup.y})
	}
	if lookup.y < y_max {
		result_list = append(result_list, position{x: lookup.x, y: lookup.y + 1})
	}
	if lookup.x > 0 {
		result_list = append(result_list, position{x: lookup.x - 1, y: lookup.y})
	}
	if lookup.y > 0 {
		result_list = append(result_list, position{x: lookup.x, y: lookup.y - 1})
	}
	return result_list
}

func insert_position(pos_list_pointer *[]prioritized_position, pos position, cost int) {
	lower_index := 0
	upper_index := 0
	*pos_list_pointer = append(*pos_list_pointer, prioritized_position{})
	pos_list := *pos_list_pointer
	if len(pos_list) == 1 {
		pos_list[0].pos = pos
	} else if cost <= pos_list[0].cost {
		copy(pos_list[1:], pos_list[0:])
		pos_list[0].pos = pos
	} else if cost >= pos_list[len(pos_list)-2].cost {
		pos_list[len(pos_list)-1].pos = pos
	} else {
		for i := 0; i < len(pos_list)-1; i++ {
			if pos_list[i].cost <= cost && pos_list[i+1].cost >= cost {
				lower_index = i
				upper_index = i + 1
			}
		}
		copy(pos_list[upper_index:], pos_list[lower_index:])
		pos_list[lower_index].pos = pos
	}
}

func process_file(filname string, mega_map bool) {
	chiton_density := load_file(filname)
	if mega_map {
		chiton_density = make_mega_map(chiton_density)
	}
	x_max := len(chiton_density[0]) - 1
	y_max := len(chiton_density) - 1
	goal := position{x: x_max, y: y_max}

	pos_list := make([]prioritized_position, 0)
	pos_list = append(pos_list, prioritized_position{pos: position{x: 0, y: 0}, cost: 0})
	visited := make(map[position]position)
	cost_so_far := make(map[position]int)

	for len(pos_list) > 0 {
		current_position := pos_list[0]
		pos_list = pos_list[1:]
		neighbours := get_neighbours(current_position.pos, x_max, y_max)
		for _, neighbour := range neighbours {
			new_cost := cost_so_far[current_position.pos] + chiton_density[neighbour.y][neighbour.x]
			if cost_so_far[neighbour] == 0 || new_cost < cost_so_far[neighbour] {
				cost_so_far[neighbour] = new_cost
				insert_position(&pos_list, neighbour, new_cost)
				visited[neighbour] = current_position.pos
			}
		}
	}

	fmt.Println("Min path cost: ", cost_so_far[goal])
}

func map_number(input int) int {
	if input < 10 {
		return input
	} else {
		return (input % 10) + 1
	}
}

func make_mega_map(chiton_density [][]int) [][]int {
	width := len(chiton_density[0])
	height := len(chiton_density)
	mega_width := width * 5
	mega_height := height * 5

	mega_map := make([][]int, 0)
	for y := 0; y < mega_height; y++ {
		row := make([]int, 0)
		for x := 0; x < mega_width; x++ {
			row = append(row, 0)
		}
		mega_map = append(mega_map, row)
	}

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			mega_map[y][x] = chiton_density[y][x]
			mega_map[y][x+(1*width)] = map_number(chiton_density[y][x] + 1)
			mega_map[y][x+(2*width)] = map_number(chiton_density[y][x] + 2)
			mega_map[y][x+(3*width)] = map_number(chiton_density[y][x] + 3)
			mega_map[y][x+(4*width)] = map_number(chiton_density[y][x] + 4)
			mega_map[y+(1*width)][x] = map_number(chiton_density[y][x] + 1)
			mega_map[y+(1*width)][x+(1*width)] = map_number(chiton_density[y][x] + 2)
			mega_map[y+(1*width)][x+(2*width)] = map_number(chiton_density[y][x] + 3)
			mega_map[y+(1*width)][x+(3*width)] = map_number(chiton_density[y][x] + 4)
			mega_map[y+(1*width)][x+(4*width)] = map_number(chiton_density[y][x] + 5)
			mega_map[y+(2*width)][x] = map_number(chiton_density[y][x] + 2)
			mega_map[y+(2*width)][x+(1*width)] = map_number(chiton_density[y][x] + 3)
			mega_map[y+(2*width)][x+(2*width)] = map_number(chiton_density[y][x] + 4)
			mega_map[y+(2*width)][x+(3*width)] = map_number(chiton_density[y][x] + 5)
			mega_map[y+(2*width)][x+(4*width)] = map_number(chiton_density[y][x] + 6)
			mega_map[y+(3*width)][x] = map_number(chiton_density[y][x] + 3)
			mega_map[y+(3*width)][x+(1*width)] = map_number(chiton_density[y][x] + 4)
			mega_map[y+(3*width)][x+(2*width)] = map_number(chiton_density[y][x] + 5)
			mega_map[y+(3*width)][x+(3*width)] = map_number(chiton_density[y][x] + 6)
			mega_map[y+(3*width)][x+(4*width)] = map_number(chiton_density[y][x] + 7)
			mega_map[y+(4*width)][x] = map_number(chiton_density[y][x] + 4)
			mega_map[y+(4*width)][x+(1*width)] = map_number(chiton_density[y][x] + 5)
			mega_map[y+(4*width)][x+(2*width)] = map_number(chiton_density[y][x] + 6)
			mega_map[y+(4*width)][x+(3*width)] = map_number(chiton_density[y][x] + 7)
			mega_map[y+(4*width)][x+(4*width)] = map_number(chiton_density[y][x] + 8)
		}
	}
	return mega_map
}

func main() {
	process_file("day_15_test.txt", false)
	process_file("day_15.txt", false)
	process_file("day_15_test.txt", true)
	process_file("day_15.txt", true)
}
