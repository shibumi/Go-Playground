package main

import "testing"

func TestThirdSmallest(t *testing.T) {
	tables := []struct {
		candidates []int
		result     int
	}{
		{[]int{1, 6, 7, 4, 7, 15, 22, 0, 54}, 4},
		{[]int{0, 5, 18, 88, 219, 61, 44, 77, 21, 5, 2, 1}, 2},
		{[]int{-400, -1, 20, 24, 54, 0, -5, 14, 22, 900, 1}, -1},
		{[]int{1, 2, 3, 3, 3, 3, 3}, 3},
		{[]int{0}, 0},
	}
	for _, table := range tables {
		res := thirdSmallest(table.candidates)
		if res != table.result {
			t.Errorf("Result for %v is incorrect, got: %d, want: %d", table.candidates, res, table.result)
		}
	}
}
