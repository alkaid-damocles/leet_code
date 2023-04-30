package code

/*
1033. 移动石子直到连续
三枚石子放置在数轴上，位置分别为 a，b，c。

每一回合，你可以从两端之一拿起一枚石子（位置最大或最小），并将其放入两端之间的任一空闲位置。形式上，假设这三枚石子当前分别位于位置 x, y, z 且 x < y < z。那么就可以从位置 x 或者是位置 z 拿起一枚石子，并将该石子移动到某一整数位置 k 处，其中 x < k < z 且 k != y。

当你无法进行任何移动时，即，这些石子的位置连续时，游戏结束。

要使游戏结束，你可以执行的最小和最大移动次数分别是多少？ 以长度为 2 的数组形式返回答案：answer = [minimum_moves, maximum_moves]
*/

func numMovesStones(a int, b int, c int) []int {
	res := []int{2, max(max(a, b), c) - min(min(a, b), c) - 2}
	if a+b+c-2*min(min(a, b), c)-max(max(a, b), c) == 1 && 2*(max(max(a, b), c))-((a+b+c)-min(min(a, b), c)) == 1 {
		res[0] = 0
	} else if a+b+c-2*min(min(a, b), c)-max(max(a, b), c) <= 2 || 2*(max(max(a, b), c))-((a+b+c)-min(min(a, b), c)) <= 2 {
		res[0] = 1
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
