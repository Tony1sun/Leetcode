package main

func findAnagrams(s string, p string) (ans []int) {
	res := make([]int, 0)
	if len(s) < len(p) {
		return res
	}
	arr := make([]int, 200)
	for i := 0; i < len(p); i++ {
		arr[s[i]]++
		arr[p[i]]--
	}
	left := 0
	right := len(p) - 1

	for {
		if isValid(arr) {
			res = append(res, left)
		}
		arr[s[left]]--
		left++
		if right+1 >= len(s) {
			break
		}
		right++
		arr[s[right]]++
	}
	return res
}

func isValid(arr []int) bool {
	for _, n := range arr {
		if n != 0 {
			return false
		}
	}
	return true
}
