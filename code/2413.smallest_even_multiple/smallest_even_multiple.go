package code

/*
2413. 最小偶倍数
给你一个正整数 n ，返回 2 和 n 的最小公倍数（正整数）。
*/
func smallestEvenMultiple(n int) int {
	return n * (1 + n%2)
}
