package code

/*
796. 旋转字符串
给定两个字符串, s 和 goal。如果在若干次旋转操作之后，s 能变成 goal ，那么返回 true 。

s 的 旋转操作 就是将 s 最左边的字符移动到最右边。

例如, 若 s = 'abcde'，在旋转一次之后结果就是'bcdea' 。
*/
func rotateString(s string, goal string) bool {
	s = s + s
	m, n := len(goal), len(s)

	if 2*m != n {
		return false
	}

	next := make([]int, m)
	for pre, sif := 0, 1; sif < m; sif++ {
		for pre != 0 && goal[pre] != goal[sif] {
			pre = next[pre-1]
		}
		if goal[pre] == goal[sif] {
			pre++
			next[sif] = pre
		}
	}

	for pre, sif := 0, 0; sif < n; sif++ {
		for pre != 0 && goal[pre] != s[sif] {
			pre = next[pre-1]
		}
		if goal[pre] == s[sif] {
			pre++
		}
		if pre == m {
			return true
		}
	}

	return false
}
