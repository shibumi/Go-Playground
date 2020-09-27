package check_permutation

import "reflect"

// The problem: We want to check if two strings
// are permutations from each other.
// Time complexity O(n) or better: O(2n)
// Space complexity  O(n)
// This solution is not so good, because it repeats code...
// We can do this better...
func checkPermutation(str1, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}
	countMap1 := map[rune]uint{}
	countMap2 := map[rune]uint{}
	for _, char := range str1 {
		countMap1[char]++
	}
	for _, char := range str2 {
		countMap2[char]++
	}
	if reflect.DeepEqual(countMap1, countMap2) {
		return true
	}
	return false
}
