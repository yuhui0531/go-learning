package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"reflect"
	"runtime"
	"strconv"
	"time"
)

func main() {
	/*
	   const filename = "abc.txt"
	   contents, err := ioutil.ReadFile(filename)
	   if err != nil {
	       fmt.Println(err)
	   } else {
	       fmt.Printf("%s\n", contents)
	   }

	   fmt.Println(grade(0))
	   fmt.Println(grade(59))
	   fmt.Println(grade(60))
	   fmt.Println(grade(80))
	   fmt.Println(grade(90))
	   fmt.Println(grade(100))
	   //fmt.Println(grade(101))
	   sum := 0
	   for i := 1; i <= 100; i++ {
	       sum += i
	   }
	   fmt.Printf("%d\n", sum)

	   fmt.Println(convert2bin(0))
	   fmt.Println(convert2bin(8))
	   fmt.Println(convert2bin(13))
	   fmt.Println(convert2bin(16))
	   fmt.Println(convert2bin(17))
	   fmt.Println(convert2bin(65535))

	   printFileContent("/tmp/abc.txt")

	   read()

	   q, r := div(13, 4)
	   println(q, r)


	   if result, err := eval(5, 4, "/"); err != nil {
	       print(err)
	   } else {
	       print(result)
	   }
	*/

	fmt.Println(apply(pow, 3, 4))

	fmt.Println(apply(func(a int, b int) int {
		return int(math.Pow(float64(a), float64(b)))
	}, 3, 4))

}

func sum(numbers ...int) int {
	s := 0
	for i := range numbers {
		s += i
	}
	return s
}

func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("calling function %s with parameters(%d, %d)\n", opName, a, b)
	return op(a, b)
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func printFileContent(filename string) {

	environ := os.Environ()
	for k, v := range environ {
		fmt.Println(k, v)
	}

	fmt.Println("aaaaa", os.Getenv("USERNAME"))

	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

}

func forever() {
	for {
		fmt.Println("aaa")
		time.Sleep(1)
	}
}

func convert2bin(n int) string {
	result := ""

	if n == 0 {
		return "0"
	}
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

func grade(score int) string {
	g := ""
	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprintf("Invalid score: %d", score))
	case score < 60:
		g = "F"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score <= 100:
		g = "A"

	}
	return g

}

func read() {
	const filename = "/tmp/abc.txt"
	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}
}

func eval(a int, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		q, _, e := div(a, b)
		return q, e
	default:
		return 0, fmt.Errorf("unsupported optation %s", op)
	}
}

func div(a, b int) (q, r int, e error) {
	if b == 0 {
		return 0, 0, fmt.Errorf("the denominator cannot be zero")
	}
	return a / b, a % b, nil
}
func div2(a, b int) (q, r int) {
	q = a / b
	r = a % b
	return
}
