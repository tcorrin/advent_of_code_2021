package main

import "fmt"

type probe struct {
	y_pos int
	x_pos int
	y_vel int
	x_vel int
}

type target_area struct {
	x_min int
	x_max int
	y_min int
	y_max int
}

const (
	Short = 0
	Hit   = 1
	Long  = 2
)

func print_result(result int) string {
	switch result {
	case 0:
		return "Short"
	case 1:
		return "Hit"
	case 2:
		return "Long"
	}
	return "unknown"
}

func run_step(search_probe *probe, search_area target_area) int {
	search_probe.x_pos += search_probe.x_vel
	search_probe.y_pos += search_probe.y_vel
	if search_probe.x_vel > 0 {
		search_probe.x_vel--
	} else if search_probe.x_vel < 0 {
		search_probe.x_vel++
	}
	search_probe.y_vel--

	if search_probe.x_pos > search_area.x_max || search_probe.y_pos < search_area.y_min {
		return Long
	}

	if search_probe.x_pos <= search_area.x_max &&
		search_probe.x_pos >= search_area.x_min &&
		search_probe.y_pos <= search_area.y_max &&
		search_probe.y_pos >= search_area.y_min {
		return Hit
	}

	return Short
}

func shoot_probe(search_area target_area, initial_x_vel int, initial_y_vel int) (int, int) {
	max_height := 0
	search_probe := probe{x_pos: 0, y_pos: 0, x_vel: initial_x_vel, y_vel: initial_y_vel}
	result := Short
	for result == Short {
		result = run_step(&search_probe, search_area)
		if search_probe.y_pos > max_height {
			max_height = search_probe.y_pos
		}
	}
	return result, max_height
}

func process_target_area(search_area target_area) {
	max_height := 0
	hit_counter := 0
	for initial_x_vel := 0; initial_x_vel < 5000; initial_x_vel++ {
		for initial_y_vel := -5000; initial_y_vel < 5000; initial_y_vel++ {
			result, shot_max_height := shoot_probe(search_area, initial_x_vel, initial_y_vel)
			if result == Hit {
				hit_counter++
				if shot_max_height > max_height {
					max_height = shot_max_height
				}
			}
		}
	}
	fmt.Println("Max Height:", max_height)
	fmt.Println("Hit Counter:", hit_counter)
}

func main() {
	process_target_area(target_area{x_min: 20, x_max: 30, y_min: -10, y_max: -5})
	process_target_area(target_area{x_min: 156, x_max: 202, y_min: -110, y_max: -69})
}
