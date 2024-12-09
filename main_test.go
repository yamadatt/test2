package main

import (
	"strconv"
	"testing"
	"os"
	"bytes"
	"io"
)

func TestMain(t *testing.T) {
	tests := []struct {
		name      string
		args      []string
		want      string
		expectErr bool
	}{
			{"ValidArgs", []string{"main.go", "2", "3"}, "Sum: 5\n", false},
			{"InvalidArgs", []string{"main.go", "a", "3"}, "Both arguments must be integers.\n", true},
			{"MissingArgs", []string{"main.go", "2"}, "Usage: go run main.go <arg1> <arg2>\n", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			oldArgs := os.Args
			defer func() { os.Args = oldArgs }()
			os.Args = tt.args

			if got := captureOutput(main); got != tt.want {
				t.Errorf("main() = %v, want %v", got, tt.want)
			}
		})
	}
}

func captureOutput(f func()) string {
	r, w, _ := os.Pipe()
	stdout := os.Stdout
	defer func() { os.Stdout = stdout }()
	os.Stdout = w

	f()
	w.Close()

	var buf bytes.Buffer
	io.Copy(&buf, r)
	return buf.String()
}
