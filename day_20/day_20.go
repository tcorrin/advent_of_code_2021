package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func load_file(filename string) ([]string, [][]string) {
	image_algo := make([]string, 0)
	image := make([][]string, 0)
	file, err := os.Open(filename)
	if err == nil {
		file_scanner := bufio.NewScanner(file)
		for file_scanner.Scan() {
			line := file_scanner.Text()
			if line != "" {
				if len(image_algo) == 0 {
					image_algo = strings.Split(line, "")
				} else {
					image = append(image, strings.Split(line, ""))
				}
			}
		}
	} else {
		fmt.Println("Error opening file")
		fmt.Println(err)
		os.Exit(23)
	}
	return image_algo, image
}

func convert_to_int(binary_string string) int {
	b := fmt.Sprintf("%064s", binary_string)
	i, _ := strconv.ParseInt(b, 2, 64)
	return int(i)
}

func make_border(w int) []string {
	b := make([]string, w)
	for i := 0; i < w; i++ {
		b[i] = "."
	}
	return b
}

func add_border(image [][]string, size int) [][]string {
	new_image := make([][]string, 0)
	new_width := len(image[0]) + (size * 2)
	for i := 0; i < size; i++ {
		new_image = append(new_image, make_border(new_width))
	}

	for _, line := range image {
		new_line := append(make_border(size), line...)
		new_image = append(new_image, append(new_line, make_border(size)...))
	}

	for i := 0; i < size; i++ {
		new_image = append(new_image, make_border(new_width))
	}
	return new_image
}

func algo_lookup(image_algo []string, image [][]string, x int, y int) string {
	row_one := make([]string, 3)
	row_two := make([]string, 3)
	row_three := make([]string, 3)
	copy(row_one, image[y-1][x-1:x+2])
	copy(row_two, image[y][x-1:x+2])
	copy(row_three, image[y+1][x-1:x+2])
	lookup := append(row_one, row_two...)
	lookup = append(lookup, row_three...)

	i := convert_lookup_to_int(lookup)
	return image_algo[i]
}

func convert_lookup_to_int(lookup []string) int {
	bin_num_slice := make([]string, len(lookup))
	for i := 0; i < len(lookup); i++ {
		if lookup[i] == "." {
			bin_num_slice[i] = "0"
		} else if lookup[i] == "#" {
			bin_num_slice[i] = "1"
		}
	}
	result := convert_to_int(strings.Join(bin_num_slice, ""))
	return result
}

func enhance_image(image_algo []string, image [][]string) [][]string {
	height := len(image)
	width := len(image[0])
	new_image := make([][]string, height)
	for i := 0; i < height; i++ {
		new_image[i] = make([]string, width)
	}

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if x > 0 && y > 0 && x < width-1 && y < height-1 {
				new_image[y][x] = algo_lookup(image_algo, image, x, y)
			} else {
				new_image[y][x] = "."
			}
		}
	}

	return new_image
}

func count_lit_pixels(image [][]string) int {
	count := 0
	for _, line := range image {
		for _, pixel := range line {
			if pixel == "#" {
				count++
			}
		}
	}
	return count
}

func trim_image(image [][]string, size int) [][]string {
	height := len(image)
	width := len(image[0])
	new_height := height - (size * 2)
	new_image := make([][]string, new_height)

	for i := size; i < height-size; i++ {
		new_image[i-size] = make([]string, 0)
		new_image[i-size] = append(new_image[i-size], image[i][size:width-size]...)
	}
	return new_image
}

func process_file(filename string, limit int) {
	image_algo, image := load_file(filename)
	image = add_border(image, limit*2)
	for i := 0; i < limit; i++ {
		image = enhance_image(image_algo, image)
		image = trim_image(image, 1)
	}
	fmt.Println("List Pixels:", count_lit_pixels(image))
}

func main() {
	process_file("day_20_test.txt", 2)
	process_file("day_20.txt", 2)
	process_file("day_20_test.txt", 50)
	process_file("day_20.txt", 50)
}
