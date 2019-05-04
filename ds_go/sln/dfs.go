package sln

import (
	"math"
	"sort"
)

//417. Pacific Atlantic Water Flow
func pacificAtlantic(matrix [][]int) (targets [][]int) {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return
	}
	R, C := len(matrix), len(matrix[0])

	makeMatrix := func() (mat [][]bool) {
		mat = make([][]bool, R)
		for i := range mat {
			mat[i] = make([]bool, C)
		}
		return
	}

	toP, toA := makeMatrix(), makeMatrix()
	//dfs to check if water in cell is reachable for two ocean
	for i := 0; i < R; i++ {
		pacificAtlanticDFS(matrix, i, 0, 0, toP)   //left
		pacificAtlanticDFS(matrix, i, C-1, 0, toA) //right
	}

	for j := 0; j < C; j++ {
		pacificAtlanticDFS(matrix, 0, j, 0, toP)   //top
		pacificAtlanticDFS(matrix, R-1, j, 0, toA) //bottom
	}

	//get the list of cell that are reachable for both oceans
	for i := 0; i < R; i++ {
		for j := 0; j < C; j++ {
			if toP[i][j] && toA[i][j] {
				targets = append(targets, []int{i, j})
			}
		}
	}

	return
}

func pacificAtlanticDFS(matrix [][]int, r, c, h int, reachAble [][]bool) {
	isValidIndex := func(x, y int) bool {
		return x >= 0 && x < len(matrix) && y >= 0 && y < len(matrix[0])
	}

	if isValidIndex(r, c) && !reachAble[r][c] && matrix[r][c] >= h {
		reachAble[r][c] = true
		pacificAtlanticDFS(matrix, r+1, c, matrix[r][c], reachAble)
		pacificAtlanticDFS(matrix, r-1, c, matrix[r][c], reachAble)
		pacificAtlanticDFS(matrix, r, c+1, matrix[r][c], reachAble)
		pacificAtlanticDFS(matrix, r, c-1, matrix[r][c], reachAble)
	}
}

//996. Number of Squareful Arrays
func numSquarefulPerms(A []int) (ans int) {
	isSquareful := func(x, y int) bool {
		s := int(math.Sqrt(float64(x + y)))
		return int(math.Pow(float64(s), 2)) == x+y
	}

	if len(A) < 2 {
		ans = 0
		return
	}
	if len(A) == 2 {
		if isSquareful(A[0], A[1]) && A[0] != A[1] {
			ans = 1
		}
		return
	}
	sort.Sort(sort.IntSlice(A))               //sort slice in order
	var cur []int                             //pick some elements from array
	used := make([]bool, len(A))              //mark element as 'used'
	numSquarefulPermsDFS(A, &cur, used, &ans) //do dfs
	return
}

func numSquarefulPermsDFS(A []int, cur *[]int, used []bool, ans *int) {
	isSquareful := func(x, y int) bool {
		s := int(math.Sqrt(float64(x + y)))
		return int(math.Pow(float64(s), 2)) == x+y
	}

	lastItem := func(items []int) int {
		if len(items) <= 0 {
			panic("index out of boundary")
		}
		return items[len(items)-1]
	}

	//if find a valid group
	if len(*cur) == len(A) {
		*ans++
		return
	}

	//iterate over all possible new num
	for i, a := range A {
		if used[i] {
			continue
		} //skip if used
		if i > 0 && !used[i-1] && A[i] == A[i-1] {
			continue
		} //skip if duplicate
		//skip if are not able to make a squareful pair
		if len(*cur) > 0 && !isSquareful(lastItem(*cur), a) {
			continue
		}

		*cur = append(*cur, a)
		used[i] = true
		numSquarefulPermsDFS(A, cur, used, ans)
		*cur = (*cur)[:len(*cur)-1]
		used[i] = false
	}
}

//980. Unique Paths III
const (
	START    = 1
	END      = 2
	EMPTY    = 0
	OBSTACLE = -1
)

func uniquePathsIII(grid [][]int) (paths int) {
	var sx, sy int
	empty := 1
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == EMPTY {
				empty++
			} //count the number of empty cells
			if grid[i][j] == START {
				sx = i
				sy = j
			} //find the start point
		}
	}
	//do dfs
	paths = uniquePathsIIIDFS(grid, sx, sy, empty)
	return
}

func uniquePathsIIIDFS(grid [][]int, x, y, empty int) (paths int) {
	//if the cell is not reachable, return 0 paths
	if x < 0 || x == len(grid) || y < 0 ||
		y == len(grid[0]) || grid[x][y] == OBSTACLE {
		return
	}

	//if the cell is 'end' and all empty cells has been visited,
	//  return 1 path, otherwise return 0
	if grid[x][y] == END {
		if empty == 0 {
			paths = 1
		}
		return
	}

	//prevent the current cell from visiting again, and do dfs from
	// the four direction of curr cell
	grid[x][y] = OBSTACLE
	paths = uniquePathsIIIDFS(grid, x+1, y, empty-1) +
		uniquePathsIIIDFS(grid, x-1, y, empty-1) +
		uniquePathsIIIDFS(grid, x, y+1, empty-1) +
		uniquePathsIIIDFS(grid, x, y-1, empty-1)
	grid[x][y] = EMPTY
	return
}
