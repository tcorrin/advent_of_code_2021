package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func remove_string(list []string, item string) []string {
	new_list := make([]string, 0)
	for _, value := range list {
		if value != item {
			new_list = append(new_list, value)
		}
	}
	return new_list
}

func contains_int(list []int, target int) bool {
	for _, value := range list {
		if value == target {
			return true
		}
	}
	return false
}

func contains_string(list []string, target string) bool {
	for _, value := range list {
		if value == target {
			return true
		}
	}
	return false
}

func elements_of_length(list []string, length int) []string {
	results := make([]string, 0)
	for _, value := range list {
		if len(value) == length {
			results = append(results, value)
		}
	}
	return results
}

func load_file(filename string) ([][]string, [][]string) {
	notes := make([][]string, 0)
	output := make([][]string, 0)

	file, err := os.Open(filename)
	if err == nil {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			split_list := strings.Split(scanner.Text(), "|")
			notes = append(notes, strings.Fields(split_list[0]))
			output = append(output, strings.Fields(split_list[1]))
		}
	} else {
		fmt.Println("Error opening file")
		fmt.Println(err)
		os.Exit(23)
	}

	return notes, output
}

func find_digit_by_matching_element(input string, input_list []string) string {
	result := ""
	input_split := strings.Split(input, "")
	for _, value := range input_list {
		value_split := strings.Split(value, "")
		found := true
		for _, letter := range input_split {
			if !contains_string(value_split, letter) {
				found = false
				break
			}
		}
		if found {
			result = value
			break
		}
	}
	return result
}

func find_digit_by_missing_element_count(input string, input_list []string, expected_count int) string {
	result := ""
	input_split := strings.Split(input, "")
	for _, value := range input_list {
		value_split := strings.Split(value, "")
		count := 0
		for _, letter := range input_split {
			if !contains_string(value_split, letter) {
				count++
			}
		}
		if count == expected_count {
			result = value
			break
		}
	}
	return result
}

func find_number_from_digit_map(digit_map map[string]string, lookup string) string {
	result := ""
	lookup_split := strings.Split(lookup, "")
	for key, value := range digit_map {
		found := true
		if len(key) == len(lookup) {
			key_split := strings.Split(key, "")
			for _, letter := range key_split {
				if !contains_string(lookup_split, letter) {
					found = false
				}
			}
			if found {
				result = value
				break
			}
		}
	}

	return result
}

func process_line(notes []string, output []string) int {
	digit_map := make(map[string]string)

	one := elements_of_length(notes, 2)[0]
	digit_map[one] = "1"

	four := elements_of_length(notes, 4)[0]
	digit_map[four] = "4"

	seven := elements_of_length(notes, 3)[0]
	digit_map[seven] = "7"

	eight := elements_of_length(notes, 7)[0]
	digit_map[eight] = "8"

	six_list := elements_of_length(notes, 6)
	five_list := elements_of_length(notes, 5)

	six := find_digit_by_missing_element_count(one, six_list, 1)
	digit_map[six] = "6"
	six_list = remove_string(six_list, six)

	zero := find_digit_by_missing_element_count(four, six_list, 1)
	digit_map[zero] = "0"

	six_list = remove_string(six_list, zero)

	nine := six_list[0]
	digit_map[nine] = "9"

	three := find_digit_by_matching_element(one, five_list)
	digit_map[three] = "3"

	five_list = remove_string(five_list, three)

	two := find_digit_by_missing_element_count(four, five_list, 2)
	digit_map[two] = "2"

	five_list = remove_string(five_list, two)
	five := five_list[0]
	digit_map[five] = "5"

	total_output := ""

	for _, value := range output {
		total_output += find_number_from_digit_map(digit_map, value)
	}

	total_output_int, _ := strconv.Atoi(total_output)

	return total_output_int
}

func process_file(filename string) {
	notes, output := load_file(filename)
	easy_digits := []int{2, 4, 3, 7}
	easy_digit_count := 0

	for _, output_line := range output {
		for _, value := range output_line {
			if contains_int(easy_digits, len(value)) {
				easy_digit_count++
			}
		}
	}

	fmt.Println("Easy Digit Count: ", easy_digit_count)

	total_output_value := 0

	for i := range output {
		total_output_value += process_line(notes[i], output[i])
	}

	fmt.Println("Total output value: ", total_output_value)
}

func main() {
	process_file("day_08_test.txt")
	process_file("day_08.txt")
}
