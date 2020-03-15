package main

import "fmt"
/**
最大不重复的子川
 */
func lengthOfNorepeatingSubstr(s string) int {

	//某个字母最后一次出现的位置
	lastOccurred := make(map[rune]int)
	//不重复开始的位置
	start := 0
	//最大长度
	maxLength := 0

	for i, ch := range []rune(s) {

		//判定是否更改其起始位置，map中存在这个字符，并且这个字符位置在当前位置的前面
		if lasti, ok := lastOccurred[ch]; ok && lasti >= start{
			start=lasti + 1
		}
		//判定起始位置到现有位置的长度是否大于最大长度
		if i - start + 1 > maxLength {
			maxLength = i - start + 1
		}
		//记录当前字母的位置
		lastOccurred[ch] = i
	}
	return maxLength
}

func lengthOfNorepeatingSubstr2(s string)  int {

	//某个字母最后一次出现的位置,利用数组，也就是字符转成数字的形式，存在数组里
	//开始默认都是0，下表也是从0开始，所以要+1
	lastOccurred := make([]int,0xfff)
	//不重复开始的位置
	start := 0
	//最大长度
	maxLength := 0

	for i, ch := range []rune(s) {

		//判定是否更改其起始位置，map中存在这个字符，并且这个字符位置在当前位置的前面
		if lasti := lastOccurred[ch];  lasti > start{
			start=lasti
		}
		//判定起始位置到现有位置的长度是否大于最大长度
		if i - start + 1 > maxLength {
			maxLength = i - start + 1
		}
		//记录当前字母的位置
		lastOccurred[ch] = i + 1
	}
	return maxLength
}

func main() {
	fmt.Println(lengthOfNorepeatingSubstr2("aabaacaaa"))
	 fmt.Println(lengthOfNorepeatingSubstr2("dasddsadddddddsssdabcdefg"))
	fmt.Println([]rune("aabaacaaa")[8:])
	fmt.Println([]rune("dasddsadddddddsssdabcdefg")[18:])
	 a :=string(11)
	 b := "ss"
	 fmt.Printf("%T,%s",a,b)
}
