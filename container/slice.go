package container

import "fmt"

func test() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println("arr[2:6]=", arr[2:6])
	fmt.Println("arr[:6]=", arr[:6])
	fmt.Println("arr[2:]=", arr[2:])
	fmt.Println("arr[:]=", arr[:])
}

func updateSlice(s []int) {
	s[0] = 100
}

func main() {
	test()

	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	s1 := arr[2:6]
	fmt.Println("before updating slice s1.")
	fmt.Println(s1)

	updateSlice(s1)
	fmt.Println("after updating slice s1.")
	fmt.Println(s1)
	fmt.Println(arr)

	arr[0], arr[2] = 0, 2
	// reslice
	s1 = s1[3:5]
	fmt.Println(s1)
	s1 = s1[2:]
	fmt.Println(s1)

}
