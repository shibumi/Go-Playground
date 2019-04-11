package main

import "testing"

func TestFaculty(t *testing.T) {
	tables := map[uint]uint{
		6: 720,
		2: 2,
		1: 1,
		9: 362880,
	}
	for k, v := range tables {
		result := faculty(k)
		if v != result {
			t.Errorf("Faculty %d! was incorrect, got: %d, want: %d", k, result, v)
		}
	}

}
