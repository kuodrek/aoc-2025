package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// var file_path string = "/home/kuodr/Documents/aoc-2025/day-4/test_input.txt"
	var file_path string = "/home/kuodr/Documents/aoc-2025/day-4/input.txt"

	file, _ := os.Open(file_path)

	defer file.Close()
	scanner := bufio.NewScanner(file)

	var rollsOfPaper [][]string
	var acc int = 0

	for scanner.Scan() {
		line := scanner.Text()
		var row []string
		for _, i := range line {

			row = append(row, string(i))
		}
		rollsOfPaper = append(rollsOfPaper, row)
	}

	max_rows := len(rollsOfPaper)
	max_columns := len(rollsOfPaper[0])

	fmt.Println("max rows: ", max_rows)
	fmt.Println("max columns: ", max_columns)

	check_adjacent_rolls := func(i int, j int) bool {
		// get submatrix from rollsOfPaper
		// sum all @'s (subtract one to take into account )
		// if sum >= 4 return false else return true

		i_lower_offset := max(i-1, 0)
		i_upper_offset := min(i+1, max_rows-1)

		j_lower_offset := max(j-1, 0)
		j_upper_offset := min(j+1, max_columns-1)

		submatrix := [][]string{}

		for r := i_lower_offset; r <= i_upper_offset; r++ {
			rowSlice := rollsOfPaper[r][j_lower_offset : j_upper_offset+1]
			submatrix = append(submatrix, rowSlice)
		}

		var count int

		for _, row := range submatrix {
			for _, col := range row {
				if col == "@" {
					count += 1
				}
			}
		}

		count = count - 1

		if count >= 4 {
			return false
		}
		return true
	}

	for i, row := range rollsOfPaper {
		for j := range row {
			if rollsOfPaper[i][j] == "@" {
				if check_adjacent_rolls(i, j) {
					acc += 1
				}
			}
		}
	}
	fmt.Println(acc)
}
