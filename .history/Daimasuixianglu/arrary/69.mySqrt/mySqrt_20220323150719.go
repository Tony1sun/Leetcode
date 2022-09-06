package main

func mySqrt(x int) int {
	low, high := , x
	for low <= high {
		mid := low + (high-low)/2
		if x/mid < mid {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return high
}
