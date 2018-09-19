package main

import "fmt"

func main() {

	const (
		a = iota //0
		b        //1
		c        //2
		d = "ha" //独立值，iota += 1
		e        //"ha"   iota += 1
		f = 100  //iota +=1
		g        //100  iota +=1
		h = iota //7,恢复计数
		i        //8
	)

	fmt.Println(a, b, c, d, e, f, g, h, i)

	enums()

}

func bounded(v int) int {
	if v > 100 {
		return 100
	} else if v < 100 {
		return 1
	}
	return 0
}

func enums() {
	const (
		cpp = iota
		java
		python
		golang
		javascript
	)

	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)

	fmt.Println(b, kb, mb, gb, tb, pb)

	fmt.Println(cpp, java, python, golang, javascript)
}
