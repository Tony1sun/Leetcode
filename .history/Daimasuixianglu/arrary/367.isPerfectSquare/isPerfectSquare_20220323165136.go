package main

func isPerfectSquare(num int) bool {
	i := 0
	for i * i <= num {
		if i * i == num {
			return true
		}
		i++
	}
}
