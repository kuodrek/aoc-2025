package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("day-1/input.txt")
	// file, err := os.Open("day-1/test_input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var counter, password int = 50, 0

	rotationMap := make(map[string]int)
	rotationMap["R"] = 1
	rotationMap["L"] = -1

	for scanner.Scan() {
		line := scanner.Text()

		rotation := string(line[0])
		value, _ := strconv.Atoi(line[1:])
		counter += rotationMap[rotation] * value

		if counter%100 == 0 {
			password += 1
		}
		fmt.Println(counter)
	}
	fmt.Println(counter)
	fmt.Println(password)
}
