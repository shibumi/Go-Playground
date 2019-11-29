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
// This has a runtime of O(nm)
func NumberInSlice(number uint64, lists [][]uint64) (result bool) {
	for _, list := range lists {
		ok := false
		for _, e := range list {
			if number == e {
				ok = true
			}
		}
		if !ok {
			return false
		}
	}
	return true
}

// UintSliceToString converts a slice of uints to a string with "," separators.
func UintSliceToString(input []uint64) (output string) {
	for i, e := range input {
		output += strconv.FormatUint(e, 10)
		if i < len(input)-1 {
			output += ", "
		}
	}
	return output
}

// FindIntersection calculates the intersection for slices of numbers.
func FindIntersection(input []string) (output string) {
	lists := Convert(input)
	var outputAsList []uint64
	for _, list := range lists {
		for _, number := range list {
			if NumberInSlice(number, lists) {
				var inSlice bool
				for _, n := range outputAsList {
					if n == number {
						inSlice = true
					}
				}
				if !inSlice {
					outputAsList = append(outputAsList, number)
				}
			}
		}
	}
	return UintSliceToString(outputAsList)
}

func main() {
	fmt.Println(FindIntersection([]string{"1, 2, 3, 4", "1, 3, 4"}))
}
