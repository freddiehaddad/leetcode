package matrix

// 54. Spiral Matrix
//
// Given an m x n matrix, return all elements of the matrix in spiral order.
//
// Example 1:
//
// Input: matrix = [[1,2,3],[4,5,6],[7,8,9]] Output: [1,2,3,6,9,8,7,4,5]
//
// Example 2:
//
// Input: matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]] Output:
// [1,2,3,4,8,12,11,10,9,5,6,7]
//
// Constraints:
//
//  1. m == matrix.length
//  2. n == matrix[i].length
//  3. 1 <= m, n <= 10
//  4. -100 <= matrix[i][j] <= 100
func spiralOrder(matrix [][]int) []int {
	nElements := len(matrix) * len(matrix[0])
	result := make([]int, nElements)

	rows, cols := len(matrix), len(matrix[0])
	top, bottom, left, right := 0, rows-1, 0, cols-1

	pos := 0
	for pos < nElements {
		for i := left; i <= right; i++ {
			result[pos] = matrix[top][i]
			pos++
		}
		top++

		for i := top; i <= bottom; i++ {
			result[pos] = matrix[i][right]
			pos++
		}
		right--

		if top <= bottom {
			for i := right; i >= left; i-- {
				result[pos] = matrix[bottom][i]
				pos++
			}
			bottom--
		}

		if left <= right {
			for i := bottom; i >= top; i-- {
				result[pos] = matrix[i][left]
				pos++
			}
			left++
		}
	}

	return result
}
