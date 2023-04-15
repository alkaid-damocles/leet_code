package code

import "fmt"

/*
214. 最短回文串
给定一个字符串 s，你可以通过在字符串前面添加字符将其转换为回文串。找到并返回可以用这种方式转换的最短回文串。
*/

func ShortestPalindrome1(s string) string {

	var prefix, suffix uint64 = 0, 0
	var base, mod, mul uint64 = 31, 1000000007, 1
	cur := 0
	for index := range s {
		prefix = (prefix*base + (uint64(s[index]) - 'a' + 1)) % mod
		suffix = (suffix + (uint64(s[index])-'a'+1)*mul) % mod
		mul = mul * base % mod
		if prefix == suffix {
			cur = index
		}
	}
	res := ""
	for i := len(s) - 1; i > cur; i-- {
		res = res + string(s[i])
	}
	res += s
	return res
}

func shortestPalindrome(s string) string {
	s = s + s

	n := len(s)

	next := make([]int, n)
	next[n-1] = 0
	for pre, sif := 0, n-2; sif >= 0; sif-- {
		for pre != 0 && s[n-1-pre] != s[sif] {
			pre = next[pre+1]
		}

		if s[n-1-pre] == s[sif] {
			pre++
			next[sif] = pre
		}
	}
	fmt.Println(next)
	for pre, sif := 0, n-1; sif >= 0; sif-- {
		for pre != 0 && s[n-1-pre] != s[sif] {
			pre = next[pre+1]
		}

		if s[n-1-pre] == s[sif] {
			pre++
		}
		fmt.Println(pre, sif)
		if pre > n/2 {
			return s[n-pre:]
		}
	}
	return ""
}
