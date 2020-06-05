package packages

// GroupAnagrams is the main function for this problem
func GroupAnagrams(strs []string) [][]string {
	m := make(map[rune][]string)
	for _, str := range strs {
		mapASCII := '0'
		sum := '0'
		mul := '0'
		for _, char := range str {
			sum += char
			mul *= char
		}
		mapASCII = mul + sum
		m[mapASCII] = append(m[mapASCII], str)
	}
	ans := [][]string{}

	for _, val := range m {
		ans = append(ans, val)
	}

	return ans
}
