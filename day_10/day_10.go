package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func reverse(list []string) []string {
	length := len(list)
	result := make([]string, length)
	for i := 0; i < length; i++ {
		result[length-i-1] = list[i]
	}
	return result
}

func load_file(filename string) [][]string {
	line_list := make([][]string, 0)

	file, err := os.Open(filename)
	if err == nil {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line_list = append(line_list, strings.Split(scanner.Text(), ""))
		}
	} else {
		fmt.Println("Error opening file")
		fmt.Println(err)
		os.Exit(23)
	}

	return line_list
}

func is_open_bracket(value string) bool {
	if value == "(" || value == "[" || value == "{" || value == "<" {
		return true
	}
	return false
}

func is_close_bracket(value string) bool {
	if value == ")" || value == "]" || value == "}" || value == ">" {
		return true
	}
	return false
}

func is_matching_bracket(open string, close string) bool {
	result := false
	if open == "(" && close == ")" {
		result = true
	} else if open == "{" && close == "}" {
		result = true
	} else if open == "[" && close == "]" {
		result = true
	} else if open == "<" && close == ">" {
		result = true
	}
	return result
}

func process_file(filename string) {
	line_list := load_file(filename)
	illegal_bracket_list := make([]string, 0)
	left_over_brackets_list := make([][]string, 0)

	for _, line := range line_list {
		open_bracket_list := make([]string, 0)
		illegal := false
		for _, bracket := range line {
			if is_open_bracket(bracket) {
				open_bracket_list = append(open_bracket_list, bracket)
			} else if is_close_bracket(bracket) {
				if is_matching_bracket(open_bracket_list[len(open_bracket_list)-1], bracket) {
					open_bracket_list[len(open_bracket_list)-1] = ""
					open_bracket_list = open_bracket_list[:len(open_bracket_list)-1]
				} else {
					illegal_bracket_list = append(illegal_bracket_list, bracket)
					illegal = true
					break
				}
			}
		}
		if !illegal {
			left_over_brackets_list = append(left_over_brackets_list, open_bracket_list)
		}
	}

	bracket_sytax_error_map := map[string]int{")": 3, "]": 57, "}": 1197, ">": 25137}
	total_syntax_error_score := 0

	for _, bracket := range illegal_bracket_list {
		total_syntax_error_score += bracket_sytax_error_map[bracket]
	}

	fmt.Println("Total syntax error score: ", total_syntax_error_score)

	autocomplete_score_map := map[string]int{"(": 1, "[": 2, "{": 3, "<": 4}
	autocomplete_score_list := make([]int, 0)

	for _, line := range left_over_brackets_list {
		autocomplete_score := 0
		reversed_line := reverse(line)
		for _, bracket := range reversed_line {
			autocomplete_score = autocomplete_score*5 + autocomplete_score_map[bracket]
		}
		autocomplete_score_list = append(autocomplete_score_list, autocomplete_score)
	}

	sort.Ints(autocomplete_score_list)

	fmt.Println("Winning score: ", autocomplete_score_list[len(autocomplete_score_list)/2])
}

func main() {
	process_file("day_10_test.txt")
	process_file("day_10.txt")
}
