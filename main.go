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
	if err1 != nil {
		fmt.Println("Error: Argument 1 is not a valid integer")
		return
	}

	arg2, err2 := strconv.Atoi(os.Args[2])
	if err2 != nil {
		fmt.Println("Error: Argument 2 is not a valid integer")
		return
	}

	result := arg1 + arg2
	fmt.Println("Result:", result)
}
