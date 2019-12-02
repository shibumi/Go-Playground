package main

import "testing"

func TestFindMissingNumber(t *testing.T) {
	correctTables := []struct {
		completeSlice   []int
		incompleteSlice []int
		result          int
	}{
		{[]int{5, 8, 9, 12, 24, 9}, []int{5, 8, 9, 24, 9}, 12},
		{[]int{0, 1}, []int{0}, 1},
		{[]int{-5, -23, -42, -99}, []int{-5, -42, -99}, -23},
	}
	for _, value := range correctTables {
		if result, err := findMissingNumber(value.completeSlice, value.incompleteSlice); result != value.result && err == nil {
			t.Errorf("result for %v and %v is incorrect, got: %d, want: %d", value.completeSlice, value.incompleteSlice, result, value.result)
		} else if err != nil {
			t.Errorf("error in findMissingNumber: %s", err)
		}
	}
}
