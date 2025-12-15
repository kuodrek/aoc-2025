package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func is_invalid_id(id string) bool {
	// todo: write function
	return true
}

func main() {
	// file, _ := os.Open("day-2/input.txt")
	file, _ := os.Open("day-2/test_input.txt")

	defer file.Close()
	scanner := bufio.NewScanner(file)

	var ranges []string
	var counter int = 0

	for scanner.Scan() {
		line := scanner.Text()
		ranges = strings.Split(line, ",")
	}

	fmt.Println(ranges)

	for _, r := range ranges {
		fmt.Println(r)
		// for n in strings.Split(r, "-")
		// check if char appears again in remaining of string
		// if true, counter += int(n)
	}

	fmt.Println(counter)
}
