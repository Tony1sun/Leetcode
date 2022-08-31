package main

func mySqrt(x int) int {
	low, high := 1, x
	for low <= high {
		mid := (low + high) / 2
		if x / mid < mid {
			high = mid + 1
		}
	}
}
