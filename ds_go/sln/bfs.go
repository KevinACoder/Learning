package sln

import (
	"fmt"
	"math"
)

func findFirstof(grids [][]int, tp int) (cell [2]int) {
	for i := range grids {
		for j := range grids[i] {
			if grids[i][j] == tp {
				cell[0] = i
				cell[1] = j
				return
			}
		}
	}
	return
}

func orangesRotting(grids [][]int) (elapse int) {
	const (
		Empty  = 0
		Fresh  = 1
		Rotten = 2
	)

	rowNum := len(grids)
	if rowNum == 0 {
		return
	}
	colNum := len(grids[0])

	//build a queue
	start := findFirstof(grids, Rotten)
	//add start pt
	queue := [][2]int{start}
	rowOff, colOff := [4]int{-1, 0, 1, 0}, [4]int{0, -1, 0, 1}

	//iterate util queue is empty
	for len(queue) > 0 {
		num := len(queue)

		for cnt := 0; cnt < num; cnt++ {
			cur := queue[0]
			queue = queue[1:]

			for i := 0; i < 4; i++ {
				x, y := cur[0]+rowOff[i], cur[1]+colOff[i]
				if x < 0 || x >= rowNum ||
					y < 0 || y >= colNum {
					continue
				}

				if grids[x][y] == Fresh {
					grids[x][y] = Rotten
					queue = append(queue, [2]int{x, y})
				}
			}
		}
		elapse++
	}

	//check if exist un-rottened fruits
	for i := range grids {
		for j := range grids[i] {
			if grids[i][j] == Fresh {
				elapse = -1
				return
			}
		}
	}

	return
}

//934. Shortest Bridge
func ShortestBridge(cells [][]int) (flipNum int) {
	rowNum := len(cells)
	if rowNum == 0 {
		return
	}
	colNum := len(cells[0])

	//color one of the island '2' with dfs
	queue := [][2]int{}
	for i := 0; i < rowNum; i++ {
		for j := 0; j < colNum; j++ {
			if cells[i][j] == 1 {
				ShortestBridgeDFS(cells, i, j, &queue)
				goto FOUND
			}
		}
	}

FOUND:
	//find the shortest distance from island '1' to island '2'
	rowOff, colOff := [4]int{-1, 0, 1, 0}, [4]int{0, -1, 0, 1}
	for len(queue) > 0 {
		size := len(queue)

		for cnt := 0; cnt < size; cnt++ {
			cur := queue[0]
			queue = queue[1:]

			for i := 0; i < 4; i++ {
				x, y := cur[0]+rowOff[i], cur[1]+colOff[i]
				if x < 0 || x >= rowNum || y < 0 || y >= colNum || cells[x][y] == 2 {
					continue
				}

				if cells[x][y] == 1 {
					return
				}
				//flip one cell
				cells[x][y] = 2
				queue = append(queue, [2]int{x, y})
			}
		}
		flipNum++
	}

	flipNum = -1
	return
}

func ShortestBridgeDFS(cells [][]int, x, y int, queue *[][2]int) {
	rowNum := len(cells)
	if rowNum == 0 {
		return
	}
	colNum := len(cells[0])
	if x < 0 || x >= rowNum || y < 0 || y >= colNum || cells[x][y] != 1 {
		return
	}

	cells[x][y] = 2 //color island
	*queue = append(*queue, [2]int{x, y})
	ShortestBridgeDFS(cells, x-1, y, queue)
	ShortestBridgeDFS(cells, x+1, y, queue)
	ShortestBridgeDFS(cells, x, y-1, queue)
	ShortestBridgeDFS(cells, x, y+1, queue)
}

//924. Minimize Malware Spread
func minMalwareSpread(graph [][]int, inital []int) (minNode int) {
	const (
		NotAffected  = 0
		Affected     = 1
		NotConnected = 0
		Connected    = 1
	)
	minAffect := math.MaxInt32
	minNode = math.MaxInt32

	//try to remove one node from inital list
	for _, rmNode := range inital {
		//init with 0, no node affected
		badNode := make([]int, len(graph))
		var queue []int

		//set inital bad node, except to be remove one
		for _, node := range inital {
			if node != rmNode {
				badNode[node] = Affected
				queue = append(queue, node)
			}
		}

		//BFS to sum all affected nodes under current
		//  setting
		affected := len(inital) - 1
		for len(queue) > 0 {
			size := len(queue)
			//visit all node of next level
			for cnt := 0; cnt < size; cnt++ {
				cur := queue[0]
				queue = queue[1:]

				//check the connection of every node to
				//  cur node
				for i := 0; i < len(graph); i++ {
					//skip if node is not connected or already
					//  affected
					if graph[cur][i] == NotConnected || badNode[i] == Affected {
						continue
					}
					affected++               //plus affected node num
					badNode[i] = Affected    //mark node as affected
					queue = append(queue, i) //search from
					// newly affected node
				}
			}
		}

		//update the min node
		if affected < minAffect {
			minAffect = affected
			minNode = rmNode
		}
	}

	return
}

func MinMalwareSpreadTest() {
	graph := [][]int{{1, 1, 0}, {1, 1, 0}, {0, 0, 1}}
	fmt.Println(minMalwareSpread(graph, []int{0, 1}))
	graph = [][]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}
	fmt.Println(minMalwareSpread(graph, []int{0, 2}))
	graph = [][]int{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}}
	fmt.Println(minMalwareSpread(graph, []int{1, 2}))
}

//542. 01 Matrix
func updateMatrixBFS(matrix [][]int) (dist [][]int) {
	rowNum, colNum := GetDim(matrix)
	dist = MakeMatrix(rowNum, colNum, math.MaxInt32)
	seen := MakeMatrix(rowNum, colNum, 0) //indicate if cell is visited
	var queue [][2]int

	for i := 0; i < rowNum; i++ {
		for j := 0; j < colNum; j++ {
			//put '0' cell in queue and start search from them
			if matrix[i][j] == 0 {
				queue = append(queue, [2]int{i, j})
				seen[i][j] = 1 //mark cell as visited
			}
		}
	}

	rowOff, colOff := [4]int{-1, 0, 1, 0}, [4]int{0, -1, 0, 1}
	var steps int //count the num of level for searching

	for len(queue) > 0 {
		size := len(queue)
		//search from all cells of next level
		for cnt := 0; cnt < size; cnt++ {
			cur := queue[0]
			queue = queue[1:]
			dist[cur[0]][cur[1]] = steps
			//search neighbor cells from curr cell
			for i := 0; i < 4; i++ {
				x, y := cur[0]+rowOff[i], cur[1]+colOff[i]
				//skip if cell not valid or have been visited
				if x < 0 || x >= rowNum ||
					y < 0 || y >= colNum ||
					seen[x][y] == 1 {
					continue
				}

				seen[x][y] = 1 //mark cell as 'visited'
				queue = append(queue, [2]int{x, y})
			}
		}
		steps++
	}

	return
}

func UpdateMatrixBFSTest() {
	cells := [][]int{{0, 0, 0}, {0, 1, 0}, {1, 1, 1}}
	fmt.Println(updateMatrixBFS(cells))
}

func minMutation(start, end string, bank []string) int {
	Q := []string{start}             //bfs queue
	visited := make(map[string]bool) //visit record
	visited[start] = true

	//closure functio to check if two gene are valid mutation
	validMutation := func(s1, s2 string) bool {
		var cnt int
		if len(s1) != len(s2) {
			return false
		}
		for i := range s1 {
			if s1[i] != s2[i] {
				cnt++
			}
		}
		return cnt == 1
	}

	var minMute int //record the level of bfs
	for len(Q) > 0 {
		size := len(Q)
		//visit all nodes of current level
		for ; size > 0; size-- {
			curr := Q[0]
			Q = Q[1:]
			if curr == end {
				return minMute
			}
			//iterate over bank to find a valid mutation
			for _, gene := range bank {
				if visited[gene] || !validMutation(curr, gene) {
					continue
				}
				visited[gene] = true
				Q = append(Q, gene)
			}
		}
		minMute++
	}

	return -1
}

//787. Cheapest Flights Within K Stops
func findCheapestPrice(n int, flights [][]int, src, dst, K int) (minCost int) {
	type nodeInfo struct {
		node int
		cost int
	}
	//build graph [src] -> [dst, cost]
	G := make(map[int][]nodeInfo)
	for _, f := range flights {
		G[f[0]] = append(G[f[0]], nodeInfo{f[1], f[2]})
	}
	minCost = math.MaxInt32 //set an invalid cost
	Q := []nodeInfo{{src, 0}}
	var steps int

	//do bfs
	for len(Q) > 0 /*&& steps <= K*/ {
		size := len(Q)
		//visit all node of current level
		for ; size > 0; size-- {
			curr := Q[0]
			Q = Q[1:]
			//update the min cost if reached dst
			if curr.node == dst {
				minCost = Min(minCost, curr.cost)
			}
			//iterate over all possible dst for curr node
			for _, ni := range G[curr.node] {
				//prune if cost is greater than current min cost
				if curr.cost+ni.cost > minCost {
					continue
				}
				//add to queue with the cost from origin src to new node
				Q = append(Q, nodeInfo{ni.node, ni.cost + curr.cost})
			}
		}
		if steps > K {
			break
		} //need to be here
		steps++
	}

	if minCost == math.MaxInt32 {
		minCost = -1
	}
	return
}
