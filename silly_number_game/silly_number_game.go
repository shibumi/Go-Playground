package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Generate random number
	randNum := int64(rand.Intn(2001) + 1)
	fmt.Println("Welcome to my silly number game")
	fmt.Println("Press 0 to exit the game")
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Please enter a number: ")
		text, err := reader.ReadString('\n')
		text = strings.Trim(text, "\n")
		if err != nil {
			log.Fatal(err)
		}
		i, err := strconv.ParseInt(text, 10, 32)
		if err != nil {
			log.Fatal(err)
		}
		if i == randNum {
			fmt.Println("BOOM! You've guessed the right number: ", randNum)
			break
		} else if i == 0 {
			fmt.Println("You have pressed 0. Exiting game...")
			break
		} else {
			fmt.Println("Too bad. Wrong number.")
		}
	}
	fmt.Println("Thanks for playing!")

}
