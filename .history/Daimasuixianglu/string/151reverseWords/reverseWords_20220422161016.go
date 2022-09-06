package Leetcode

// https://leetcode-cn.com/problems/reverse-words-in-a-string/
// 给定一个字符串，逐个翻转字符串中的每个单词。
// 示例 1：
// 输入: "the sky is blue"
// 输出: "blue is sky the"

// 示例 2：
// 输入: "  hello world!  "
// 输出: "world! hello"
// 解释: 输入字符串可以在前面或者后面包含多余的空格，但是反转后的字符不能包括。

// 示例 3：
// 输入: "a good   example"
// 输出: "example good a"
// 解释: 如果两个单词间有多余的空格，将反转后单词间的空格减少到只含一个。

// 思路
// 1、移除多余空格
// 2、将整个字符串反转
// 3、将每个单词反转
func reverseWords(s string) string {
	// 使用双指针删除冗余的空格
	slowIndex, fastIndex := 0, 0
	b := []byte(s)
	// 删除头部冗余空格
	for len(b) > 0 && fastIndex < len(b) && b[fastIndex] == ' ' {
		fastIndex++
	}
	// 删除单词间冗余空格
	for fastIndex := 0; fastIndex < len(b); fastIndex++ {
		if fastIndex-1 > 0 && b[fastIndex-1] == b[fastIndex] && b[fastIndex] == ' ' {
			continue
		}
		b[slowIndex] = b[fastIndex]
		slowIndex++
	}
	// 删除尾部冗余空格
	if slowIndex-1 > 0 && b[slowIndex-1] == ' ' {
		b = b[:slowIndex-1]
	} else {
		b = b[:slowIndex]
	}
	// 反转整个字符串
	reverse(&b, 0, len(b)-1)
	// 反转单个单词,i开始位置,j结束位置
	i := 0
	for i < len(b) {
		j := i
		// 如果没有空格，遍历到有空格为止(这样才是整个单词，为了反转)
		for ; j < len(b) && b[j] != ' '; j++ {
		}
		reverse(&b, i, j-1)
		i = j
		i++
	}
	return string(b)
}

func reverse(b []byte, left, right int) {
	for left < right {
		b[left], b[right] = b[right], )[left]
		left++
		right--
	}
}
