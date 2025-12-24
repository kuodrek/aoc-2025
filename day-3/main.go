package main

import (
	"bufio"
	"fmt"
	"os"
)

func get_highest_joltage(bank []int) int {
	var max_value int = 0

	for i, j := range bank {
		for _, j2 := range bank[i+1:] {
			joltage := j*10 + j2
			if joltage > max_value {
				max_value = joltage
			}
		}
	}
	return max_value
}

func main() {
	// var file_path string = "day-3/test_input.txt"
	var file_path string = "day-3/input.txt"

	file, _ := os.Open(file_path)

	defer file.Close()
	scanner := bufio.NewScanner(file)

	var banks [][]int
	var acc int = 0

	for scanner.Scan() {
		line := scanner.Text()
		var joltages []int
		for _, i := range line {

			joltages = append(joltages, int(i-'0'))
		}
		banks = append(banks, joltages)
	}

	for _, bank := range banks {
		acc += get_highest_joltage(bank)
	}
	fmt.Println((banks))
	fmt.Println(acc)
}
