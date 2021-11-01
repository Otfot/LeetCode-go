package leetcode

func twoSum(numbers []int, target int) []int {

	st, end := 0, len(numbers)-1

	for st < end {
		val := numbers[st] + numbers[end]
		if target == val {
			return []int{st + 1, end + 1}
		} else if target > val {
			st++
		} else {
			end--
		}

	}

	return nil

}
