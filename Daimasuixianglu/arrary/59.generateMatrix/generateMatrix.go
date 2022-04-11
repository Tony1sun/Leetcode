package Leetcode

// https://leetcode-cn.com/problems/spiral-matrix-ii/
// 59. 螺旋矩阵 II
// 给你一个正整数 n ，生成一个包含 1 到 n2 所有元素，且元素按顺时针顺序螺旋排列的 n x n 正方形矩阵 matrix 。
func generateMatrix(n int) [][]int {
	// 上 下
	top, bottom := 0, n-1
	// 左 右
	left, right := 0, n-1
	num := 1
	tar := n * n
	// 初始化要填充的正方形
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}
	for num <= tar {
		// 模拟上侧从左到右
		for i := left; i <= right; i++ {
			matrix[top][i] = num
			num++
		}
		top++
		// 模拟右侧从上到下
		for i := top; i <= bottom; i++ {
			matrix[i][right] = num
			num++
		}
		right--
		// 模拟下侧从右到左
		for i := right; i >= left; i-- {
			matrix[bottom][i] = num
			num++
		}
		bottom--
		// 模拟左侧从下到上
		for i := bottom; i >= top; i-- {
			matrix[i][left] = num
			num++
		}
		left++
	}
	return matrix
}
