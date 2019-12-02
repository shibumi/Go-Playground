package main

import "fmt"

var T9 = map[int]string{
	1: "_",
	2: "ABC",
	3: "DEF",
	4: "GHI",
	5: "JKL",
	6: "MNO",
	7: "PQRS",
	8: "TUV",
	9: "WXYZ",
	0: "+"}

func T9Permutator(numberblock []int) (results []string) {
	for index, value := range numberblock {
		for _, charvalue := range T9[value] {
			if index == 0 {
				results = append(results, string(charvalue))
			} else {
				results[index] += string(charvalue)
			}

		}
	}
	fmt.Println(results)
	return results
}

func main() {
	T9Permutator([]int{2, 2})
}
