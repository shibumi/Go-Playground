package main

import (
	"fmt"
	"sort"
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

// FindIntersectionPerformance should be smaller and perform better than the other functions in this file. Runtime should be O(n).
func FindIntersectionPerformance(input []string) (output string) {
	dict := make(map[string]uint64)
	for _, s := range input {
		chars := strings.Split(s, ", ")
		for _, char := range chars {
			dict[char]++
		}
	}
	// We need this work around to generate a sorted output string.
	// golang maps are unsorted, therefore we would generate different output strings on every run.
	sortedKeyList := make([]string, len(dict))
	for k := range dict {
		sortedKeyList = append(sortedKeyList, k)
	}
	// We need to sort with this inline function.
	// If we would use sort.Strings() we would get strings such as "1, 13, 4".
	// But we want to sort them by the weight of the number.
	sort.Slice(sortedKeyList, func(i, j int) bool {
		numA, _ := strconv.Atoi(sortedKeyList[i])
		numB, _ := strconv.Atoi(sortedKeyList[j])
		return numA < numB
	})
	for _, k := range sortedKeyList {
		if dict[k] == uint64(len(input)) {
			output += k + ", "
		}
	}
	// We need to trim the last comma and whitespace from the string.
	output = strings.TrimSuffix(output, ", ")
	return output
}

// The challenge is the following:
// We get as input a slice of strings with sorted numbers.
// We want as output a string with the intersection of these slices.
// The major challenge is the intersection algorithm.
// The side quest is the input/output handling, as we get strings and not
// integers (what would make this challenge a lot easier).
func main() {
	fmt.Println(FindIntersectionPerformance([]string{"1, 2, 3, 4", "1, 3, 4"}))
}
