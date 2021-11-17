package leetcode

func convert(s string, numRows int) string {

	if numRows <= 1 {
		return s
	}

	colStep := numRows + numRows - 2

	// 使用 8 位的整型来存储字符串的单个字符
	ss := make([]uint8, len(s), len(s))

	diagStep := colStep - 2

	i := 0

	// 开始和最后一行的字符间隔相等
	for row := 0; row < numRows; row++ {
		// 对中间行进行判别
		diag := row > 0 && row < numRows-1

		for j := row; j < len(ss); j += colStep {
			ss[i] = s[j]
			i++

			// 在这里并不更新 j 而是直接通过计算判断在一个循环内还有哪些值
			if diag && j+diagStep < len(s) {
				ss[i] = s[j+diagStep]
				i++
			}
		}

		if diag {
			diagStep -= 2
		}
	}

	return string(ss)

}


