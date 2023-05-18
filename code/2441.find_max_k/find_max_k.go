package code

/*
2441. 与对应负数同时存在的最大正整数
给你一个 不包含 任何零的整数数组 nums ，找出自身与对应的负数都在数组中存在的最大正整数 k 。

返回正整数 k ，如果不存在这样的整数，返回 -1 。
*/

func findMaxK(nums []int) int {

	set := map[int]struct{}{}

	res := -1
	for _, num := range nums {
		_, ok := set[-num]
		if ok {
			res = max(res, abs(num))
		} else {
			set[num] = struct{}{}
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}
