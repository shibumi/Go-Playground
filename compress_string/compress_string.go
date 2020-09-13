package compress

import "strconv"

// Time and Space complexity is O(n).
func compressString(input string) (output string) {
	if len(input) == 0 {
		return ""
	}
	counter := 0
	var lastChar rune
	for i, char := range input {
		if i == 0 {
			lastChar = char
			counter++
			output += string(char)
			continue
		}
		if lastChar == char {
			counter++
			continue
		} else {
			n := strconv.Itoa(counter)
			if counter != 1 {
				output += n
			}
			lastChar = char
			output += string(char)
			counter = 1
		}
	}
	if counter != 1 {
		output += strconv.Itoa(counter)
	}
	if len(input) <= len(output) {
		return input
	}
	return
}
