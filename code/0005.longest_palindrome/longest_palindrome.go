package code

/*
5. 最长回文子串
给你一个字符串 s，找到 s 中最长的回文子串。

如果字符串的反序与原始字符串相同，则该字符串称为回文字符串。
*/

func longestPalindrome(s string) string {
	return LongestPalindrome3(s)
}

// LongestPalindrome1 第一种方法, 使用了动态规划
func LongestPalindrome1(s string) string {

	n := len(s)
	if n < 2 {
		return s
	}

	var dp = make([][]bool, n)
	for index := range dp {
		dp[index] = make([]bool, n)
		dp[index][index] = true
	}

	begin := 0
	end := 0
	for l := 2; l <= n; l++ {
		for i := 0; i < n; i++ {
			j := i + l - 1
			if j >= n {
				break
			}
			if s[i] == s[j] && (j-i < 3 || dp[i+1][j-1]) {
				dp[i][j] = true
				begin = i
				end = j
			}
		}
	}
	return s[begin : end+1]
}

// LongestPalindrome2 中心扩展法, 遍历字符串, 选取中心, 扩展两边 直到不是回文串
func LongestPalindrome2(s string) string {
	begin := 0
	end := 0

	for index := range s {
		left1, right1 := expandAroundCenter(s, index, index)
		left2, right2 := expandAroundCenter(s, index, index+1)
		if right1-left1 > end-begin {
			begin, end = left1, right1
		}
		if right2-left2 > end-begin {
			begin, end = left2, right2
		}
	}
	return s[begin : end+1]
}

func expandAroundCenter(s string, left, right int) (int, int) {
	for ; left >= 0 && right < len(s) && s[left] == s[right]; left, right = left-1, right+1 {

	}
	return left + 1, right - 1
}

// LongestPalindrome3 马拉车算法
func LongestPalindrome3(s string) string {

	begin := 0
	end := 0

	t := "#"

	for index := range s {
		t = t + string(s[index]) + "#"
	}
	s = t
	arm_len := []int{}
	j, right := -1, -1
	for i := 0; i < len(s); i++ {
		cur := 0
		if right > i {
			mix_arm := mix(right-i, arm_len[2*j-i])
			cur = expand(s, i-mix_arm, i+mix_arm)
		} else {
			cur = expand(s, i, i)
		}
		arm_len = append(arm_len, cur)
		if i+cur > j {
			j = i
			right = i + cur
		}
		if 2*cur+1 > end-begin {
			begin = i - cur
			end = i + cur
		}
	}

	res := ""
	for i := begin; i <= end; i++ {
		if s[i] != '#' {
			res += string(s[i])
		}
	}
	return res
}

func expand(s string, left, right int) int {
	for ; left >= 0 && right < len(s) && s[left] == s[right]; left, right = left-1, right+1 {

	}
	return (right - left - 2) / 2
}

func mix(a, b int) int {
	if a < b {
		return a
	}
	return b
}
