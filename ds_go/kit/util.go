package kit

import "math"

func Abs(n int) int {
	if n < 0 {
		return -n
	} else {
		return n
	}
}

func Min(xs ...int) int {
	min := math.MaxInt32
	for _, x := range xs {
		if min > x {
			min = x
		}
	}
	return min
}

func Max(xs ...int) int {
	max := math.MinInt32
	for _, x := range xs {
		if max < x {
			max = x
		}
	}
	return max
}

func Sum(nums []int) (sum int) {
	for _, n := range nums {
		sum = sum + n
	}
	return
}

func GetDim(matrix [][]int) (row, col int) {
	row = len(matrix)
	if row == 0 {
		return
	}
	col = len(matrix[0])
	return
}

func MakeMatrix(row, col, val int) (mat [][]int) {
	mat = make([][]int, row)
	for i := 0; i < row; i++ {
		mat[i] = make([]int, col)
		for j := 0; j < col; j++ {
			mat[i][j] = val
		}
	}
	return mat
}
