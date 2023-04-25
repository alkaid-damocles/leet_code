package code

/*
1163. 按字典序排在最后的子串
给你一个字符串 s ，找出它的所有子串并按字典序排列，返回排在最后的那个子串。
*/

func lastSubstring(s string) string {
	i, n := 0, len(s)

	for j, k := 1, 0; j+k < n; {
		if s[i+k] == s[j+k] {
			k++
		} else if s[i+k] < s[j+k] {
			i = i + k + 1
			k = 0
			if i >= j {
				j = i + 1
			}
		} else {
			j = j + k + 1
			k = 0
		}
	}
	return s[i:]
}
