package main

import (
	"reflect"
	"testing"
)

func TestConsecutiveMaxArray(t *testing.T) {
	tables := []struct {
		input  []uint
		output []uint
	}{{
		input:  []uint{10, 7, 4, 6, 8, 10, 11},
		output: []uint{4, 6, 8, 10},
	},
		{
			input:  []uint{9, 7, 5, 3},
			output: []uint{9, 7, 5, 3},
		},
		{
			input:  []uint{5, 5, 4, 5, 5, 5, 4, 5, 6},
			output: []uint{5, 5, 5},
		},
		{
			input:  []uint{5, 4, 3, 2, 1, 2, 3, 4, 5, 6},
			output: []uint{1, 2, 3, 4, 5, 6},
		},
	}
	for _, table := range tables {
		result := algorithmicArray(table.input)
		if !reflect.DeepEqual(table.output, result) {
			t.Errorf("algorithmicArray calculated wrong result for %v. Should have: %v, got: %v", table.input, table.output, result)
		}
	}
}
