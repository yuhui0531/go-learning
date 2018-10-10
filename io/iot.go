package main

import (
	"fmt"
	"runtime"
)

const (
	B = 1 << (10 * iota)
	KB
	MB
	GB
	TB
)

func main() {
	fmt.Println(1 << 3)
	fmt.Println(B)
	fmt.Println(KB)
	fmt.Println(MB)
	fmt.Println(GB)
	fmt.Println(TB)

	runtime.NumCPU()

	var arr1 [5]int
	var arr2 = [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr1, arr2)

	fmt.Println(deferTest(10))
	fmt.Println(deferTest2(10))
	fmt.Println('a' * 20)
}

func deferTest2(number int) int {
	defer func() {
		number++
		fmt.Println("three:", number)
	}()

	defer func() {
		number++
		fmt.Println("two:", number)
	}()

	defer func() {
		number++
		fmt.Println("one:", number)
	}()

	return number
}

func deferTest(number int) int {
	defer func(n int) {
		n++
		fmt.Println("three:", n)
	}(number)

	defer func(n int) {
		n++
		fmt.Println("two:", n)
	}(number)

	defer func(n int) {
		n++
		fmt.Println("one:", n)
	}(number)

	return number
}
