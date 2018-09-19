package container

import "fmt"

// GO语言不直接使用数组,直接使用slice
func printArray(arr [5]int) {
	for k, v := range arr {
		fmt.Println(k, v)
	}
}

func printArray2(arr *[5]int) {
	arr[0] = 100
	for k, v := range arr {
		fmt.Println(k, v)
	}
}

func findMax(arr [5]int) {
	maxi, maxv := -1, -1

	for k, v := range arr {
		if v > maxv {
			maxi, maxv = k, v
			fmt.Printf("maxi=%d, maxv=%d\n", maxi, maxv)
		}
	}
}

func main() {
	var arr1 [5]int
	arr2 := [3]int{1, 3, 5}
	arr3 := [...]int{2, 6, 5, 15, 8}
	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println("original array list:")
	fmt.Println(arr3)

	fmt.Println("printArray(arr3)")
	printArray(arr3)

	fmt.Println("changed array list:")
	fmt.Println(arr3)

	fmt.Println("printArray2(arr3)")
	printArray2(&arr3)

	fmt.Println("changed array list:")
	fmt.Println(arr3)

	fmt.Println("find max element from array list:")
	findMax(arr3)

	// 数组传参是值传递,会直接拷贝过去,生成一模一样的数组拷贝 .
}
