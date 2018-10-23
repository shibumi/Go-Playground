package main

import "fmt"

// small golang script for the famous FizzBuzz test
// http://wiki.c2.com/?FizzBuzzTest
// Task:
// Write a program that prints the numbers from 1 to 100. But for multiples
// of three print “Fizz” instead of the number and for the multiples of
// five print “Buzz”. For numbers which are multiples of both three and
// five print “FizzBuzz
func main() {
	for i := 1; i < 101; i++ {
		switch {
		case i%3 == 0:
			if i%5 == 0 {
				fmt.Println("FizzBuzz")
			} else {
				fmt.Println("Fizz")
			}
		case i%5 == 0:
			fmt.Println("Buzz")
		default:
			fmt.Println(i)
		}
	}
}
