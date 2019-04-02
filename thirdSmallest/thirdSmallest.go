package main

import "fmt"

// My implemention for https://www.youtube.com/watch?v=ztyCQjiAAdU
func thirdSmallest(candidates []int) int {
	var comparingSet []int
	if len(candidates) < 3 {
		return 0
	}
	for _, numb := range candidates {
		fmt.Println("Number ", numb)
		switch {
		case len(comparingSet) == 1:
			if numb < comparingSet[0] {
				comparingSet[0] = numb
			} else {
				comparingSet = append(comparingSet, numb)
			}
		case len(comparingSet) == 2:
			if numb < comparingSet[1] && numb > comparingSet[0] {
				comparingSet[1] = numb
			} else if numb < comparingSet[0] {
				comparingSet[1] = comparingSet[0]
				comparingSet[0] = numb
			} else {
				comparingSet = append(comparingSet, numb)
			}
		case len(comparingSet) == 3:
			switch {
			case numb < comparingSet[0]:
				comparingSet[2] = comparingSet[1]
				comparingSet[1] = comparingSet[0]
				comparingSet[0] = numb
			case numb < comparingSet[1]:
				comparingSet[2] = comparingSet[1]
				comparingSet[1] = numb
			case numb < comparingSet[2]:
				comparingSet[2] = numb
			}
		default:
			comparingSet = append(comparingSet, numb)
		}
	}
	fmt.Println(comparingSet)
	return comparingSet[2]
}

func main() {
	candidates := []int{1, 6, 7, 4, 7, 15, 22, 0, 54}
	fmt.Println(thirdSmallest(candidates))
}
