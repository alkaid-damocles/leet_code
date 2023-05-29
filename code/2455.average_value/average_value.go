package code

func averageValue(nums []int) int {

	sum, cut := 0, 0
	for _, num := range nums {
		if num%6 == 0 {
			sum += num
			cut++
		}
	}
	if cut == 0 {
		return 0
	}
	return sum / cut
}
