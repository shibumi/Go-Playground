package main

import (
	"reflect"
	"testing"
)

func TestUintSliceToString(t *testing.T) {
	correctTables := []struct {
		input  []uint64
		result string
	}{
		{[]uint64{1, 2, 3, 4}, "1, 2, 3, 4"},
	}
	for _, table := range correctTables {
		if result := UintSliceToString(table.input); result != table.result {
			t.Errorf("Incorrect result for input: %v, got: %v, want: %v", table.input, result, table.result)
		}
	}
}

func TestFindIntersection(t *testing.T) {
	correctTables := []struct {
		input  []string
		result string
	}{
		{[]string{"1, 2, 3, 4", "1, 2, 3, 4, 5"}, "1, 2, 3, 4"},
		{[]string{"1, 3, 4, 7, 13", "1, 2, 4, 13, 15"}, "1, 4, 13"},
		{[]string{"1, 2, 6", "1, 5, 7", "1, 4, 8"}, "1"},
	}
	for _, table := range correctTables {
		result := FindIntersection(table.input)
		if result != table.result {
			t.Errorf("Incorrect result for input: %v, got: %v, want: %v", table.input, result, table.result)
		}
	}
}

func TestNumberInSlice(t *testing.T) {
	correctTables := []struct {
		number uint64
		lists  [][]uint64
		result bool
	}{
		{1, [][]uint64{{1, 2, 3, 4}, {1, 5, 6, 7, 8}}, true},
	}
	for _, table := range correctTables {
		result := NumberInSlice(table.number, table.lists)
		if result != table.result {
			t.Errorf("Incorrect result for number: %v in lists %v, got: %v, want: %v", table.number, table.lists, result, table.result)
		}
	}
}

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
