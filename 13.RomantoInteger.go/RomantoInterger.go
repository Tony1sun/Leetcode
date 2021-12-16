package Leetcode

var roman = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToInt(s string) int {
	var ans int
	for i := range s {
		value := roman[s[i]]
		if i < len(s)-1 && value < roman[s[i+1]] {
			ans -= value
		} else {
			ans += value
		}
	}
	return ans
}
