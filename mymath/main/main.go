package main

import (
	"fmt"
	"go-learning/mymath"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("please input number")
		return
	}
	num, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("input must be number", err)
		return
	}
	fmt.Println("Fibonacci :", num, mymath.Fibonacci(num))
}
