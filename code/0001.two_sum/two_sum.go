package code

func TwoSum(nums []int, target int) []int {
	hashTable := map[int]int{}
	for i, x := range nums {
		if j, ok := hashTable[target-x]; ok {
			return []int{i, j}
		} else {
			hashTable[x] = i
		}
	}
	return nil
}
