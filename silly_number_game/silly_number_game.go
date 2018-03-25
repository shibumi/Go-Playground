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

func generateRandomNumber() (randNum int64) {
	return int64(rand.Intn(2001) + 1)
}

func main() {
    var endGame bool = false
    var counter int = 0
	randNum := generateRandomNumber()
	fmt.Println("Welcome to my silly number game")
	fmt.Println("Press 0 to exit the game")
	reader := bufio.NewReader(os.Stdin)
	for !endGame {
		fmt.Print("Please enter a number: ")
		text, err := reader.ReadString('\n')
		text = strings.Trim(text, "\n")
		if err != nil {
			log.Fatal(err)
			continue
		}
		i, err := strconv.ParseInt(text, 10, 32)
		if err != nil {
			log.Fatal(err)
			continue
		}
		switch {
		case i == randNum:
			fmt.Println("Boom! You've guessed the right number: ", randNum)
            endGame = true
		case (i > randNum) && (i != 0):
			fmt.Println("Too high")
		case (i < randNum) && (i != 0):
			fmt.Println("Too low")
		case i == 0:
			fmt.Println("You have pressed 0. Exiting game...")
            endGame = true
		default:
			fmt.Println("Too bad. Wrong number.")
		}
        counter++
	}
	fmt.Println("Thanks for playing!")
    fmt.Println("You've needed ", counter, " tries")
}
