package main

import (
	"errors"
	"sort"
)

// short function to check if two int slices are equal
// "equal" returns an error if:
//  one of the slices is empty
func equal(x []int, y []int) (result bool, err error) {
	if x == nil || y == nil {
		return false, errors.New("one of the slices is empty")
	}
	if len(x) != len(y) {
		return false, nil
	}
	sort.Ints(x)
	sort.Ints(y)
	for i := range x {
		if x[i] != y[i] {
			return false, nil
		}
	}
	return true, nil
}

// findMissingNumber will check two int slices on equality,
// if they are not equal, findMissingNumber will add the values from the slices to sums
// then it will substract these sums to get the missing number
// Note: this works for one missing number and slices of integers only.
// The time complexity is O(n), because we go through the slices only once
func findMissingNumber(completeSlice []int, incompleteSlice []int) (missingNumber int, err error) {
	if result, err := equal(completeSlice, incompleteSlice); result && err != nil {
		return 0, err
	} else if result && err == nil {
		return 0, errors.New("both slices are equal")
	}
	sumCompleteSlice := 0
	for _, value := range completeSlice {
		sumCompleteSlice = value + sumCompleteSlice
	}
	sumIncompleteSlice := 0
	for _, value := range incompleteSlice {
		sumIncompleteSlice = value + sumIncompleteSlice
	}
	result := sumCompleteSlice - sumIncompleteSlice
	return result, nil
}

func main() {

}
