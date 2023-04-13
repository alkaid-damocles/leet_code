package code

func lengthOfLongestSubstring(s string) int {

	strIndex := map[rune]int{}

	left, res := 0, 0

	for index, value := range s {
		if strIndex[value] != 0 {
			left = max(strIndex[value], left)
		}
		res = max(res, index-left+1)
		strIndex[value] = index + 1
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
