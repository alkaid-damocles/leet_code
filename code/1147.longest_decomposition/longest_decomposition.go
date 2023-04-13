package code

/*
1147. 段式回文
你会得到一个字符串 text 。你应该把它分成 k 个子字符串 (subtext1, subtext2，…， subtextk) ，要求满足:

subtexti 是 非空 字符串
所有子字符串的连接等于 text ( 即subtext1 + subtext2 + ... + subtextk == text )
对于所有 i 的有效值( 即 1 <= i <= k ) ，subtexti == subtextk - i + 1 均成立
返回k可能最大值
*/

func longestDecomposition(text string) int {
	var prefix, suffix uint64 = 0, 0
	var base, mod, mul uint64 = 31, 1000000007, 1

	n := len(text)
	res := 0
	cur := -1
	for i := 0; i < n/2; i++ {
		prefix = (prefix*base + (uint64(text[i]) - 'a')) % mod
		suffix = (suffix + (uint64(text[n-1-i])-'a')*mul) % mod
		mul = mul * base % mod
		if prefix == suffix {
			cur = i + 1
			res += 2
			prefix, suffix, mul = 0, 0, 1
		}
	}
	if cur < n/2 || n%2 == 1 {
		res++
	}
	return res
}
