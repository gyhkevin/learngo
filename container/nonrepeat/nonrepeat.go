package main

import "fmt"
// 寻找最长不含有重复字符的子串
func lengthOfNonRepeatingSubString(s string) int {
	lastOccurred := make(map[rune]int)
	start := 0
	maxLength := 0
	for i, ch := range []rune(s) {

		if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
			start = lastOccurred[ch] + 1
		}
		if i-start+1 > maxLength {
			maxLength = i-start+1
		}
		lastOccurred[ch] = i
	}

	return maxLength
}

func main()  {
	fmt.Println(lengthOfNonRepeatingSubString("abcabcbb"))
	fmt.Println(lengthOfNonRepeatingSubString("bbbbbb"))
	fmt.Println(lengthOfNonRepeatingSubString("pwwkew"))
	fmt.Println(lengthOfNonRepeatingSubString("一二三三二"))
	fmt.Println(lengthOfNonRepeatingSubString("寻找最长不含有重复字符的子串"))

}
