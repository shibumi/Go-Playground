package main

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
