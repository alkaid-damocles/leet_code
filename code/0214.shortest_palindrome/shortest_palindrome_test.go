package code

import (
	"fmt"
	"testing"
)

/*
214. 最短回文串
给定一个字符串 s，你可以通过在字符串前面添加字符将其转换为回文串。找到并返回可以用这种方式转换的最短回文串。
*/

func TestShortestPalindrome(t *testing.T) {
	fmt.Println(shortestPalindrome("aacecaaa"))
}
