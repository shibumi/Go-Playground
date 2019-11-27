package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Convert converts a slice of strings like {"1, 2, 3, 4"} into a slice of slices of uint64
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

func findIntersection(input []string) (output string) {
	fmt.Println(Convert(input))
	return "test"
}

func main() {
	fmt.Println(findIntersection([]string{"1, 2, 3, 4", "1, 3, 4"}))
}
