package container

import "fmt"

func find(s string) int {

	lastOccurred := make(map[byte]int)
	start, maxLength := 0, 0

	for i, ch := range []byte(s) {

		if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
			start = lastI + 1
		}

		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}
	return maxLength
}

func find2(s string) int {

	lastOccurred := make(map[rune]int)
	start, maxLength := 0, 0

	for i, ch := range []rune(s) {

		if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
			start = lastI + 1
		}

		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}

	return maxLength
}

func main() {
	fmt.Println(find("abcabcbb"))
	fmt.Println(find("bbbbb"))
	fmt.Println(find("b"))
	fmt.Println(find("abcde"))
	fmt.Println(find(""))
	fmt.Println(find2("我是中国人,我爱中华国!テスト"))
}
