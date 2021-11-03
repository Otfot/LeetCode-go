package leetcode

import (
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

	data := []struct {
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
			answer{},
		},
	}

	for _, pa := range data {
		actual := twoSum(pa.param.nums, pa.param.target)

		if !reflect.DeepEqual(actual, pa.answer.one) {
			t.Errorf("twoSum(%v, %v) = %v; expected %v", pa.param.nums, pa.param.target, actual, pa.answer.one)
		}
	}

}
