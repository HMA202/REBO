package main

func RotateMatrix(matrix [][]int) [][]int {
	n := len(matrix)

	result := make([][]int, n)
	for i := 0; i < n; i++ {
		result[i] = make([]int, n)
	}

	for row := 0; row < n; row++ {
		for col := 0; col < n; col++ {
			result[col][n-1-row] = matrix[row][col]
		}
	}

	return result
}
