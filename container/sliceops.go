package container

import "fmt"

func printSlice(s []int) {
	fmt.Println("slice content")
	fmt.Println(s)
	fmt.Println("slice properties")
	fmt.Printf("len=%d, cap=%d\n", len(s), cap(s))
}

func main() {
	//var s []int // Zero value for slice is nil
	//fmt.Println(s)
	//
	//for i := 0; i < 100; i++ {
	//	printSlice(s)
	//	s = append(s, 2*i+1)
	//}
	//
	//fmt.Println(s)

	createSlice()
	copySlice()
	deleteSlice()
	popSlice()
	shiftSlice()

}

func createSlice() {
	fmt.Println("creating slice...")
	s2 := make([]int, 16)
	s3 := make([]int, 16, 32)

	printSlice(s2)
	printSlice(s3)
}

func copySlice() {
	fmt.Println("copying slice...")
	s1 := []int{2, 4, 6, 8}
	s2 := make([]int, 16)
	copy(s2, s1)
	printSlice(s1)
	printSlice(s2)
}

func deleteSlice() {
	fmt.Println("deleting slice...")
	s1 := []int{2, 4, 6, 8, 10, 12, 14}
	s1 = append(s1[:3], s1[4:]...)
	printSlice(s1)
}

func popSlice() {
	fmt.Println("popping slice...")
	s1 := []int{2, 4, 6, 8, 10, 12, 14}
	s1 = s1[1:]
	printSlice(s1)
}

func shiftSlice() {
	fmt.Println("shifting slice...")
	s1 := []int{2, 4, 6, 8, 10, 12, 14}
	s1 = s1[:len(s1)-1]
	printSlice(s1)
}
