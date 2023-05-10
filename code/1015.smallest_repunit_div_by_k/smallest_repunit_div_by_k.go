package code

func smallestRepunitDivByK(k int) int {
	if k%2 == 0 || k%5 == 0 {
		return -1
	}
	n, res := 0, 1
	for {
		n = (n*10 + 1) % k
		if n == 0 {
			return res
		}
		res++

	}
}
