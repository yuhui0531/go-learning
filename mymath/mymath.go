package mymath

func Fibonacci(num int) int {
	if num == 0 || num == 1 {
		return num
	}
	return Fibonacci(num-1) + Fibonacci(num-2)
}
