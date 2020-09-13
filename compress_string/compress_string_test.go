package compress

import "testing"

func TestCompressString(t *testing.T) {
	tables := []struct {
		input  string
		output string
	}{
		{input: "AAABCCDDDD", output: "A3BC2D4"},
		{input: "", output: ""},
		{input: "AAABCD", output: "A3BCD"},
		{input: "ABCD", output: "ABCD"},
		{input: "AABBCC", output: "AABBCC"},
	}
	for _, table := range tables {
		output := compressString(table.input)
		if table.output != output {
			t.Errorf("error in test %s, %s expected: %s", table.input, output, table.output)
		}
	}
}
