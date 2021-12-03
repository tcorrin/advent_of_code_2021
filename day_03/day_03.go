package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func load_file(filename string) [][]string {

	file, err := os.Open(filename)
	number_list := make([][]string, 0)

	if err == nil {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line_slice := strings.Split(scanner.Text(), "")
			number_list = append(number_list, line_slice)
		}
	} else {
		fmt.Println("Error opening file")
		fmt.Println(err)
	}

	return number_list
}

func invert_binary_number(binary_number string) string {
	result := ""
	for _, bit := range binary_number {
		if bit == '1' {
			result += "0"
		} else if bit == '0' {
			result += "1"
		}
	}
	return result
}

func calculate_common_bit(number_list [][]string, pos int, match bool) string {
	number_of_ones := 0
	number_of_zeros := 0

	for _, number := range number_list {
		if number[pos] == "1" {
			number_of_ones++
		} else if number[pos] == "0" {
			number_of_zeros++
		}
	}

	if match {
		if number_of_ones > number_of_zeros {
			return "1"
		} else if number_of_zeros > number_of_ones {
			return "0"
		} else {
			return "1"
		}
	} else {
		if number_of_ones > number_of_zeros {
			return "0"
		} else if number_of_zeros > number_of_ones {
			return "1"
		} else {
			return "0"
		}
	}
}

func filter_number_list_on_bit(bit string, pos int, number_list [][]string) [][]string {
	result_list := make([][]string, 0)
	for _, number := range number_list {
		if number[pos] == bit {
			result_list = append(result_list, number)
		}
	}
	return result_list
}

func process_file(filename string) {
	number_list := load_file(filename)
	number_of_bits := len(number_list[0])
	result := ""

	fmt.Println("Number of bits: ", number_of_bits)

	for i := 0; i < number_of_bits; i++ {
		result += calculate_common_bit(number_list, i, true)
	}

	base_gamma_rate := result
	base_epsilon_rate := invert_binary_number(result)

	gamma_rate, _ := strconv.ParseUint(fmt.Sprintf("%064s", base_gamma_rate), 2, 64)
	epsilon_rate, _ := strconv.ParseUint(fmt.Sprintf("%064s", base_epsilon_rate), 2, 64)

	power_consumption := gamma_rate * epsilon_rate

	println("pt1 gamma_rate: ", gamma_rate, " epsilon_rate: ", epsilon_rate, " power_consumption: ", power_consumption)

	oxygen_rate_list := make([][]string, len(number_list))
	copy(oxygen_rate_list, number_list)
	oxygen_rate_binary := ""

	scrubber_rate_list := make([][]string, len(number_list))
	copy(scrubber_rate_list, number_list)
	scrubber_rating_binary := ""

	for i := 0; i < number_of_bits; i++ {
		if len(oxygen_rate_list) > 1 {
			bit := calculate_common_bit(oxygen_rate_list, i, true)
			oxygen_rate_list = filter_number_list_on_bit(string(bit), i, oxygen_rate_list)
		}
		if len(oxygen_rate_list) == 1 {
			oxygen_rate_binary = strings.Join(oxygen_rate_list[0], "")
			break
		}
	}

	for i := 0; i <= number_of_bits; i++ {
		if len(scrubber_rate_list) > 1 {
			bit := calculate_common_bit(scrubber_rate_list, i, false)
			scrubber_rate_list = filter_number_list_on_bit(string(bit), i, scrubber_rate_list)
		}
		if len(scrubber_rate_list) == 1 {
			scrubber_rating_binary = strings.Join(scrubber_rate_list[0], "")
			break
		}
	}

	oxygen_rate, _ := strconv.ParseUint(fmt.Sprintf("%064s", oxygen_rate_binary), 2, 64)
	scrubber_rating, _ := strconv.ParseUint(fmt.Sprintf("%064s", scrubber_rating_binary), 2, 64)
	life_support_rating := oxygen_rate * scrubber_rating

	println("pt2 oxygen_rate: ", oxygen_rate, " scrubber_rating: ", scrubber_rating, " life_support_rating: ", life_support_rating)
}

func main() {
	process_file("day_03_test.txt")
	process_file("day_03.txt")
}
