package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func load_file(filename string) (map[string]string, []string) {
	pair_map := make(map[string]string)
	var polymer_chain []string

	file, err := os.Open(filename)
	if err == nil {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			if strings.Contains(line, ">") {
				var pair string
				var insert string
				fmt.Sscanf(line, "%s -> %s", &pair, &insert)
				pair_map[pair] = insert
			} else if line != "" {
				polymer_chain = strings.Split(line, "")
			}
		}
	} else {
		fmt.Println("Error opening file")
		fmt.Println(err)
		os.Exit(23)
	}

	return pair_map, polymer_chain
}

func make_empty_pair_count(pair_map map[string]string) map[string]int {
	pair_count := make(map[string]int)
	for key := range pair_map {
		pair_count[key] = 0
	}
	return pair_count
}

func run_step(pair_count map[string]int, pair_map map[string]string) map[string]int {
	new_pair_count := make_empty_pair_count(pair_map)
	for key, value := range pair_count {
		insert := pair_map[key]
		string_split := strings.Split(key, "")
		new_pair_count[string_split[0]+insert] += value
		new_pair_count[insert+string_split[1]] += value
	}
	return new_pair_count
}

func generate_initial_pair_count(polymer_chain []string, pair_map map[string]string) map[string]int {
	pair_count := make_empty_pair_count(pair_map)
	for i := 0; i < len(polymer_chain)-1; i++ {
		pair_count[polymer_chain[i]+polymer_chain[i+1]] += 1
	}
	return pair_count
}

func calc_result(pair_count map[string]int, last_element string) int {
	result_map := make(map[string]int)
	result_map[last_element] = 1
	for key, value := range pair_count {
		result_map[strings.Split(key, "")[0]] += value
	}
	max := 0
	for _, value := range result_map {
		if value > max {
			max = value
		}
	}
	min := max
	for _, value := range result_map {
		if value < min {
			min = value
		}
	}
	return max - min
}

func process_file(filename string, steps int) {
	pair_map, polymer_chain := load_file(filename)
	pair_count := generate_initial_pair_count(polymer_chain, pair_map)

	for i := 0; i < steps; i++ {
		pair_count = run_step(pair_count, pair_map)
	}

	fmt.Println("Result: ", calc_result(pair_count, polymer_chain[len(polymer_chain)-1]))
}

func main() {
	process_file("day_14_test.txt", 10)
	process_file("day_14.txt", 10)
	process_file("day_14_test.txt", 40)
	process_file("day_14.txt", 40)
}
