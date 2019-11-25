package main

import "fmt"
/**
最大不重复的子川
 */
func lengthOfNorepeatingSubstr(s string) int {
	lastOccurred := make(map[byte]int)
	start := 0
	maxLength := 0

	for i, ch := range []byte(s) {

		if lasti, ok := lastOccurred[ch]; ok && lasti >= start{
			start=lasti + 1
		}
		if i - start + 1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}
	return maxLength
}

func main() {
	fmt.Println(lengthOfNorepeatingSubstr("aabaacaaa"))
}
