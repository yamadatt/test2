package main

import (
	"os"
	"testing"
)

func TestMain(t *testing.T) {
	tests := []struct {
		args     []string
		expected string
	}{
		{[]string{"main.go", "1", "2"}, "Result: 3\n"},
		{[]string{"main.go", "a", "2"}, "Error: Argument 1 is not a valid integer\n"},
		{[]string{"main.go", "1", "b"}, "Error: Argument 2 is not a valid integer\n"},
		{[]string{"main.go"}, "Usage: go run main.go <arg1> <arg2>\n"},
	}

	for _, test := range tests {
		os.Args = test.args
		output := captureOutput(main)
		if output != test.expected {
			t.Errorf("For args %v, expected %q but got %q", test.args, test.expected, output)
		}
	}
}

func captureOutput(f func()) string {
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	w.Close()
	var buf bytes.Buffer
	io.Copy(&buf, r)
	return buf.String()
}
