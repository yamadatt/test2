package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run main.go <arg1> <arg2>")
		return
	}

	arg1, err1 := strconv.Atoi(os.Args[1])
	arg2, err2 := strconv.Atoi(os.Args[2])

	if err1 != nil || err2 != nil {
		fmt.Println("Both arguments must be integers.")
		return
	}

	sum := arg1 + arg2
	fmt.Println("Sum:", sum)
}
