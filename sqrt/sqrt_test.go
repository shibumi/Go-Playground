package main

import "testing"

func TestSqrt(t *testing.T) {
	tables := []struct {
		x float64
		d float64
		z float64
	}{
		{2, 0.001, 1.4142135623746899},
		{9, 0.001, 3.000000001396984},
		{13, 0.001, 3.6055513629176015},
	}
	for _, table := range tables {
		result := Sqrt(table.x, table.d)
		if result != table.z {
			t.Errorf("Square root for: x: %f, d: %f was incorrect, got: %f, want: %f", table.x, table.d, result, table.z)
		}
	}
}
