package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func is_int(input string) bool {
	_, err := strconv.Atoi(input)
	return err == nil
}

func load_file(filename string) []string {
	result := make([]string, 0)
	file, err := os.Open(filename)
	if err == nil {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			result = append(result, scanner.Text())
		}
	} else {
		fmt.Println("Error opening file")
		fmt.Println(err)
		os.Exit(23)
	}
	return result
}

func pop_number(number_list *[]string) string {
	result := (*number_list)[0]
	*number_list = (*number_list)[1:]
	return result
}

func add_numbers(lhs string, rhs string) string {
	number := "[" + lhs + "," + rhs + "]"
	for can_be_exploded(number) || can_be_split(number) {
		reduce_number(&number)
	}
	return number
}

func can_be_exploded(number string) bool {
	max_depth := 0
	split_list := strings.Split(number, "")
	depth := 0
	for _, x := range split_list {
		if x == "[" {
			depth++
		} else if x == "]" {
			depth--
		}
		if depth > max_depth {
			max_depth = depth
		}
	}
	return max_depth >= 5
}

func can_be_split(number string) bool {
	split_list := strings.Split(number, "")
	for i, x := range split_list {
		if is_int(x) && is_int(split_list[i+1]) {
			return true
		}
	}
	return false
}

func split_number(number *string) {
	split_list := strings.Split(*number, "")
	new_split := make([]string, 0)
	split := false
	var extra_digit_index int
	for i, x := range split_list {
		if is_int(x) && is_int(split_list[i+1]) && !split {
			split = true
			extra_digit_index = i + 5
			value_string := x + split_list[i+1]
			value, _ := strconv.Atoi(value_string)
			lhs := strconv.Itoa(value / 2)
			rhs := strconv.Itoa(value/2 + value%2)
			new_split = append(new_split, "[")
			new_split = append(new_split, lhs)
			new_split = append(new_split, ",")
			new_split = append(new_split, rhs)
			new_split = append(new_split, "]")
		} else {
			new_split = append(new_split, x)
		}
	}
	new_split = append(new_split[:extra_digit_index], new_split[extra_digit_index+1:]...)
	*number = strings.Join(new_split, "")
}

func explode_number(number *string) {
	split_list := strings.Split(*number, "")
	new_split := make([]string, 0)
	depth := 0
	explosion := false
	previous_number_index := -1
	next_number_index := -1
	skip := 0
	rhs := -1
	lhs := -1
	trim := false
	for i, x := range split_list {
		if x == "[" {
			depth++
		} else if x == "]" {
			depth--
		}
		if depth < 5 || explosion {
			if depth == 4 && x == "]" && !explosion {
				explosion = true
				new_split = append(new_split, "0")
			} else {
				if is_int(x) {
					if !explosion {
						previous_number_index = i
					} else if next_number_index == -1 {
						next_number_index = i
					}
				}
				new_split = append(new_split, x)
			}
		} else {
			if is_int(x) && !is_int(split_list[i+1]) {
				if skip == 1 && lhs == -1 {
					lhs, _ = strconv.Atoi(x)
				} else if (skip == 3 || skip == 4) && rhs == -1 {
					rhs, _ = strconv.Atoi(x)
				}
			} else if is_int(x) && is_int(split_list[i+1]) {
				if skip == 1 && lhs == -1 {
					lhs, _ = strconv.Atoi(x + split_list[i+1])
				} else if (skip == 3 || skip == 4) && rhs == -1 {
					rhs, _ = strconv.Atoi(x + split_list[i+1])
				}
			}
			skip++
		}
	}
	if previous_number_index != -1 {
		if is_int(new_split[previous_number_index]) {
			var value int
			if !is_int(new_split[previous_number_index-1]) {
				value, _ = strconv.Atoi(new_split[previous_number_index])
			} else {
				value, _ = strconv.Atoi(new_split[previous_number_index-1] + new_split[previous_number_index])
			}
			new_split[previous_number_index] = strconv.Itoa(value + lhs)
			if value > 9 {
				trim = true
				new_split = append(new_split[:previous_number_index-1], new_split[previous_number_index:]...)
			}

		}
	}
	if next_number_index != -1 {
		i := next_number_index - skip
		if trim {
			i--
		}
		if is_int(new_split[i]) {
			var value int
			if !is_int(new_split[i+1]) {
				value, _ = strconv.Atoi(new_split[i])
			} else {
				value, _ = strconv.Atoi(new_split[i] + new_split[i+1])
			}
			new_split[i] = strconv.Itoa(value + rhs)
			if value > 9 {
				new_split = append(new_split[:i+1], new_split[i+2:]...)
			}
		}
	}
	*number = strings.Join(new_split, "")
}

func reduce_number(number *string) {
	if can_be_exploded(*number) {
		explode_number(number)
	} else if can_be_split(*number) {
		split_number(number)
	}
}

func calc_magnitude(number string) int {
	max_depth := 0
	depth := 0
	result := 0
	split_list := strings.Split(number, "")
	for _, x := range split_list {
		if x == "[" {
			depth++
		} else if x == "]" {
			depth--
		}
		if depth > max_depth {
			max_depth = depth
		}
	}
	new_split := make([]string, 0)
	for i := max_depth; i > 0; i-- {
		skip := 0
		for j, x := range split_list {
			if x == "[" {
				depth++
			} else if x == "]" {
				depth--
			}
			if depth == i {
				if is_int(x) && skip == 1 {
					rhs, _ := strconv.Atoi(x)
					rhs = rhs * 3
					lhs, _ := strconv.Atoi(split_list[j+2])
					lhs = lhs * 2
					new_split = append(new_split, strconv.Itoa(lhs+rhs))
				}
				skip++
			} else if !(depth == i-1 && x == "]") {
				new_split = append(new_split, x)
				skip = 0
			}
		}
		split_list = new_split
		new_split = make([]string, 0)
	}
	result, _ = strconv.Atoi(split_list[0])
	return result
}

func calc_magnitude_test() {
	input := make([]string, 0)
	input = append(input, "[9,1]")
	input = append(input, "[[9,1],[1,9]]")
	input = append(input, "[[1,2],[[3,4],5]]")
	input = append(input, "[[[[0,7],4],[[7,8],[6,0]]],[8,1]]")
	input = append(input, "[[[[1,1],[2,2]],[3,3]],[4,4]]")
	input = append(input, "[[[[3,0],[5,3]],[4,4]],[5,5]]")
	input = append(input, "[[[[5,0],[7,4]],[5,5]],[6,6]]")
	input = append(input, "[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]")
	input = append(input, "[[[[6,6],[7,6]],[[7,7],[7,0]]],[[[7,7],[7,7]],[[7,8],[9,9]]]]")
	input = append(input, "[[[[7,7],[7,7]],[[7,8],[7,8]]],[[[8,7],[8,0]],[[8,8],[8,8]]]]")

	expected_results := make([]int, 0)
	expected_results = append(expected_results, 29)
	expected_results = append(expected_results, 129)
	expected_results = append(expected_results, 143)
	expected_results = append(expected_results, 1384)
	expected_results = append(expected_results, 445)
	expected_results = append(expected_results, 791)
	expected_results = append(expected_results, 1137)
	expected_results = append(expected_results, 3488)
	expected_results = append(expected_results, 4140)
	expected_results = append(expected_results, 4457)

	for i := 0; i < len(input); i++ {
		result := calc_magnitude(input[i])
		if result == expected_results[i] {
			fmt.Println("PASS -", input[i], "-", expected_results[i])
		} else {
			fmt.Println("FAIL -", input[i], "- Got:", result, " Expected:", expected_results[i])
		}
	}
}

func process_file(filename string) {
	number_list := load_file(filename)
	current_number := pop_number(&number_list)
	for len(number_list) > 0 {
		current_number = add_numbers(current_number, pop_number(&number_list))
	}
	fmt.Println(filename, "Pt1 - Final number:", current_number)
	magnitude := calc_magnitude(current_number)
	fmt.Println(filename, "Pt1 - Final number magnitude:", magnitude)
	max_magnitude := 0
	number_list = load_file(filename)
	for i := 0; i < len(number_list); i++ {
		for j := 0; j < len(number_list); j++ {
			if i != j {
				result := add_numbers(number_list[i], number_list[j])
				current_magnitude := calc_magnitude(result)
				if current_magnitude > max_magnitude {
					max_magnitude = current_magnitude
				}
			}
		}
	}
	fmt.Println(filename, "Pt2 - Max magnitude:", max_magnitude)
}

func main() {
	calc_magnitude_test()
	process_file("day_18_test_01.txt")
	process_file("day_18_test_02.txt")
	process_file("day_18_test_03.txt")
	process_file("day_18_test_05.txt")
	process_file("day_18_test_04.txt")
	process_file("day_18.txt")
}
