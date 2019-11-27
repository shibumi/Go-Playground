package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Convert converts a slice of strings like {"1, 2, 3, 4"} into a slice of slices of uint64
// Has O(n^2), this can be very possibly achieved much faster...
func Convert(input []string) (output [][]uint64) {
	for _, element := range input {
		numbers := strings.Split(element, ",")
		var list []uint64
		for _, number := range numbers {
			u, err := strconv.ParseUint(strings.TrimSpace(number), 10, 64)
			if err != nil {
				fmt.Println(err)
			}
			list = append(list, u)
		}
		output = append(output, list)
	}
	return output
}

// NumberInSlice checks if a number exists in all Slices in a Slice
func NumberInSlice(number uint64, lists [][]uint64) (result bool) {
	var counter uint
	var counterSlices uint
	for _, list := range lists {
		for _, e := range list {
			if number == e {
				counter++
			}
		}
		counterSlices++
	}
	if counter != counterSlices {
		return false
	}
	return true
}

func findIntersection(input []string) (output string) {
	return "test"
}

func main() {
	fmt.Println(findIntersection([]string{"1, 2, 3, 4", "1, 3, 4"}))
}
