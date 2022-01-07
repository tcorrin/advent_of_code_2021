package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

type beacon struct {
	x int
	y int
	z int
}

type scanner struct {
	name        string
	beacon_list []beacon
}

func compare_offsets(a_list []int, b_list []int) bool {
	match_count := 0
	for _, a := range a_list {
		for _, b := range b_list {
			if a == b {
				match_count++
			}
		}
	}
	return match_count > 5
}

func load_file(filename string) []scanner {
	scanner_list := make([]scanner, 0)
	file, err := os.Open(filename)
	if err == nil {
		file_scanner := bufio.NewScanner(file)
		var s scanner
		for file_scanner.Scan() {
			line := file_scanner.Text()
			if strings.Contains(line, "---") {
				if s.name != "" {
					scanner_list = append(scanner_list, s)
				}
				s = scanner{beacon_list: make([]beacon, 0), name: line}
			} else if strings.Contains(line, ",") {
				var x int
				var y int
				var z int
				fmt.Sscanf(line, "%d,%d,%d", &x, &y, &z)
				s.beacon_list = append(s.beacon_list, beacon{x: x, y: y, z: z})
			}
		}
	} else {
		fmt.Println("Error opening file")
		fmt.Println(err)
		os.Exit(23)
	}
	return scanner_list
}

func match_scanners(scanner_a scanner, scanner_b scanner) (bool, []beacon, []beacon) {
	matching_beacons_a := make([]beacon, 0)
	matching_beacons_b := make([]beacon, 0)
	a_offset_map := make(map[beacon][]int)
	b_offset_map := make(map[beacon][]int)
	for _, i := range scanner_a.beacon_list {
		a_offset_map[i] = make([]int, 0)
		for _, j := range scanner_a.beacon_list {
			if i != j {
				md := int(math.Abs(float64(j.x-i.x)) + math.Abs(float64(j.y-i.y)) + math.Abs(float64(j.z-i.z)))
				a_offset_map[i] = append(a_offset_map[i], md)
			}
		}
	}

	for _, i := range scanner_b.beacon_list {
		b_offset_map[i] = make([]int, 0)
		for _, j := range scanner_b.beacon_list {
			if i != j {
				md := int(math.Abs(float64(j.x-i.x)) + math.Abs(float64(j.y-i.y)) + math.Abs(float64(j.z-i.z)))
				b_offset_map[i] = append(b_offset_map[i], md)
			}
		}
	}

	for a, mn_list_a := range a_offset_map {
		for b, mn_list_b := range b_offset_map {
			if compare_offsets(mn_list_a, mn_list_b) {
				matching_beacons_a = append(matching_beacons_a, a)
				matching_beacons_b = append(matching_beacons_b, b)
			}
		}
	}
	return len(matching_beacons_a) >= 12, matching_beacons_a, matching_beacons_b
}

func process_file(filename string) {
	scanner_list := load_file(filename)
	for _, scanner_a := range scanner_list {
		for _, scanner_b := range scanner_list {
			if scanner_a.name != scanner_b.name {
				match, scanner_a_matching_beacons, scanner_b_matching_beacons := match_scanners(scanner_a, scanner_b)
				if match {
					println(scanner_a.name, "matches", scanner_b.name)
					println(scanner_a_matching_beacons)
					println(scanner_b_matching_beacons)
				}
			}
		}
	}
}

func main() {
	process_file("day_19_test.txt")
	//process_file("day_19.txt")
}
