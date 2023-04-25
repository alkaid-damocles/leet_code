package code

import (
	"sort"
)

/*
2418. 按身高排序
给你一个字符串数组 names ，和一个由 互不相同 的正整数组成的数组 heights 。两个数组的长度均为 n 。

对于每个下标 i，names[i] 和 heights[i] 表示第 i 个人的名字和身高。

请按身高 降序 顺序返回对应的名字数组 names 。
*/
func sortPeople(names []string, heights []int) []string {

	heightMap := map[int]string{}

	for i := range heights {
		heightMap[heights[i]] = names[i]
	}
	sort.Slice(heights, func(i, j int) bool {
		return heights[i] > heights[j]
	})
	for i := range heights {
		names[i] = heightMap[heights[i]]
	}
	return names
}
