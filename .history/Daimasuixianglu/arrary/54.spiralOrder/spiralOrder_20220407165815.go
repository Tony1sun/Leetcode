package main

// https://leetcode-cn.com/problems/spiral-matrix/
// 给你一个 m 行 n 列的矩阵 matrix ，请按照 顺时针螺旋顺序 ，返回矩阵中的所有元素。
// 输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
// 输出：[1,2,3,6,9,8,7,4,5]
func spiralOrder(matrix [][]int) []int {

	top, bottom := 0, len(matrix)
	left, right := 0, len(matrix)

	n := len(matrix) * len(matrix)
	ans := []int{}
	for {
		// 第一次循环 上层循环 从左到右
		for i := left; i < right; i++ {
			ans = append(ans, matrix[top][i])
		}

		right--
		// 第二次循环 最右循环 从上到下
		for i := top + 1; i < bottom; i++ {
			ans = append(ans, matrix[i][right])
		}

		bottom--
		// 第三次循环 下层循环 从右到左
		for i := right - 1; i >= left; i-- {
			ans = append(ans, matrix[bottom][i])
		}

		top++
		// 第四次循环 最左循环 从下到上
		for i := bottom - 1; i >= top; i-- {
			ans = append(ans, matrix[i][left])
		}
		left++

		if len(ans) >= n {
			break
		}
	}
	return ans[:n]
}
