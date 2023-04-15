package code

import "unicode"

func camelMatch(queries []string, pattern string) []bool {
	n := len(queries)
	res := make([]bool, n)
	for i := 0; i < n; i++ {
		res[i] = true
		p := 0
		for _, c := range queries[i] {
			if p < len(pattern) && pattern[p] == byte(c) {
				p++
			} else if unicode.IsUpper(c) {
				res[i] = false
				break
			}
		}
		if p < len(pattern) {
			res[i] = false
		}
	}
	return res
}
