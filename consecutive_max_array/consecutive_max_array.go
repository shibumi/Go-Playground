package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func calcDifference(x uint, y uint) uint {
	if x >= y {
		return x - y
	}
	return y - x
}

func algorithmicArray(input []uint) (solution []uint) {
	hashMap := map[uint][]uint{}
	var difference uint
	difference = 0
	hashMap[difference] = []uint{}
	for i := 0; i < len(input); i++ {
		result := uint(0)
		if i == len(input)-1 {
			result = calcDifference(input[i], input[i-1])
		} else {
			result = calcDifference(input[i], input[i+1])
		}
		hashMap[result] = append(hashMap[result], input[i])
	}
	var result uint
	result = 0
	var key uint
	for k, v := range hashMap {
		if result < uint(len(v)) {
			result = uint(len(v))
			key = k
		}
	}
	return hashMap[key]
}

func main() {
	//algorithmicArray([]uint{10, 7, 4, 6, 8, 10, 11})
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
