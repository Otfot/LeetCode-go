package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	type param struct {
		nums   []int
		target int
	}

	type answer struct {
		one []int
	}

	ts := []struct {
		param
		answer
	}{
		{
			param{[]int{3, 2, 4}, 6},
			answer{[]int{1, 2}},
		},

		{
			param{[]int{3, 2, 4}, 5},
			answer{[]int{0, 1}},
		},

		{
			param{[]int{0, 8, 7, 3, 3, 4, 2}, 11},
			answer{[]int{1, 3}},
		},

		{
			param{[]int{0, 1}, 1},
			answer{[]int{0, 1}},
		},

		{
			param{[]int{0, 3}, 5},
			answer{[]int{}},
		},
	}

	fmt.Printf("%10s%s%s\n", "===", " Leetcode Problem 1 ", "===")

	for _, pa := range ts {
		actual := twoSum(pa.param.nums, pa.param.target)

		if actual == nil && len(pa.answer.one) == 0 {
			continue
		} else if !reflect.DeepEqual(actual, pa.answer.one) {
			t.Errorf("twoSum(%v, %v) = %v; expected %v", pa.param.nums, pa.param.target, actual, pa.answer.one)
		}
	}

}
