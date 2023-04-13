package code

func checkDistances(s string, distance []int) bool {

	raneDistance := map[rune]int{}
	for index, value := range s {
		if raneDistance[value] == 0 {
			raneDistance[value] = index + 1
		} else {
			if index-raneDistance[value] != distance[value-97] {
				return false
			}
		}
	}
	return true
}
