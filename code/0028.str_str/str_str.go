package code

import "fmt"

/*

28. 找出字符串中第一个匹配项的下标
给你两个字符串 haystack 和 needle ，请你在 haystack 字符串中找出 needle 字符串的第一个匹配项的下标（下标从 0 开始）。如果 needle 不是 haystack 的一部分，则返回  -1 。
*/
//  这道题其实还是考查的kmp算法, 不过是需要将两个字符串匹配在一起.
func strStr(haystack string, needle string) int {

	n, m := len(haystack), len(needle)

	if m == 0 {
		return 0
	}
	next := make([]int, m)
	next[0] = 0
	for p, s := 0, 1; s < m; s++ {
		for p != 0 && needle[p] != needle[s] {
			p = next[p-1]
		}
		if needle[p] == needle[s] {
			p++
			next[s] = p
		}

	}
	fmt.Println(next)
	// 这里其实是把字符串needle 和 haystack 按照 needle#haystack的形式拼接, 然后用kmp算法而已, 直到p = needle的长度的时候, 就可以退出了
	// 但是为了压缩内存占用, 所以不再显示的拼接字符串, 只需要找到第一个, 所以不需要保存后续的next数组,
	for p, s := 0, 0; s < n; s++ {
		for p != 0 && needle[p] != haystack[s] {
			p = next[p-1]
		}
		if needle[p] == haystack[s] {
			p++
		}
		fmt.Println(p, s)
		if p == m {
			return s - m + 1
		}
	}
	return -1
}
