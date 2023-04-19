package code

/*
1043. 分隔数组以得到最大和
给你一个整数数组 arr，请你将该数组分隔为长度 最多 为 k 的一些（连续）子数组。分隔完成后，每个子数组的中的所有值都会变为该子数组中的最大值。

返回将数组分隔变换后能够得到的元素最大和。本题所用到的测试用例会确保答案是一个 32 位整数。
*/

func maxSumAfterPartitioning(arr []int, k int) int {

	dp := make([]int, len(arr))
	dp[0] = arr[0]
	for i := 1; i < len(arr); i++ {
		if i < k {
			dp[i] = max(arr[i], dp[i-1]/i) * (i + 1)
		} else {
			cur := arr[i]
			for j := i; j > i-k; j-- {
				cur = max(cur, arr[j])
				dp[i] = max(dp[i], dp[j-1]+cur*(i-j+1))
			}
		}
	}
	return dp[len(arr)-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
