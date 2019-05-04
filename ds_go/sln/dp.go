package sln

import (
	. "..\kit"
)

//542. 01 Matrix
func updateMatrix(matrix [][]int) (dist [][]int) {
	rowNum, colNum := GetDim(matrix)
	if rowNum == 0 || colNum == 0 {return}
	//make matrix, init cell value that is impossible to overflow
	dist = MakeMatrix(rowNum, colNum, math.MaxInt32 - rowNum * colNum)

	//up, left
	for i := 0; i < rowNum; i++ {
		for j := 0; j < colNum; j++ {
			if matrix[i][j] == 1 {
				//if current cell is '1', find the smallest possible
				//  distance
				if i > 0 {dist[i][j] = Min(dist[i][j], dist[i-1][j] + 1)}
				if j > 0 {dist[i][j] = Min(dist[i][j], dist[i][j-1] + 1)}	
			}else{
				dist[i][j] = 0;
			}
		}
	}

	//down, right
	for i := rowNum - 1; i >= 0; i-- {
		for j := colNum - 1; j >= 0; j-- {
			if i < rowNum - 1 {dist[i][j] = Min(dist[i][j], dist[i+1][j] + 1)}
			if j < colNum - 1 {dist[i][j] = Min(dist[i][j], dist[i][j+1] + 1)}
		}
	}

	return
}