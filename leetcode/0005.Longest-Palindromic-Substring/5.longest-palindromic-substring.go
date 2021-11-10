
package leetcode

func longestPalindrome(s string) string {
	l := len(s)
	if l == 0 {
		return ""
	}

	res := &[2]int{0, 0}
	for idx, _ := range s {
		findPalindrome(idx, idx, s, res)
		findPalindrome(idx, idx+1, s, res)
	}

	return s[res[0]:res[1]]
}

func findPalindrome(left, right int, s string, sub *[2]int) {
	for left >= 0 && right < len(s) {
		if s[left] == s[right] {
			left--
			right++
		} else {
			break
		}
	}
	if right-left-1 > sub[1]-sub[0] {
		sub[0] = left + 1
		sub[1] = right
	}
}


