package main

import "fmt"

// go 语言只有值传递一种方式

func main() {

	a, b := 3, 4
	swap(a, b)
	fmt.Println(a, b)

	swap2(&a, &b)
	fmt.Println(a, b)

	a, b = swap3(a, b)
	fmt.Println(a, b)

}

func swap(a, b int) {
	a, b = b, a
}

func swap2(a, b *int) {
	*a, *b = *b, *a
}

func swap3(a, b int) (int, int) {
	return b, a
}
