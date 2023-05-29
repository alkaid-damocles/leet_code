package code

import (
	"fmt"
	"math"
)

func shortestPathBinaryMatrix(grid [][]int) int {
	var res int = math.MaxInt64
	backTrack(0, 0, grid, &res, 1)
	if grid[0][0] == 1 || res == math.MaxInt64 {
		return -1
	}
	return res
}

func backTrack(i, j int, grid [][]int, res *int, cur int) {
	fmt.Println(i, j, *res, cur)
	if i == len(grid) || j == len(grid[0]) {
		return
	}
	if i == len(grid)-1 && j == len(grid[0])-1 {
		*res = min(*res, cur)
	}
	if grid[i][j] != 0 {
		return
	}
	backTrack(i+1, j, grid, res, cur+1)
	backTrack(i, j+1, grid, res, cur+1)
	backTrack(i+1, j+1, grid, res, cur+1)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
