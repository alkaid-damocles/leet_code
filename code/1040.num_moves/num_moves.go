package code

import "sort"

/*
//
在一个长度 无限 的数轴上，第 i 颗石子的位置为 stones[i]。如果一颗石子的位置最小/最大，那么该石子被称作 端点石子 。

每个回合，你可以将一颗端点石子拿起并移动到一个未占用的位置，使得该石子不再是一颗端点石子。

值得注意的是，如果石子像 stones = [1,2,5] 这样，你将 无法 移动位于位置 5 的端点石子，因为无论将它移动到任何位置（例如 0 或 3），该石子都仍然会是端点石子。

当你无法进行任何移动时，即，这些石子的位置连续时，游戏结束。

要使游戏结束，你可以执行的最小和最大移动次数分别是多少？ 以长度为 2 的数组形式返回答案：answer = [minimum_moves, maximum_moves] 。
*/

func numMovesStonesII(stones []int) []int {

	n := len(stones)
	sort.Ints(stones)

	iMax := max(stones[n-1]-stones[1], stones[n-2]-stones[0]) - n + 2
	iMin := iMax
	j := 0
	for i := 0; i < n; i++ {
		for j+1 < n && stones[j+1]+1 <= stones[i]+n {
			j++
		}

		if (j-i+1) == n-1 && stones[j]-stones[i]+1 == n-1 {
			iMin = min(iMin, 2)
		} else {
			iMin = min(iMin, n-(j-i+1))
		}
	}

	return []int{iMin, iMax}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
