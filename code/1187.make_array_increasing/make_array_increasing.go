package code

import (
	"math"
	"sort"
)

/*
1187. 使数组严格递增
给你两个整数数组 arr1 和 arr2，返回使 arr1 严格递增所需要的最小「操作」数（可能为 0）。

每一步「操作」中，你可以分别从 arr1 和 arr2 中各选出一个索引，分别为 i 和 j，0 <= i < arr1.length 和 0 <= j < arr2.length，然后进行赋值运算 arr1[i] = arr2[j]。

如果无法让 arr1 严格递增，请返回 -1。
*/

func makeArrayIncreasing(arr1 []int, arr2 []int) int {
	sort.Ints(arr2)
	n := len(arr1)
	m := len(arr2)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, min(m, n)+1)
		for j := range dp[i] {
			dp[i][j] = math.MaxInt
		}
	}
	dp[0][0] = -1
	for i := 1; i <= len(arr1); i++ {
		for j := 0; j <= min(i, m); j++ {
			if arr1[i-1] > dp[i-1][j] {
				dp[i][j] = arr1[i-1]
			}
			if j > 0 && dp[i-1][j-1] != math.MaxInt {
				k := j - 1 + sort.SearchInts(arr2[j-1:], dp[i-1][j-1]+1)
				if k < m {
					dp[i][j] = min(dp[i][j], arr2[k])
				}
			}
			if i == n && dp[i][j] != math.MaxInt {
				return j
			}
		}
	}
	return -1
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
