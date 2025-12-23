package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func splitInt(n int) []int {
	slc := []int{}
	for n > 0 {
		slc = append(slc, n%10)
		n /= 10
	}

	result := []int{}
	for i := range slc {
		result = append(result, slc[len(slc)-1-i])
	}

	return result
}

func join_digits(digits []int) string {
	// fmt.Println("Digits to be joined:", digits)
	// var acc int = 0
	// slices.Reverse(digits)
	// for i, d := range digits {
	// 	acc += d * int(math.Pow10(i))
	// }
	// return acc
	var digits_str []string
	for _, d := range digits {
		digits_str = append(digits_str, strconv.Itoa(d))
	}
	return strings.Join(digits_str, "")
}

func containsDuplicate(digits []int) bool {
	for i := range len(digits) / 2 {
		left_str := join_digits(digits[:i+1])
		right_str := join_digits(digits[i+1:])

		if left_str == right_str {
			return true
		}
	}
	return false
}

func sum_invalid_ids(id_range string) int {
	id_interval := strings.Split(id_range, "-") // 11-22 -> [11, 22]
	start, _ := strconv.Atoi(id_interval[0])
	end, _ := strconv.Atoi(id_interval[1])

	var acc int = 0
	var invalid_ids []int

	for i := start; i <= end; i++ {
		digits := splitInt(i)
		has_duplicates := containsDuplicate(digits)
		if has_duplicates && !slices.Contains(invalid_ids, i) {
			invalid_ids = append(invalid_ids, i)
			acc += i
		}
	}
	// fmt.Println(invalid_ids)
	return acc
}

func main() {
	// var file_path string = "day-2/test_input.txt"
	var file_path string = "day-2/input.txt"

	file, _ := os.Open(file_path)

	defer file.Close()
	scanner := bufio.NewScanner(file)

	var ranges []string
	var acc int = 0

	for scanner.Scan() {
		line := scanner.Text()
		r := strings.Split(line, ",")
		ranges = append(ranges, r...) // Input is only one line but Im prepending the lines just in case
	}

	// fmt.Println(ranges)

	for _, r := range ranges {
		invalid_id_sum := sum_invalid_ids(r)
		acc += invalid_id_sum
	}

	fmt.Println(acc)
}
