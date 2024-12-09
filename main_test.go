package main

import (
	"strings"
	"testing"

	//	"strconv"
	"bytes"
	"io"
	"os"
)

func TestMain(t *testing.T) {
	tests := []struct {
		args     []string
		expected string
	}{
		{[]string{"main.go", "3", "5"}, "8"},
		{[]string{"main.go", "10", "20"}, "30"},
		{[]string{"main.go", "0", "0"}, "0"},
		{[]string{"main.go", "a", "b"}, "Both arguments must be integers."},
		{[]string{"main.go", "1"}, "Usage: go run main.go <num1> <num2>"},
	}

	for _, test := range tests {
		os.Args = test.args
		output := captureOutput(main)
		if output != test.expected {
			t.Errorf("For args %v, expected %v but got %v", test.args, test.expected, output)
		}
	}
}

func captureOutput(f func()) string {
	r, w, _ := os.Pipe()
	stdout := os.Stdout
	os.Stdout = w

	outC := make(chan string)
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outC <- buf.String()
	}()

	f()
	w.Close()
	os.Stdout = stdout
	out := <-outC

	return strings.TrimSpace(out)
}
