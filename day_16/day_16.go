package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const MaxInt = int64(9223372036854775807)

func load_file(filename string) string {
	result := ""
	file, err := os.Open(filename)
	if err == nil {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			result = scanner.Text()
			break
		}
	} else {
		fmt.Println("Error opening file")
		fmt.Println(err)
		os.Exit(23)
	}
	return result
}

func decode_hex_to_binary(hex_message string) string {
	binary_message := ""
	decode_map := make(map[string]string)
	decode_map["0"] = "0000"
	decode_map["1"] = "0001"
	decode_map["2"] = "0010"
	decode_map["3"] = "0011"
	decode_map["4"] = "0100"
	decode_map["5"] = "0101"
	decode_map["6"] = "0110"
	decode_map["7"] = "0111"
	decode_map["8"] = "1000"
	decode_map["9"] = "1001"
	decode_map["A"] = "1010"
	decode_map["B"] = "1011"
	decode_map["C"] = "1100"
	decode_map["D"] = "1101"
	decode_map["E"] = "1110"
	decode_map["F"] = "1111"

	for _, hex_value := range strings.Split(hex_message, "") {
		binary_message += decode_map[hex_value]
	}
	return binary_message
}

func convert_to_int(binary_string string) int64 {
	b := fmt.Sprintf("%064s", binary_string)
	i, _ := strconv.ParseInt(b, 2, 64)
	return i
}

func process_literal_value_packet(packet *string) int64 {
	carry_on := true
	binary_output := ""
	for carry_on {
		substring := (*packet)[:5]
		carry_on = substring[:1] == "1"
		binary_output += substring[1:]
		*packet = (*packet)[5:]
	}
	return convert_to_int(binary_output)
}

func calculate_operator_result(results []int64, type_id int64) int64 {
	result := int64(0)
	if type_id == 0 {
		for _, val := range results {
			result += val
		}
	} else if type_id == 1 {
		result = 1
		for _, val := range results {
			result = result * val
		}
	} else if type_id == 2 {
		result = MaxInt
		for _, val := range results {
			if val < result {
				result = val
			}
		}
	} else if type_id == 3 {
		for _, val := range results {
			if val > result {
				result = val
			}
		}
	} else if type_id == 5 {
		if results[0] > results[1] {
			result = 1
		}
	} else if type_id == 6 {
		if results[0] < results[1] {
			result = 1
		}
	} else if type_id == 7 {
		if results[0] == results[1] {
			result = 1
		}
	}
	return result
}

func process_packet(packet *string, version_count *int64) int64 {
	result := int64(0)
	operator_results := make([]int64, 0)
	packet_version := convert_to_int((*packet)[:3])
	*packet = (*packet)[3:]
	*version_count += packet_version

	type_id := convert_to_int((*packet)[:3])
	*packet = (*packet)[3:]

	if type_id == 4 {
		result = process_literal_value_packet(packet)
	} else {
		length_type_id := convert_to_int((*packet)[:1])
		*packet = (*packet)[1:]
		if length_type_id == 0 {
			length := int(convert_to_int((*packet)[:15]))
			*packet = (*packet)[15:]
			sub_packet := (*packet)[:length]
			*packet = (*packet)[length:]
			for len(sub_packet) > 0 {
				operator_results = append(operator_results, process_packet(&sub_packet, version_count))
			}
		} else if length_type_id == 1 {
			length := convert_to_int((*packet)[:11])
			*packet = (*packet)[11:]
			for i := int64(0); i < length; i++ {
				operator_results = append(operator_results, process_packet(packet, version_count))
			}
		}
		result = calculate_operator_result(operator_results, type_id)
	}
	return result
}

func process_hex_message(message string) {
	binary_string := decode_hex_to_binary(message)
	total_version := int64(0)
	result := process_packet(&binary_string, &total_version)
	fmt.Println(message)
	fmt.Println("Total version: ", total_version)
	fmt.Println("Calculation result: ", result)
}

func main() {
	main_hex_message := load_file("day_16.txt")
	process_hex_message(main_hex_message)
}
