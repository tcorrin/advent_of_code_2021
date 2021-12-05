package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type bingo_number struct {
	value int
	mark  bool
}

func process_number_list(input string) []int {
	number_sequence := make([]int, 0)
	split_list := strings.Split(input, ",")
	for _, number := range split_list {
		number_int, _ := strconv.Atoi(number)
		number_sequence = append(number_sequence, number_int)
	}
	return number_sequence
}

func process_bingo_board(input []string, number_sequence []int) [][]bingo_number {
	board := make([][]bingo_number, 0)
	for _, row := range input {
		board_row := make([]bingo_number, 0)
		number_list := strings.Fields(row)
		for _, number := range number_list {
			number_int, _ := strconv.Atoi(number)
			board_row = append(board_row, bingo_number{value: number_int, mark: false})
		}
		board = append(board, board_row)
	}
	return board
}

func load_file(filename string) ([]int, [][][]bingo_number) {

	file, err := os.Open(filename)
	board_list := make([][][]bingo_number, 0)
	var number_sequence []int
	bingo_board_raw := make([]string, 0)

	if err == nil {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			if strings.Contains(line, ",") {
				number_sequence = process_number_list(line)
			} else if line == "" {
				if len(bingo_board_raw) > 0 {
					bingo_board := process_bingo_board(bingo_board_raw, number_sequence)
					board_list = append(board_list, bingo_board)
				}
				bingo_board_raw = make([]string, 0)
			} else {
				bingo_board_raw = append(bingo_board_raw, line)
			}
		}
	} else {
		fmt.Println("Error opening file")
		fmt.Println(err)
	}

	return number_sequence, board_list
}

func mark_numbers(board *[][]bingo_number, number int) {
	b := *board
	for i, row := range *board {
		for j := range row {
			if b[i][j].value == number {
				b[i][j].mark = true
			}
		}
	}
}

func transpose_bingo_board(bingo_board [][]bingo_number) [][]bingo_number {
	new_board := make([][]bingo_number, 0)
	for i := 0; i < 5; i++ {
		row := make([]bingo_number, 5)
		new_board = append(new_board, row)
	}
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			new_board[i][j] = bingo_board[j][i]
		}
	}
	return new_board
}

func check_for_victory(board [][]bingo_number) bool {
	victory := false
	for _, row := range board {
		for _, number := range row {
			victory = number.mark
			if !victory {
				break
			}
		}
		if victory {
			return true
		}
	}
	return false
}

func print_result(board [][]bingo_number, number int) {
	unmarked_total := 0
	for _, row := range board {
		for _, bn := range row {
			if !bn.mark {
				unmarked_total += bn.value
			}
		}
	}
	fmt.Println("unmarked_total: ", unmarked_total,
		" winning_number: ",
		number,
		" result: ",
		unmarked_total*number)
}

func process_file(filename string) {
	number_sequence, board_list := load_file(filename)
	for _, number := range number_sequence {

		// mark numbers
		for _, board := range board_list {
			mark_numbers(&board, number)
		}

		// check for victory
		for _, board := range board_list {
			if check_for_victory(board) || check_for_victory(transpose_bingo_board(board)) {
				print_result(board, number)
				return
			}
		}
	}
}

func remove_boards(board_list [][][]bingo_number, boards_to_remove []int) [][][]bingo_number {
	new_board_list := make([][][]bingo_number, 0)
	for i, value := range board_list {
		skip := false
		for _, index := range boards_to_remove {
			if i == index {
				skip = true
				break
			}
		}
		if !skip {
			new_board_list = append(new_board_list, value)
		}
	}
	return new_board_list
}

func process_file_pt2(filename string) {
	number_sequence, board_list := load_file(filename)
	for _, number := range number_sequence {
		boards_to_remove := make([]int, 0)

		// mark numbers
		for _, board := range board_list {
			mark_numbers(&board, number)
			//print_bingo_board(board)
		}

		// check for victory
		for i, board := range board_list {
			if check_for_victory(board) || check_for_victory(transpose_bingo_board(board)) {
				if len(board_list) == 1 {
					print_result(board, number)
					return
				} else {
					boards_to_remove = append(boards_to_remove, i)
				}
			}
		}

		board_list = remove_boards(board_list, boards_to_remove)
	}
}

func main() {
	process_file("day_04_test.txt")
	process_file("day_04.txt")
	process_file_pt2("day_04_test.txt")
	process_file_pt2("day_04.txt")
}
