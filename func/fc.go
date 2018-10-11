package main

func max(num1, num2 int) int {

	var result int

	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}

	return result
}

func main() {
	a := 100
	b := 10
	c := max(a, b)
	println(c)

	d, e, f := 5, 7, "abc"
	println(d, e, f)
}
