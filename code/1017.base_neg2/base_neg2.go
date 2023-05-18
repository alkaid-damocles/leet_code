package code

import "strconv"

func baseNeg2(n int) string {
	if n == 0 || n == 1 {
		return strconv.Itoa(n)
	}
	res := []byte{}
	for n != 0 {
		remainder := n & 1
		res = append(res, '0'+byte(remainder))
		n -= remainder
		n /= -2
	}
	for i, n := 0, len(res); i < n/2; i++ {
		res[i], res[n-1-i] = res[n-1-i], res[i]
	}
	return string(res)
}
