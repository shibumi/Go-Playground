package check_permutation

import "testing"

func TestCheckPermutation(t *testing.T) {
	tables := []struct {
		str1   string
		str2   string
		result bool
	}{
		{str1: "abcde", str2: "acdeb", result: true},
		{str1: "abcdef", str2: "abdef", result: false},
		{str1: "", str2: "", result: true},
	}
	for _, table := range tables {
		result := checkPermutation(table.str1, table.str2)
		if result != table.result {
			t.Errorf("Error, parameters: %s, %s, result %t, should have: %t", table.str1, table.str2, result, table.result)
		}
	}
}
