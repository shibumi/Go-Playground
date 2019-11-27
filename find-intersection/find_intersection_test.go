package main

import (
	"reflect"
	"testing"
)

// func TestFindIntersection(t *testing.T) {
// 	correctTables := []struct {
// 		input  []string
// 		result string
// 	}{
// 		{[]string{"1, 2, 3, 4", "1, 2, 3, 4, 5"}, "1, 2, 3, 4"},
// 		{[]string{"1, 3, 4, 7, 13", "1, 2, 4, 13, 15"}, "1, 4, 13"},
// 		{[]string{"1, 2, 6", "1, 5, 7", "1, 4, 8"}, "1"},
// 	}
// }

func TestConvert(t *testing.T) {
	correctTables := []struct {
		input  []string
		result [][]uint64
	}{
		{[]string{"1, 2, 3, 4", "5, 6, 7, 8"}, [][]uint64{{1, 2, 3, 4}, {5, 6, 7, 8}}},
	}
	for _, table := range correctTables {
		result := Convert(table.input)
		if !reflect.DeepEqual(result, table.result) {
			t.Errorf("Conversion for: input array %s was incorrect, got: %v, want: %v", table.input, result, table.result)
		}
	}
}
