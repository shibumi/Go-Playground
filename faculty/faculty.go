package main

import "log"

// This is a simple faculty implementation in Golang
// It's runtime is O(n)
// I am pretty sure this can be done recursive as well,
// but golang has no TCO (tail call optimization).
// So an iterative way should be preferred.
func faculty(number uint) (result uint) {
	if number == 1 {
		return 1
	}
	for result = 1; number >= 1; number-- {
		result = result * number
	}
	return result
}

func main() {
	log.Println(faculty(9))
}
