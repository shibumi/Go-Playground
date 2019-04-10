package main

import (
	"log"
)

// Standard partition process of QuickSort().
// It considers the last element as pivot
// and moves all smaller elements to the left of the pivot.
// Greater elements will be placed right of the pivot.
func partition(candidates []int, left int, right int) (index int) {
	index = left
	for j := left; j <= right-1; j++ {
		if candidates[j] <= candidates[right] {
			candidates[index], candidates[j] = candidates[j], candidates[index]
			index++
		}
	}
	candidates[index], candidates[right] = candidates[right], candidates[index]
	//log.Println(candidates)
	//log.Println(index)
	return index
}

func kthSmallest(candidates []int, left int, right int, k int) int {
	// If k is smaller than number of
	// elements in array
	if k > 0 && k <= right-left+1 {

		// Partition the array around last
		// element and get position of pivot
		// element in sorted array
		index := partition(candidates, left, right)

		// If position is same as k
		if index-left == k-1 {
			return candidates[index]
		}

		// If position is more, recur
		// for left subarray
		if index-left > k-1 {
			return kthSmallest(candidates, left, index-1, k)
		}

		// Else recur for right subarray
		return kthSmallest(candidates, index+1, right, k-index+left-1)
	}

	// If k is more than number of
	// elements in array
	return 0
}

// This is a quickselect implementation.
// See also: https://www.geeksforgeeks.org/quickselect-algorithm/
func main() {
	candidates := []int{1, 6, 77, 4, 7, 15, 22, 0, 54}
	//partition(candidates, 0, 8)
	log.Println(kthSmallest(candidates, 0, 8, 3))
}
