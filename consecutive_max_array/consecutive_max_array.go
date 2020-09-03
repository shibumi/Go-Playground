package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func algorithmicArray(input []uint) (solution []uint) {
	if len(input) < 2 {
		return []uint{}
	}
	var last uint
	currentChain := []uint{}
	for i := 0; i < len(input)-1; i++ {
		result := input[i] - input[i+1]
		if len(currentChain) == 0 {
			currentChain = append(currentChain, input[i], input[i+1])
			last = result
		} else {
			if result != last {
				currentChain = []uint{}
				currentChain = append(currentChain, input[i], input[i+1])
				last = result
				continue
			}
			currentChain = append(currentChain, input[i+1])
		}
		if len(currentChain) > len(solution) {
			solution = currentChain
		}
	}
	return
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	// just scan first line to get number of tests
	scanner.Scan()
	i := 1
	caseNr := 1
	for scanner.Scan() {
		if i%2 != 0 {
			i++
			continue
		}
		line := strings.Split(scanner.Text(), " ")
		array := []uint{}
		for _, v := range line {
			n, err := strconv.ParseUint(v, 10, 64)
			if err != nil {
				log.Fatalln(err)
			}
			i := uint(n)
			array = append(array, i)
		}
		fmt.Printf("Case #%d: %d\n", caseNr, len(algorithmicArray(array)))
		caseNr++
		i++
	}
}
