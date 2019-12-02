package main

import (
	"reflect"
	"testing"
)

func TestT9Permutator(t *testing.T) {
	correctTables := []struct {
		numberblock []int
		results     []string
	}{
		{[]int{2, 2}, []string{"AA", "BB", "CC", "AB", "BA", "AC", "CA", "BC", "CB"}},
	}
	for _, value := range correctTables {
		if results := T9Permutator(value.numberblock); reflect.DeepEqual(results, value.results) {
			t.Errorf("result for %v is incorecct, got: %v, want: %v", value.numberblock, results, value.results)
		}
	}
}
