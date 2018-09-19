package container

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func printString(s string) {
	for i, ch := range s {
		fmt.Printf("(%d %X) ", i, ch)
	}

	fmt.Println(utf8.RuneCountInString(s))

	bytes := []byte(s)

	for len(bytes) > 0 {
		r, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c ", r)
	}

	fmt.Println()

	fmt.Println("rune....")
	for i, ch := range []rune(s) {
		fmt.Printf("(%d %c)", i, ch)
	}
}

func main() {
	//printString("I am Ronny")
	//printString("I'm Ronny")
	//printString("我是中国人,我爱中华国!中国人テスト")

	prefix := strings.HasPrefix("I have", "I")
	print(prefix)
}
