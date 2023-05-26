package code

func oddString(words []string) string {
	if getDifference(words[0]) != getDifference(words[1]) {
		if getDifference(words[0]) == getDifference(words[2]) {
			return words[1]
		}
		return words[0]
	}

	cur := getDifference(words[0])
	for i := 2; i < len(words); i++ {
		if getDifference(words[i]) != cur {
			return words[i]
		}
	}
	return ""
}

func getDifference(str string) string {
	res := ""
	for i := 1; i < len(str); i++ {
		res += string(str[i] - str[i-1])
	}
	return res
}
