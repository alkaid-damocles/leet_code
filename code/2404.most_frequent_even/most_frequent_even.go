package code

import (
	"math"
)

/*
2404. 出现最频繁的偶数元素
给你一个整数数组 nums ，返回出现最频繁的偶数元素。

如果存在多个满足条件的元素，只需要返回 最小 的一个。如果不存在这样的元素，返回 -1 。
*/

func mostFrequentEven(nums []int) int {

	counts := map[int]int{}
	count := 0
	res := math.MaxInt64
	for _, num := range nums {
		if num%2 != 0 {
			continue
		}
		counts[num]++
		if counts[num] > count {
			res = num
			count = counts[num]
		} else if counts[num] == count && num < res {
			res = num
		}
	}
	if len(counts) == 0 {
		return -1
	}
	return res
}
