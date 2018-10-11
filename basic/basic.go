package main

import (
	"fmt"
	"golang.org/x/tools/container/intsets"
)

func variableWithEmtpy() {
	var a int
	var b string

	fmt.Printf("a=%d,b=%q\n", a, b)
}

func variableWithInitValue() {
	var a, b = 3, 4
	var s = "abc"
	fmt.Printf("a=%d,b=%d,s=%q\n", a, b, s)
}

func variableTypeDetection() {
	var a, b, c, s = 3, 4, true, "abc"
	fmt.Printf("a=%d,b=%d,c=%v,s=%q\n", a, b, c, s)
}

func main() {
	variableWithEmtpy()
	variableWithInitValue()
	variableTypeDetection()
	testSparse()
}

func testSparse() {

	s := intsets.Sparse{}
	s.Insert(1)
	s.Insert(1000)
	fmt.Println(s.Has(1000))
}
