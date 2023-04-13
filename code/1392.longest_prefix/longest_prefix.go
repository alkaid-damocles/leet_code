package code

import "fmt"

/*
1392. 最长快乐前缀
「快乐前缀」 是在原字符串中既是 非空 前缀也是后缀（不包括原字符串自身）的字符串。

给你一个字符串 s，请你返回它的 最长快乐前缀。如果不存在满足题意的前缀，则返回一个空字符串 "" 。
*/

//

/* 这里是字符串编码的解法
func longestPrefix(s string) string {
	var prefix, suffix uint64 = 0, 0
	var base, mod, mul uint64 = 31, 1000000007, 1
	n := len(s)

	cur := 0
	for i := 1; i < n; i++ {
		prefix = (prefix*base + (uint64(s[i-1]) - 'a')) % mod
		suffix = (suffix + (uint64(s[n-i])-'a')*mul) % mod
		mul = mul * base % mod
		if prefix == suffix {
			cur = i
		}
	}
	return s[:cur]
}

*/

func longestPrefix(s string) string {
	return s[:next(s)[len(s)-1]]
}

func next(str string) []int {
	next := make([]int, len(str))
	next[0] = 0

	p := 0
	for s := 1; s < len(str); s++ {
		for p != 0 && str[p] != str[s] {
			p = next[p-1]
		}
		if str[p] == str[s] {
			p++
			next[s] = p
		}
	}
	fmt.Println(next)
	return next
}
