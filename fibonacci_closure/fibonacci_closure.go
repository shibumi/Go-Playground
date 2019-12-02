package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
// See go-tour: https://go-tour-de.appspot.com/moretypes/25
func fibonacci() func() int {
	first, second := 0, 1
	return func() int {
		ret := first
		first, second = second, first+second
		return ret
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
