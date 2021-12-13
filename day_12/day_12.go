package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func contains_path(list_of_paths [][]string, lookup_path []string) bool {
	for _, path := range list_of_paths {
		match := true
		for i, cave := range path {
			if lookup_path[i] != cave {
				match = false
				break
			}
		}
		if match {
			return true
		}
	}
	return false
}

func is_small_cave(cave_name string) bool {
	for _, letter := range cave_name {
		if unicode.IsUpper(letter) {
			return false
		}
	}
	return true
}

func contains_cave(path []string, cave_name string, limit int) bool {
	count := 0
	for _, cave := range path {
		if cave == cave_name {
			count++
			if count == limit {
				return true
			}
		}
	}
	return false
}

func add_cave(path []string, cave_name string) []string {
	new_path := make([]string, len(path))
	copy(new_path, path)
	new_path = append(new_path, cave_name)
	return new_path
}

func process_line(line string, path_map map[string][]string) {
	split_list := strings.Split(line, "-")
	node_one := split_list[0]
	node_two := split_list[1]
	if node_one == "start" {
		path_map[node_one] = append(path_map[node_one], node_two)
	} else if node_two == "start" {
		path_map[node_two] = append(path_map[node_two], node_one)
	} else if node_one == "end" {
		path_map[node_two] = append(path_map[node_two], node_one)
	} else if node_two == "end" {
		path_map[node_one] = append(path_map[node_one], node_two)
	} else {
		path_map[node_one] = append(path_map[node_one], node_two)
		path_map[node_two] = append(path_map[node_two], node_one)
	}
}

func next_step_in_path(list_of_paths [][]string, path_map map[string][]string, double_dip string) [][]string {
	new_list_of_paths := make([][]string, 0)
	for _, path := range list_of_paths {
		last_cave := path[len(path)-1]
		if last_cave == "end" {
			new_list_of_paths = append(new_list_of_paths, path)
		} else {
			connecting_caves := path_map[last_cave]
			for _, cave := range connecting_caves {
				if double_dip == "" && !(is_small_cave(cave) && contains_cave(path, cave, 1)) {
					new_path := add_cave(path, cave)
					if !(cave == "end" && contains_path(new_list_of_paths, new_path)) {
						new_list_of_paths = append(new_list_of_paths, new_path)
					}
				} else if !(cave == double_dip && is_small_cave(cave) && contains_cave(path, cave, 2)) &&
					!(cave != double_dip && is_small_cave(cave) && contains_cave(path, cave, 1)) {
					new_path := add_cave(path, cave)
					if !(cave == "end" && contains_path(new_list_of_paths, new_path)) {
						new_list_of_paths = append(new_list_of_paths, new_path)
					}
				}
			}
		}
	}
	return new_list_of_paths
}

func check_for_paths_complete(list_of_paths [][]string) bool {
	for _, path := range list_of_paths {
		if path[len(path)-1] != "end" {
			return false
		}
	}
	return true
}

func load_file(filename string) map[string][]string {
	path_map := make(map[string][]string)
	file, err := os.Open(filename)
	if err == nil {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			process_line(scanner.Text(), path_map)
		}
	} else {
		fmt.Println("Error opening file")
		fmt.Println(err)
		os.Exit(23)
	}
	return path_map
}

func process_file(filename string) {
	path_map := load_file(filename)
	list_of_paths := [][]string{{"start"}}
	for !check_for_paths_complete(list_of_paths) {
		list_of_paths = next_step_in_path(list_of_paths, path_map, "")
	}

	fmt.Println("Number of paths: ", len(list_of_paths))
	//get list of small caves
	small_cave_list := make([]string, 0)
	for entry_cave, cave_list := range path_map {
		if is_small_cave(entry_cave) && !contains_cave(small_cave_list, entry_cave, 1) && entry_cave != "start" {
			small_cave_list = append(small_cave_list, entry_cave)
		}
		for _, cave := range cave_list {
			if is_small_cave(cave) && !contains_cave(small_cave_list, cave, 2) && cave != "end" {
				small_cave_list = append(small_cave_list, cave)
			}
		}
	}

	master_list_of_paths := make([][]string, 0)
	for _, double_dip_small_cave := range small_cave_list {
		sub_list_of_paths := [][]string{{"start"}}
		for !check_for_paths_complete(sub_list_of_paths) {
			sub_list_of_paths = next_step_in_path(sub_list_of_paths, path_map, double_dip_small_cave)
		}
		for _, path := range sub_list_of_paths {
			if !contains_path(master_list_of_paths, path) {
				master_list_of_paths = append(master_list_of_paths, path)
			}
		}
	}

	fmt.Println("Number of paths when double dipping: ", len(master_list_of_paths))

}

func main() {
	process_file("day_12_test01.txt")
	process_file("day_12_test02.txt")
	process_file("day_12_test03.txt")
	process_file("day_12.txt")
}
