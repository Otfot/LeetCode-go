package leetcode 

func lengthOfLongestSubstring(s string) int {
	
	m, max, left := make(map[rune]int), 0, 0

	for idx, c := range s {
		if _, ok :=m[c]; ok && m[c] >= left {
			if idx - left > max {
				max = idx - left
			}
			left = m[c]+1
		}
		m[c] = idx
	}

	if len(s) - left > max {
		max = len(s) - left
	}
	return max
}


