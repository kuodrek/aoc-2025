package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type IdRange struct {
	LowerBound int
	UpperBound int
}

func (r IdRange) checkIdValid(id int) bool {
	return id >= r.LowerBound && id <= r.UpperBound
}

func (r IdRange) ArithmeticSum() int {
	return (r.UpperBound - r.LowerBound) * (r.UpperBound + r.LowerBound) / 2
}

func main() {
	// var file_path string = "day-5/test_input.txt"
	var file_path string = "day-5/input.txt"

	fmt.Println("hello world")

	file, _ := os.Open(file_path)

	defer file.Close()
	scanner := bufio.NewScanner(file)

	var AvailableIdRanges []IdRange
	var IdsToTest []int
	var counter int
	var validIds []int

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" { // skip blank line between id range and ids to test
			continue
		}

		if strings.Contains(line, "-") {
			ids_str := strings.Split(line, "-")
			lower_bound, _ := strconv.Atoi(ids_str[0])
			upper_bound, _ := strconv.Atoi(ids_str[1])
			AvailableIdRanges = append(AvailableIdRanges, IdRange{lower_bound, upper_bound})
		} else {
			id, _ := strconv.Atoi(line)
			IdsToTest = append(IdsToTest, id)
		}

	}

	// part one
	for _, id := range IdsToTest {
		for _, idRange := range AvailableIdRanges {
			if idRange.checkIdValid(id) && !slices.Contains(validIds, id) {
				// fmt.Println("Id ", id, "is between ", idRange)
				validIds = append(validIds, id)
				counter += 1
			}
		}
	}

	fmt.Println(counter)
}
